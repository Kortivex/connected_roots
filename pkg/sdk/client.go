package sdk

import (
	"context"
	"net/http"
	"net/http/httptrace"
	"time"

	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/Kortivex/connected_roots/pkg/telemetry/telhttp"

	"go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"github.com/go-resty/resty/v2"
)

const (
	HeaderContentType          = "Content-Type"
	ContentTypeApplicationJSON = "application/json"
)

type ExternalAPI struct {
	RestAPIOptions        *APIConfig
	ConnectedRootsService *ConnectedRootsService
}

type APIHosts struct {
	ConnectedRootsService string
}

type APIConfig struct {
	Verbose  bool
	TimeOut  time.Duration
	APIHosts *APIHosts
	APIKey   string
	Logger   *logger.Logger
}

type ConnectedRootsService struct {
	API ConnectedRootsServiceAPI
	SDK ConnectedRootsServiceSDK
}

type Client struct {
	Host   string
	Client *resty.Client
	logger *logger.Logger
}

type APIError struct {
	Err Error `json:"error"`
}

func (e *APIError) Error() string {
	return e.Err.Message
}

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func New(options *APIConfig) *ExternalAPI {
	if options.TimeOut == 0 {
		options.TimeOut = 1 * time.Minute
	}

	logLevel := "info"
	if options.Verbose {
		logLevel = "debug"
	}
	log := logger.NewLogger(logLevel)

	c := &ExternalAPI{RestAPIOptions: options}
	c.ConnectedRootsService = NewConnectedRootsClient(options.APIHosts.ConnectedRootsService, setClient(options), &log)

	return c
}

func setClient(options *APIConfig) *resty.Client {
	httpClient := &http.Client{
		Transport: otelhttp.NewTransport( // official http transport for OTel
			telhttp.NewTransport( // Transport with options like adding custom headers based on a context key
				http.DefaultTransport,
				telhttp.WithtHeaderFromCtx(telhttp.HeaderCtx{"X-Request-Id": "cid"})), // set a header from a context key
			otelhttp.WithClientTrace(func(ctx context.Context) *httptrace.ClientTrace {
				return otelhttptrace.NewClientTrace(ctx)
			})),
	}

	client := resty.NewWithClient(httpClient)
	client.SetAuthScheme("Bearer")
	client.SetAuthToken(options.APIKey)
	client.SetDebug(options.Verbose)
	client.SetTimeout(options.TimeOut)

	return client
}
