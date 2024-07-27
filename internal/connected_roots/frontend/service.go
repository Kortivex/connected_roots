package frontend

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/ferrors"

	"github.com/labstack/echo-contrib/session"
	"github.com/wader/gormstore/v2"

	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/web/templates"

	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/pkg/httpserver"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/thejerf/suture/v4"

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
	i18n   *i18n.Bundle
	ctx    *connected_roots.Context
}

// NewService This function creates a new service object holding multiple configurations and databases like Cache, DB and RabbitMQ.
func NewService(name string, ctx *connected_roots.Context) *Service {
	loggerEmpty := ctx.Logger.NewEmpty()
	log := loggerEmpty.WithTag(commons.TagPlatformClient)

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
		i18n:   ctx.I18n,
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
	s.Echo.Debug = s.conf.Frontend.Debug
	s.Echo.Server.ReadTimeout = time.Duration(s.conf.Frontend.Timeouts.Read) * time.Second
	s.Echo.Server.WriteTimeout = time.Duration(s.conf.Frontend.Timeouts.Write) * time.Second
	s.Echo.Server.IdleTimeout = time.Duration(s.conf.Frontend.Timeouts.Idle) * time.Second
	s.Echo.Static("/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/users/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/users/profile/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/users/activities/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/users/activities/edit/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/users/activities/view/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/users/activities/delete/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/users/orchards/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/users/orchards/view/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/users/orchards/view/images", "images")
	s.Echo.Static("/users/sensors/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/admin/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/admin/roles/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/admin/roles/edit/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/admin/roles/view/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/admin/roles/delete/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/admin/users/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/admin/users/edit/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/admin/users/view/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/admin/users/delete/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/admin/orchards/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/admin/orchards/edit/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/admin/orchards/view/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/admin/orchards/view/images", "images")
	s.Echo.Static("/admin/orchards/delete/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/admin/crop-types/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/admin/crop-types/edit/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/admin/crop-types/view/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/admin/crop-types/view/images", "images")
	s.Echo.Static("/admin/crop-types/delete/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/admin/sensors/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/admin/sensors/edit/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/admin/sensors/view/assets", s.conf.Frontend.Assets)
	s.Echo.Static("/admin/sensors/delete/assets", s.conf.Frontend.Assets)
	s.Echo.Renderer = &templates.TemplateRenderer{
		Templates: template.Must(templates.ParseTemplates(s.conf.Frontend.Templates)),
	}
}

// setMiddlewares add all middlewares to "echo" server.
func (s *Service) setMiddlewares() {
	s.Params = s.conf.GetFrontendParams()

	s.Params.PreMiddlewares = []echo.MiddlewareFunc{
		echoframework.PreMiddlewareLogger(s.logger, true),
		middleware.RemoveTrailingSlash(),
	}

	s.Params.Middlewares = []echo.MiddlewareFunc{
		otelecho.Middleware(s.conf.Observability.Otel.Service),  // enable official echo middleware support for OTel
		telecho.MiddlewareRequestIDToContext("cid"),             // add the request ID to a context key, "cid" in this case.
		telecho.MiddlewareDataToSpan(telecho.DefaultDataToSpan), // enrich the span with custom data
		telecho.DumpContentToSpan(),                             // save req/resp bodies in the span

		session.Middleware(s.setSessionCookie()),
		s.I18n(),

		echoframework.PostMiddlewareLogger(250*time.Millisecond, "warn"),
	}

	s.Params.HTTPErrorHandlerMiddlewares = []httpserver.HTTPErrorHandlerMiddleware{
		telecho.MiddlewareErrorHandler, // records automatically the error in the span
		echoframework.ErrorHandlerMiddlewareLogger(s.logger),
	}

	s.Params.HTTPErrorHandler = s.errorHandler()

	s.Params.RecoverDisabled = !s.conf.Frontend.Recover

	if s.conf.App.LogLevel == debug {
		pprof.Register(s.Echo)
	}
}

func (s *Service) setSessionCookie() *gormstore.Store {
	store := gormstore.NewOptions(
		s.gorm,
		gormstore.Options{
			TableName:       s.conf.Cookie.Table,
			SkipCreateTable: false,
		},
		[]byte(s.conf.Cookie.Key),
	)

	store.SessionOpts.Secure = true
	store.SessionOpts.HttpOnly = true
	store.SessionOpts.SameSite = http.SameSiteStrictMode
	store.SessionOpts.MaxAge = s.conf.Cookie.MaxAge

	quit := make(chan struct{})
	go store.PeriodicCleanup(1*time.Second, quit)
	return store
}

func (s *Service) errorHandler() func(err error, c echo.Context) {
	return func(err error, c echo.Context) {
		var eS commons.ErrorI
		var eC *echo.HTTPError
		if ok := errors.As(err, &eC); ok {
			eS = commons.NewErrorS(eC.Code, eC.Message.(string), nil, err)
		} else if ok = errors.As(err, &eS); !ok {
			eS = commons.NewDefaultErrorS(err)
		}

		switch eS.ErrorStatus() {
		case http.StatusUnauthorized:
			if err = ferrors.Error401(c); err != nil {
				return
			}
			return
		case http.StatusForbidden:
			if err = ferrors.Error403(c); err != nil {
				return
			}
			return
		case http.StatusNotFound:
			if err = ferrors.Error404(c); err != nil {
				return
			}
			return
		default:
			if err = ferrors.Error500(c); err != nil {
				return
			}
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
				hostPort := net.JoinHostPort(s.conf.Frontend.Host, strconv.Itoa(s.conf.Frontend.Port))
				url := fmt.Sprintf("%s://%s/health/alive", "http", hostPort)
				go func() {
					ticker := time.NewTicker(time.Duration(s.conf.Frontend.Health.Frequency) * time.Second)
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
