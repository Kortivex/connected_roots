package app

import (
	"os"
	"strconv"

	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/pkg/logger"
)

// SetOtelEnvVars set main envs vars to OTEL.
func SetOtelEnvVars(conf *config.Config, logger *logger.Logger) {
	if err := os.Setenv("OTEL_SDK_DISABLED", strconv.FormatBool(conf.Observability.Otel.Disabled)); err != nil {
		logger.Fatal(err.Error())
	}

	if err := os.Setenv("OTEL_BODY_DUMP_ENABLED", strconv.FormatBool(conf.Observability.Otel.DumpEnabled)); err != nil {
		logger.Fatal(err.Error())
	}

	if err := os.Setenv("OTEL_SERVICE_NAME", conf.Observability.Otel.Service); err != nil {
		logger.Fatal(err.Error())
	}

	if err := os.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", conf.Observability.Otel.Addr); err != nil {
		logger.Fatal(err.Error())
	}
}
