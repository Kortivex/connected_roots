// Package telecho adds cross-service funcionality to the echo package.
package telecho

import (
	"bytes"
	"context"
	"io"
	"os"

	"github.com/Kortivex/connected_roots/pkg/telemetry"
	"github.com/labstack/echo/v4"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/trace"
)

const (
	// account or space value get from uris like: /spaces/:space/clusters.
	uriAccountKey = "account"
	uriSpaceKey   = "space"
)

// CtxHeader stores contextKeys and http headers, it extrats the value from header and stored it as contextKeys.
// e.g: CtxHeader{contextKey: "X-Request-Id"}.
type CtxHeader map[any]string

// MiddlewareDataToSpan is a middleware to allow set custom data to a span.
func MiddlewareDataToSpan(f func(c echo.Context, span trace.Span)) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			span := trace.SpanFromContext(c.Request().Context())

			f(c, span)

			return next(c)
		}
	}
}

// DefaultDataToSpan is a default func to use with the MiddlewareDataToSpan middleware.
func DefaultDataToSpan(c echo.Context, span trace.Span) {
	if space := c.Param(uriAccountKey); space != "" {
		span.SetAttributes(attribute.String(telemetry.TelSpaceKey, space))
	}

	if space := c.Param(uriSpaceKey); space != "" {
		span.SetAttributes(attribute.String(telemetry.TelSpaceKey, space))
	}

	cid := requestID(c)
	span.SetAttributes(attribute.String(telemetry.TelRequestIDKey, cid))

	request := c.Request()
	ctx := request.Context()

	member, _ := baggage.NewMemberRaw(telemetry.TelRequestIDKey, cid)
	b, _ := baggage.New(member)
	ctxb := baggage.ContextWithBaggage(ctx, b)

	c.SetRequest(request.WithContext(ctxb))
}

// DumpContentToSpan dumps the req/resp body to the span created by Echo.
func DumpContentToSpan() echo.MiddlewareFunc {
	var dumpContent bool

	if os.Getenv(telemetry.EnvvarBodyDumpEnabled) == "true" {
		dumpContent = true
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if telemetry.IsOTelSDKDisabled() || !dumpContent {
				return next(c)
			}

			sp := trace.SpanFromContext(c.Request().Context())
			attrs := make([]attribute.KeyValue, 0, 2)
			req := c.Request()

			var respWriter *responseDumpWriter

			if req.ContentLength > 0 && req.ContentLength < telemetry.TelMaxBytesToLog {
				body, err := io.ReadAll(req.Body)
				if err != nil {
					return err
				}

				attrs = append(attrs, attribute.String(telemetry.TelHTTPReqBodyKey, string(body)))

				req.Body = io.NopCloser(bytes.NewReader(body))
			}

			respWriter = newResponseDumpWriter(c.Response().Writer)
			c.Response().Writer = respWriter

			err := next(c)

			bodyResp := respWriter.response()
			if bodyResp.Len() > 0 && bodyResp.Len() < telemetry.TelMaxBytesToLog {
				attrs = append(attrs, attribute.String(telemetry.TelHTTPRespBodyKey, bodyResp.String()))
			}

			if len(attrs) > 0 {
				sp.AddEvent("content info", trace.WithAttributes(attrs...))
			}

			return err
		}
	}
}

// MiddlewareRequestIDToContext set the request ID to a context key interface.
func MiddlewareRequestIDToContext(contextKeys ...interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			for _, ck := range contextKeys {
				ctx := context.WithValue(c.Request().Context(), ck, requestID(c))

				c.SetRequest(c.Request().WithContext(ctx))
			}

			return next(c)
		}
	}
}

// MiddlewareRequestHeaderToContext set a header value to a context key interface.
func MiddlewareRequestHeaderToContext(ctxHeader CtxHeader) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			for ck, header := range ctxHeader {
				if _, ok := c.Request().Header[header]; ok {
					ctx := context.WithValue(c.Request().Context(), ck, c.Request().Header.Get(header))

					c.SetRequest(c.Request().WithContext(ctx))
				}
			}

			return next(c)
		}
	}
}

// requestID returns the request id of the current request from echo.Context.
func requestID(c echo.Context) string {
	rid := c.Request().Header.Get(echo.HeaderXRequestID)
	if rid == "" {
		rid = c.Response().Header().Get(echo.HeaderXRequestID)
	}

	return rid
}

// MiddlewareErrorHandler is an OTel error handler middleware that store error info in the OTel traces.
func MiddlewareErrorHandler(h echo.HTTPErrorHandler) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		// removing this can cause calling HTTPErrorHandler twice.
		// see https://github.com/labstack/echo/blob/v4.12.0/echo.go#L420
		if c.Response().Committed {
			return
		}

		telemetry.RecordSpanErrorFromContext(c.Request().Context(), err)

		h(err, c)
	}
}
