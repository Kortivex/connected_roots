// Package telemetry implements OpenTelemetry and general tracing functionality for different packages.
package telemetry

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"sync"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	sdkresource "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

const (
	// SW keys.
	TelRequestIDKey = "sw.cid"
	TelSpaceKey     = "sw.space"

	// custom keys.
	TelHTTPRespBodyKey  = "http.resp.body"
	TelHTTPReqBodyKey   = "http.req.body"
	TelMessagingBodyKey = "messaging.body"
	TelErrorStatusKey   = "error.status"
	TelErrorMessageKey  = "error.message"
	TelErrorDetailsKey  = "error.details"

	// max bytes to log.
	TelMaxBytesToLog = 60_000

	EnvvarBodyDumpEnabled = "OTEL_BODY_DUMP_ENABLED"
	envvarOtelSDKDisabled = "OTEL_SDK_DISABLED"
)

var (
	resource      *sdkresource.Resource
	resourcesOnce sync.Once

	errResource error

	// value read from envvar envvarOtelSDKDisabled.
	_isOTelSDKDisabled bool
)

// ErrorOtel is the error insterface for setting OTel errors.
type ErrorOtel interface {
	ErrorStatus() int
	ErrorMessage() string
	ErrorDetails() map[string]any
	Error() string
	Unwrap() error
}

type ShutdownFunc func(ctx context.Context) error

var noopCloserFunc = func(ctx context.Context) error { return nil }

type TracerOptions struct {
	tracerProvider *sdktrace.TracerProvider
	exporter       *otlptrace.Exporter
	resource       *sdkresource.Resource
}

type ApplyOpt func(o *TracerOptions)

// WithTracerProvider configures a TracerProvider.
func WithTracerProvider(tracerProvider *sdktrace.TracerProvider) ApplyOpt {
	return func(o *TracerOptions) {
		o.tracerProvider = tracerProvider
	}
}

// WithExporter configures an exporter, it is used when no custom TracerProvider is provided.
func WithExporter(exporter *otlptrace.Exporter) ApplyOpt {
	return func(o *TracerOptions) {
		o.exporter = exporter
	}
}

// WithResource configures a resource, it is used when no custom TracerProvider is provided.
func WithResource(resource *sdkresource.Resource) ApplyOpt {
	return func(o *TracerOptions) {
		o.resource = resource
	}
}

// getResource returns the resource.
func getResource() (*sdkresource.Resource, error) {
	resourcesOnce.Do(func() {
		var extraResources *sdkresource.Resource

		extraResources, errResource = sdkresource.New(
			context.Background(),
			sdkresource.WithFromEnv(),
			sdkresource.WithOS(),
			sdkresource.WithProcess(),
			sdkresource.WithContainer(),
			sdkresource.WithHost(),
		)
		if errResource != nil {
			return
		}

		resource, errResource = sdkresource.Merge(
			sdkresource.Default(),
			extraResources,
		)
	})

	return resource, errResource
}

// IsOTelSDKDisabled returns whether the OTel SDK is disabled or not.
func IsOTelSDKDisabled() bool {
	return _isOTelSDKDisabled
}

// InitGlobalTracerProvider creates and register a tracerProvider as a global.
func InitGlobalTracerProvider(ctx context.Context, opts ...ApplyOpt) (ShutdownFunc, error) {
	// see https://github.com/open-telemetry/opentelemetry-go/issues/3559
	if os.Getenv(envvarOtelSDKDisabled) == "true" {
		_isOTelSDKDisabled = true

		return noopCloserFunc, nil
	}

	var options TracerOptions
	for _, opt := range opts {
		opt(&options)
	}

	tp := options.tracerProvider
	if tp == nil {
		exporter := options.exporter
		if exporter == nil {
			var err error

			exporter, err = otlptracehttp.New(ctx)
			if err != nil {
				return nil, err
			}
		}

		resource := options.resource
		if resource == nil {
			var err error

			resource, err = getResource()
			if err != nil {
				return nil, err
			}
		}

		tp = sdktrace.NewTracerProvider(
			sdktrace.WithBatcher(exporter),
			sdktrace.WithResource(resource),
		)
	}

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return tp.Shutdown, nil
}

// eventAttrsFromError returns the attrs used to record a single
// error in an error event.
//
// Unwrap is not called on the error.
func eventAttrsFromError(err error) trace.SpanStartEventOption {
	const detailsEmpty = "{}"

	if eo, ok := err.(ErrorOtel); ok {
		details, errMarshal := json.Marshal(eo.ErrorDetails())
		if string(details) == "null" || errMarshal != nil {
			details = []byte(detailsEmpty)
		}

		return trace.WithAttributes(
			attribute.Int(TelErrorStatusKey, eo.ErrorStatus()),
			attribute.String(TelErrorMessageKey, eo.Error()),
			attribute.String(TelErrorDetailsKey, string(details)),
		)
	}

	return trace.WithAttributes(
		attribute.Int(TelErrorStatusKey, 0),
		attribute.String(TelErrorMessageKey, err.Error()),
		attribute.String(TelErrorDetailsKey, detailsEmpty),
	)
}

// AddErrorEventToSpan add an error envent to the span unwrapping the err.
func AddErrorEventToSpan(sp trace.Span, err error) {
	var attrs []trace.EventOption

	for e := err; e != nil; e = errors.Unwrap(e) {
		attrs = append(attrs, eventAttrsFromError(e))
	}

	if attrs != nil {
		sp.AddEvent("error", attrs...)
	}
}

// RecordSpanErrorFromContext extracts the current span from the context
// and marks the span as a span error.
// It also records an error event.
func RecordSpanErrorFromContext(ctx context.Context, err error) {
	sp := trace.SpanFromContext(ctx)

	RecordSpanError(sp, err)
}

// RecordSpanError marks the span as a span error.
// It also records an error event.
func RecordSpanError(sp trace.Span, err error) {
	sp.RecordError(err)
	sp.SetStatus(codes.Error, err.Error())

	AddErrorEventToSpan(sp, err)
}
