package httpserver

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/pkg/httpserver"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/thejerf/suture/v4"

	errs "github.com/Kortivex/connected_roots/internal/connected_roots/httpserver/errors"

	"github.com/Kortivex/connected_roots/pkg/telemetry/telecho"
	"github.com/labstack/echo/v4/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"

	"github.com/Kortivex/connected_roots/pkg/logger/extend/echoframework"

	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/Kortivex/connected_roots/pkg/logger/commons"

	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/pkg/service"
	"github.com/labstack/echo-contrib/pprof"
	"gorm.io/gorm"
)

const debug = "debug"

type Service struct {
	service.Service
	Echo   *echo.Echo
	Params httpserver.Params
	gorm   *gorm.DB
	logger *logger.Logger
	conf   *config.Config
	ctx    *connected_roots.Context
}

// NewService This function creates a new service object holding multiple configurations and databases like Cache, DB and RabbitMQ.
func NewService(name string, ctx *connected_roots.Context) *Service {
	loggerEmpty := ctx.Logger.NewEmpty()
	log := loggerEmpty.WithTag(commons.TagPlatformHttpserver)

	srv := &Service{
		Service: service.Service{
			Name:     name,
			Started:  make(chan bool),
			Status:   make(chan int),
			Release:  make(chan bool),
			Stop:     make(chan bool),
			Existing: 0,
			M:        sync.Mutex{},
			Running:  false,
		},
		Echo:   echo.New(),
		gorm:   ctx.Gorm,
		logger: log,
		conf:   ctx.Conf,
		ctx:    ctx,
	}

	return srv
}

// provide This function sets up an Echo server and registers the app's routes, preparing the app to listen for requests.
func (s *Service) provide() {
	// Set Config for echo.
	s.setSetup()
	// Set Middlewares for echo.
	s.setMiddlewares()

	s.logger.Debug("Config Loaded")

	// Set all HTTP Routes.
	s.registerRoutes(s.ctx)
	s.logger.Debug("Routes Loaded")

	s.logger.Debug("Starting http server")
	go httpserver.Start(context.Background(), s.Echo, s.Params)
}

// setSetup apply the standard configuration to the "echo" server.
func (s *Service) setSetup() {
	s.Echo.Logger = echoframework.NewLogger(s.logger)
	s.Echo.HideBanner = true
	s.Echo.HidePort = true
	s.Echo.Debug = s.conf.API.Debug
	s.Echo.Server.ReadTimeout = time.Duration(s.conf.API.Timeouts.Read) * time.Second
	s.Echo.Server.WriteTimeout = time.Duration(s.conf.API.Timeouts.Write) * time.Second
	s.Echo.Server.IdleTimeout = time.Duration(s.conf.API.Timeouts.Idle) * time.Second
}

// setMiddlewares add all middlewares to "echo" server.
func (s *Service) setMiddlewares() {
	s.Params = s.conf.GetAPIParams()

	s.Params.PreMiddlewares = []echo.MiddlewareFunc{
		echoframework.PreMiddlewareLogger(s.logger, true),
		middleware.RemoveTrailingSlash(),
	}

	s.Params.Middlewares = []echo.MiddlewareFunc{
		otelecho.Middleware(s.conf.Observability.Otel.Service),  // enable official echo middleware support for OTel
		telecho.MiddlewareRequestIDToContext("cid"),             // add the request ID to a context key, "cid" in this case.
		telecho.MiddlewareDataToSpan(telecho.DefaultDataToSpan), // enrich the span with custom data
		telecho.DumpContentToSpan(),                             // save req/resp bodies in the span

		echoframework.PostMiddlewareLogger(250*time.Millisecond, "warn"),

		middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
			KeyLookup:  "header:" + echo.HeaderAuthorization,
			AuthScheme: "Bearer",
			Validator: func(key string, c echo.Context) (bool, error) {
				return key == s.conf.API.APIKey, nil
			},
		}),
	}

	s.Params.HTTPErrorHandlerMiddlewares = []httpserver.HTTPErrorHandlerMiddleware{
		telecho.MiddlewareErrorHandler, // records automatically the error in the span
		echoframework.ErrorHandlerMiddlewareLogger(s.logger),
	}

	s.Params.HTTPErrorHandler = s.errorHandler()

	s.Params.RecoverDisabled = !s.conf.API.Recover

	if s.conf.App.LogLevel == debug {
		pprof.Register(s.Echo)
	}
}

func (s *Service) errorHandler() func(err error, c echo.Context) {
	return func(err error, c echo.Context) {
		var eS commons.ErrorI
		if ok := errors.As(err, &eS); !ok {
			eS = commons.NewDefaultErrorS(err)
		}

		if eS.ErrorStatus() == 0 {
			err = &errs.APIResponseError{Err: commons.ErrorS{
				Status:  http.StatusInternalServerError,
				Message: "something went wrong",
			}}
			if err = c.JSON(http.StatusInternalServerError, err); err != nil {
				return
			}
			return
		}

		if err = c.JSON(eS.ErrorStatus(), err); err != nil {
			return
		}
	}
}

// Serve method from interface suture.Service to handle Service cycle life.
func (s *Service) Serve(ctx context.Context) error {
	s.M.Lock()
	if s.Existing != 0 {
		(&sync.Mutex{}).Unlock()
	}
	s.Existing++
	s.Running = true
	s.M.Unlock()

	defer func() {
		s.M.Lock()
		s.Running = false
		s.M.Unlock()
	}()

	releaseExistence := func() {
		s.M.Lock()
		s.Existing--
		s.M.Unlock()
	}

	s.Started <- true

	useStopChan := false

	for {
		select {
		case val := <-s.Status:
			switch val {
			case service.Run:
				// Start Service
				s.provide()
			case service.Heartbeat:
				client := resty.New().R().EnableTrace()
				client.SetAuthScheme("Bearer")
				client.SetAuthToken(s.conf.API.APIKey)
				hostPort := net.JoinHostPort(s.conf.API.Host, strconv.Itoa(s.conf.API.Port))
				url := fmt.Sprintf("%s://%s/health/alive", "http", hostPort)
				go func() {
					ticker := time.NewTicker(time.Duration(s.conf.API.Health.Frequency) * time.Second)
					for range ticker.C {
						if _, err := client.Get(url); err != nil {
							s.logger.Debug(service.PingKO)
							s.logger.Error(err.Error())
							releaseExistence()
							os.Exit(1)
						}
						s.logger.Debug(service.PingOK)
					}
				}()
			case service.Fail:
				releaseExistence()
				if useStopChan {
					s.Stop <- true
				}
				return nil
			case service.Panic:
				releaseExistence()
				panic(service.ErrPanicService)
			case service.Hang:
				<-s.Release
			case service.UseStopChan:
				useStopChan = true
			case service.TerminateTree:
				return suture.ErrTerminateSupervisorTree
			case service.DoNotRestart:
				return suture.ErrDoNotRestart
			}
		case <-ctx.Done():
			releaseExistence()
			if useStopChan {
				s.Stop <- true
			}
			return fmt.Errorf(service.ErrFailureServiceEnding+": %w", ctx.Err())
		}
	}
}
