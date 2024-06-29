package app

import (
	"context"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"

	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend"
	"github.com/Kortivex/connected_roots/pkg/sdk"

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
	Frontend      = "service.frontend"
)

// Start This function sets up the Supervisor, adds the various services as supervisor children, prints the banner, sets up the telemetry, and starts the services.
func Start() {
	supervisor := suture.NewSimple("Supervisor")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Init Supervisor.
	supervisor.ServeBackground(ctx)

	// Init Base Services.
	appCtx := &connected_roots.Context{}
	// - Init Config.
	addConfigService(supervisor, appCtx)
	// - Print Banner.
	printBanner(appCtx)
	// - Init Logger.
	addLoggerService(supervisor, appCtx)

	// Set Telemetry Env.
	setTelemetry(appCtx)

	// Init Global Tracer.
	closer, err := telemetry.InitGlobalTracerProvider(ctx)
	if err != nil {
		appCtx.Logger.Fatal(fmt.Errorf("error on initTracer: %w", err).Error())
	}
	defer func() {
		if err = closer(ctx); err != nil {
			appCtx.Logger.Error(fmt.Errorf("error on shutdown tracer: %w", err).Error())
		}
	}()

	// Ini SDK.
	setAPIConfig(appCtx)

	// Init i18n config.
	setI18NConfig(appCtx)

	// Init Core Services.
	addDBService(supervisor, appCtx)

	// Init HTTP Server.
	addHTTPService(supervisor, appCtx)

	// Init Frontend Server.
	addFrontend(supervisor, appCtx)
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
}

func addFrontend(supervisor *suture.Supervisor, ctx *connected_roots.Context) {
	// Add HTTP (Echo) Frontend Service.
	frontendSrv := frontend.NewService(Frontend, ctx)
	supervisor.Add(frontendSrv)
	<-frontendSrv.Started
	frontendSrv.Status <- service.UseStopChan
	frontendSrv.Status <- service.Run
	frontendSrv.Status <- service.Heartbeat

	<-frontendSrv.Stop
}

func setAPIConfig(appCtx *connected_roots.Context) {
	appCtx.SDK = sdk.New(
		&sdk.APIConfig{
			APIKey: appCtx.Conf.Thirds.SDK.ConnectedRootsService.APIKey,
			APIHosts: &sdk.APIHosts{
				ConnectedRootsService: appCtx.Conf.Thirds.SDK.ConnectedRootsService.Host,
			},
			Verbose: appCtx.Conf.Thirds.SDK.Verbose,
		},
	)
}

func setI18NConfig(appCtx *connected_roots.Context) {
	appCtx.I18n = i18n.NewBundle(language.English)
	appCtx.I18n.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	appCtx.I18n.MustLoadMessageFile(appCtx.Conf.I18n.Path + "/" + appCtx.Conf.I18n.En)
	appCtx.I18n.MustLoadMessageFile(appCtx.Conf.I18n.Path + "/" + appCtx.Conf.I18n.Es)
}

func printBanner(ctx *connected_roots.Context) {
	fmt.Println(GenerateBanner(ctx.Conf))
}

func setTelemetry(ctx *connected_roots.Context) {
	SetOtelEnvVars(ctx.Conf, ctx.Logger)
}
