// Package httpserver provides an http server wrapping echo framework.
package httpserver

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"slices"
	"syscall"
	"time"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	shutdownTimeout = 30 * time.Second
	readTimeout     = 10 * time.Second
	writeTimeout    = 10 * time.Second
	idleTimeout     = 15 * time.Second
	bodyLimit       = "1M"
)

// ServerShutdownFunc is a func to run where the server shutdowns.
type ServerShutdownFunc func(e *echo.Echo)

// Params are the server config params.
type Params struct {
	Port                        string
	BodyLimit                   string
	PrometheusServiceName       string
	PrometheusDisabled          *bool
	OnServerShutdown            []ServerShutdownFunc
	WriteTimeout                *time.Duration
	ReadTimeout                 *time.Duration
	IdleTimeout                 *time.Duration
	PreMiddlewares              []echo.MiddlewareFunc
	Middlewares                 []echo.MiddlewareFunc
	HTTPErrorHandler            echo.HTTPErrorHandler
	HTTPErrorHandlerMiddlewares []HTTPErrorHandlerMiddleware
	RecoverDisabled             bool
}

// Start starts the server.
func Start(ctx context.Context, e *echo.Echo, params Params) {
	wTimeout := writeTimeout
	if params.WriteTimeout != nil {
		wTimeout = *params.WriteTimeout
	}

	rTimeout := readTimeout
	if params.ReadTimeout != nil {
		rTimeout = *params.ReadTimeout
	}

	iTimeout := idleTimeout
	if params.IdleTimeout != nil {
		iTimeout = *params.IdleTimeout
	}

	if params.BodyLimit == "" {
		params.BodyLimit = bodyLimit
	}

	e.Server.ReadTimeout = rTimeout
	e.Server.WriteTimeout = wTimeout
	e.Server.IdleTimeout = iTimeout

	// Pre-middlewares registering
	e.Pre(middleware.RemoveTrailingSlash())
	for _, m := range params.PreMiddlewares {
		e.Pre(m)
	}

	// Post-middlewares registering
	e.Use(middleware.RequestID())
	e.Use(middleware.BodyLimit(params.BodyLimit))
	for _, m := range params.Middlewares {
		e.Use(m)
	}

	if params.PrometheusDisabled != nil && !*params.PrometheusDisabled {
		p := prometheus.NewPrometheus(params.PrometheusServiceName, nil)
		p.Use(e)
	}

	e.HTTPErrorHandler = e.DefaultHTTPErrorHandler

	// Add HTTPError handler.
	if params.HTTPErrorHandler != nil {
		e.HTTPErrorHandler = params.HTTPErrorHandler
	}

	// Add HTTPErrorHandler middlewares
	slices.Reverse(params.HTTPErrorHandlerMiddlewares)
	for _, m := range params.HTTPErrorHandlerMiddlewares {
		e.HTTPErrorHandler = m(e.HTTPErrorHandler)
	}

	// Add Recover middleware with default config. Print Stack Traces and Recover from 'panic'.
	if !params.RecoverDisabled {
		e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{}))
	}

	e.HideBanner = true
	e.HidePort = true
	e.DisableHTTP2 = true

	done := make(chan struct{})
	go gracefulShutdown(ctx, e, params, done)

	e.Logger.Error(e.Start(":" + params.Port).Error())

	<-done
}

// gracefulShutdown graceful shutdown callback. This method is waiting for interrupt signals or context is done.
// When signal is received or context is done the Params.OnServerShutdown are called.
func gracefulShutdown(ctx context.Context, server *echo.Echo, params Params, done chan<- struct{}) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
	case <-quit:
	}

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("could not gracefully shutdown the server")
	}

	for _, f := range params.OnServerShutdown {
		f(server)
	}

	close(done)
}

// GetRequestID returns the correlation-id or request-id based on echo standard headers.
//
// 'X-Correlation-Id' or 'X-Request-Id'.
func GetRequestID(c echo.Context) string {
	req := c.Request()
	res := c.Response()
	id := req.Header.Get(echo.HeaderXCorrelationID)
	if id == "" {
		id = res.Header().Get(echo.HeaderXCorrelationID)
		if id == "" {
			id = req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}
		}
	}

	return id
}

// GetHeader returns header value if not exist returns defaultValue.
func GetHeader(c echo.Context, key string, defaultValue string) string {
	value := c.Request().Header.Get(key)

	if value == "" {
		return defaultValue
	}

	return value
}

// SetHeader add header to Request.
func SetHeader(c echo.Context, key string, value string) {
	c.Response().Header().Set(key, value)
}

// FindRouteName returns the route name of the current request.
func FindRouteName(c echo.Context) string {
	for _, r := range c.Echo().Routes() {
		if r.Path == c.Path() && r.Method == c.Request().Method {
			return r.Name
		}
	}

	return "-"
}
