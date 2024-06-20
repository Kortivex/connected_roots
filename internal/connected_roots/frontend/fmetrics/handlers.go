package fmetrics

import (
	"net/http"

	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/pkg/logger"

	"go.opentelemetry.io/otel"

	"github.com/labstack/echo/v4"
)

const (
	tracingMetricsHandlers = "http-handler.metrics"

	tracingMetricsGetUnitsHandler = "http-handler.metrics.get-unit-handler"
)

type MetricsHandlers struct {
	conf   *config.Config
	logger *logger.Logger
}

// NewMetricsHandlers This function builds and returns a new UnitsHandlers object with all necessary dependencies.
func NewMetricsHandlers(ctx *connected_roots.Context) *MetricsHandlers {
	loggerEmpty := ctx.Logger.NewEmpty()
	log := loggerEmpty.WithTag(tracingMetricsHandlers)

	return &MetricsHandlers{
		logger: log,
		conf:   ctx.Conf,
	}
}

// GetHealthAliveHandler This function returns an "OK" response, indicating that the status of the service is healthy.
func (h *MetricsHandlers) GetHealthAliveHandler(c echo.Context) error {
	_, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingMetricsGetUnitsHandler)
	defer span.End()

	loggerEmpty := h.logger.New()
	log := loggerEmpty.WithTag(tracingMetricsGetUnitsHandler)

	log.Debug("health alive called")

	return c.String(http.StatusOK, "OK")
}
