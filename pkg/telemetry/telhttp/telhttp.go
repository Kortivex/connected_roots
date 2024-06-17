// Package telhttp adds cross-service funcionality to the http package.
package telhttp

import (
	"bytes"
	"io"
	"mime"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/Kortivex/connected_roots/pkg/telemetry"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const (
	_mimeJSON = "application/json"
)

// NewTransport returns a new Transport.
func NewTransport(base http.RoundTripper, opts ...ApplyOpt) *Transport {
	if base == nil {
		base = http.DefaultTransport
	}

	var options OptionsTransport

	if os.Getenv(telemetry.EnvvarBodyDumpEnabled) == "true" {
		options.dumpContent = true
	}

	for _, opt := range opts {
		opt(&options)
	}

	return &Transport{
		rt:            base,
		headerContext: options.headerContext,
		dumpContent:   options.dumpContent,
	}
}

// Transport is a transport that can propagate context keys as http request headers.
type Transport struct {
	rt            http.RoundTripper
	headerContext HeaderCtx
	dumpContent   bool
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()

	r := req.Clone(ctx) // According to RoundTripper spec, we shouldn't modify the origin request.

	for header, contextKey := range t.headerContext {
		if v := ctx.Value(contextKey); v != nil {
			if v, ok := v.(string); ok {
				r.Header.Add(header, v)
			}
		}
	}

	sp := trace.SpanFromContext(ctx)
	attrs := make([]attribute.KeyValue, 0, 2)

	if t.isRequestLogabble(r) {
		bodyR := &bytes.Buffer{}
		reader := io.TeeReader(r.Body, bodyR)

		body, err := io.ReadAll(reader)
		if err != nil {
			return nil, err
		}

		r.Body.Close()
		r.Body = io.NopCloser(bodyR)

		attrs = append(attrs, attribute.String(telemetry.TelHTTPReqBodyKey, string(body)))
	}

	resp, err := t.rt.RoundTrip(r)

	if t.isResponseLoggable(resp) {
		bodyR := &bytes.Buffer{}
		reader := io.TeeReader(resp.Body, bodyR)

		body, err := io.ReadAll(reader)
		if err != nil {
			return resp, err
		}

		resp.Body.Close()
		resp.Body = io.NopCloser(bodyR)

		attrs = append(attrs, attribute.String(telemetry.TelHTTPRespBodyKey, string(body)))
	}

	if len(attrs) > 0 {
		sp.AddEvent("content info", trace.WithAttributes(attrs...))
	}

	return resp, err
}

func (t *Transport) isRequestLogabble(r *http.Request) bool {
	return !telemetry.IsOTelSDKDisabled() &&
		t.dumpContent &&
		r != nil &&
		r.Body != nil &&
		hasMimeType(r.Header.Get("Content-type"), _mimeJSON) &&
		r.ContentLength < telemetry.TelMaxBytesToLog
}

func (t *Transport) isResponseLoggable(r *http.Response) bool {
	return !telemetry.IsOTelSDKDisabled() &&
		t.dumpContent &&
		r != nil &&
		r.Body != nil &&
		hasMimeType(r.Header.Get("Content-type"), _mimeJSON) &&
		r.ContentLength < telemetry.TelMaxBytesToLog
}

func hasMimeType(contentType string, mimeType ...string) bool {
	for _, v := range strings.Split(contentType, ",") {
		t, _, err := mime.ParseMediaType(v)
		if err != nil {
			continue
		}

		if slices.Contains(mimeType, t) {
			return true
		}
	}

	return false
}

// HeaderCtx stores http headers and context keys to be sent as http request headers.
//
// The underlying value of the context key is expected to be a string type, otherwise it won't be send.
// e.g: HeaderCtx{"X-Request-Id": contextKey}.
type HeaderCtx map[string]any

// OptionsTransport are the Transport options.
type OptionsTransport struct {
	headerContext HeaderCtx

	// create an event span with the request/response body.
	dumpContent bool
}

// ApplyOpt sets the option in the OptionsTransport struct.
type ApplyOpt func(o *OptionsTransport)

// WithtHeaderFromCtx configures the headers-contextKeys to be sent as http request headers.
func WithtHeaderFromCtx(ch HeaderCtx) ApplyOpt {
	return func(o *OptionsTransport) {
		o.headerContext = ch
	}
}
