package sensor

import (
	"fmt"
	"github.com/Kortivex/connected_roots/internal/connected_roots/httpserver/errors"
	"github.com/Kortivex/connected_roots/internal/connected_roots/sensor"
	"net/http"

	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

const (
	tracingSensorHandlers = "http-handler.sensor"

	tracingPostSensorHandlers   = "http-handler.sensor.post-sensor"
	tracingPutSensorHandlers    = "http-handler.sensor.put-sensor"
	tracingGetSensorHandlers    = "http-handler.sensor.get-sensor"
	tracingListSensorHandlers   = "http-handler.sensor.list-sensors"
	tracingDeleteSensorHandlers = "http-handler.sensor.delete-sensor"

	sensorIDParam = "sensor_id"
)

type SensorsHandlers struct {
	gorm   *gorm.DB
	logger *logger.Logger
	conf   *config.Config
	// Services.
	sensorSvc *sensor.Service
}

func NewSensorsHandlers(appCtx *connected_roots.Context) *SensorsHandlers {
	loggerEmpty := appCtx.Logger.NewEmpty()
	log := loggerEmpty.WithTag(tracingSensorHandlers)

	return &SensorsHandlers{
		gorm:      appCtx.Gorm,
		logger:    log,
		conf:      appCtx.Conf,
		sensorSvc: sensor.New(appCtx.Conf, appCtx.Gorm, appCtx.Logger),
	}
}

func (h *SensorsHandlers) PostSensorHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingPostSensorHandlers)
	defer span.End()

	sensorBody := connected_roots.Sensors{}
	if err := c.Bind(&sensorBody); err != nil {
		err = fmt.Errorf("%s: %w", tracingPostSensorHandlers, errors.ErrInvalidPayload)
		return errors.NewErrorResponse(c, err)
	}

	sensorRes, err := h.sensorSvc.Save(ctx, &sensorBody)
	if err != nil {
		err = fmt.Errorf("%s: %w", tracingPostSensorHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated, sensorRes)
}

func (h *SensorsHandlers) PutSensorHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingPutSensorHandlers)
	defer span.End()

	sensorID := c.Param(sensorIDParam)
	if sensorID == "" {
		return errors.NewErrorResponse(c, errors.ErrPathParamInvalidValue)
	}

	sensorBody := connected_roots.Sensors{}
	if err := c.Bind(&sensorBody); err != nil {
		err = fmt.Errorf("%s: %w", tracingPutSensorHandlers, errors.ErrInvalidPayload)
		return errors.NewErrorResponse(c, err)
	}
	sensorBody.ID = sensorID

	sensorRes, err := h.sensorSvc.Update(ctx, &sensorBody)
	if err != nil {
		err = fmt.Errorf("%s: %w", tracingPutSensorHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, sensorRes)
}

func (h *SensorsHandlers) GetSensorHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingGetSensorHandlers)
	defer span.End()

	sensorID := c.Param(sensorIDParam)
	if sensorID == "" {
		return errors.NewErrorResponse(c, errors.ErrPathParamInvalidValue)
	}

	sensorRes, err := h.sensorSvc.Obtain(ctx, sensorID)
	if err != nil {
		err = fmt.Errorf("%s: %w", tracingGetSensorHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, sensorRes)
}

func (h *SensorsHandlers) ListSensorsHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingListSensorHandlers)
	defer span.End()

	filters := connected_roots.SensorPaginationFilters{}
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &filters); err != nil {
		err = fmt.Errorf("%s: %w", tracingListSensorHandlers, errors.ErrQueryParamInvalidValue)
		return errors.NewErrorResponse(c, err)
	}

	sensorsRes, err := h.sensorSvc.ObtainAll(ctx, &filters)
	if err != nil {
		err = fmt.Errorf("%s: %w", tracingListSensorHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, sensorsRes)
}

func (h *SensorsHandlers) DeleteSensorHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingDeleteSensorHandlers)
	defer span.End()

	sensorID := c.Param(sensorIDParam)
	if sensorID == "" {
		return errors.NewErrorResponse(c, errors.ErrPathParamInvalidValue)
	}

	if err := h.sensorSvc.Remove(ctx, sensorID); err != nil {
		err = fmt.Errorf("%s: %w", tracingDeleteSensorHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}
