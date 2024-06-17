package app

import (
	"context"
	"fmt"
	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/pkg/telemetry"
	"github.com/thejerf/suture/v4"

	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/internal/connected_roots/httpserver"
	"github.com/Kortivex/connected_roots/internal/connected_roots/logger"
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql"
	"github.com/Kortivex/connected_roots/pkg/service"
)

const (
	ConfigService = "service.config"
	LoggerService = "service.logger"
	DBService     = "service.db"
	HTTPService   = "service.httpserver"
)

// Start This function sets up the Supervisor, adds the various services as supervisor children, prints the banner, sets up the telemetry, and starts the services.
func Start() {
	supervisor := suture.NewSimple("Supervisor")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Init Supervisor.
	supervisor.ServeBackground(ctx)

	// Init Base Services.
	globCtx := &connected_roots.Context{}
	// - Init Config.
	addConfigService(supervisor, globCtx)
	// - Print Banner.
	printBanner(globCtx)
	// - Init Logger.
	addLoggerService(supervisor, globCtx)

	// Set Telemetry Env.
	setTelemetry(globCtx)

	// Init Global Tracer.
	closer, err := telemetry.InitGlobalTracerProvider(ctx)
	if err != nil {
		globCtx.Logger.Fatal(fmt.Errorf("error on initTracer: %w", err).Error())
	}
	defer func() {
		if err = closer(ctx); err != nil {
			globCtx.Logger.Error(fmt.Errorf("error on shutdown tracer: %w", err).Error())
		}
	}()

	// Init Core Services.
	addDBService(supervisor, globCtx)

	// Init HTTP Server.
	addHTTPService(supervisor, globCtx)
}

func addConfigService(supervisor *suture.Supervisor, ctx *connected_roots.Context) {
	configSrv := config.NewService(ConfigService)
	supervisor.Add(configSrv)
	<-configSrv.Started

	ctx.Conf = configSrv.Conf
}

func addLoggerService(supervisor *suture.Supervisor, ctx *connected_roots.Context) {
	loggerSrv := logger.NewService(LoggerService, ctx.Conf)
	supervisor.Add(loggerSrv)
	<-loggerSrv.Started

	ctx.Logger = loggerSrv.Logger
}

func addDBService(supervisor *suture.Supervisor, ctx *connected_roots.Context) {
	dbSrv := postgresql.NewService(DBService, ctx.Conf, ctx.Logger)
	supervisor.Add(dbSrv)
	<-dbSrv.Started
	dbSrv.Status <- service.Run
	dbSrv.Status <- service.Heartbeat

	ctx.Gorm = dbSrv.Gorm
}

func addHTTPService(supervisor *suture.Supervisor, ctx *connected_roots.Context) {
	// Add HTTP (Echo) Service.
	httpSrv := httpserver.NewService(HTTPService, ctx)
	supervisor.Add(httpSrv)
	<-httpSrv.Started
	httpSrv.Status <- service.UseStopChan
	httpSrv.Status <- service.Run
	httpSrv.Status <- service.Heartbeat

	<-httpSrv.Stop
}

func printBanner(ctx *connected_roots.Context) {
	fmt.Println(GenerateBanner(ctx.Conf))
}

func setTelemetry(ctx *connected_roots.Context) {
	SetOtelEnvVars(ctx.Conf, ctx.Logger)
}
