package sensor

import (
	"fmt"
	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/bars"
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/ferrors"
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/i18n/translator"
	sessionServ "github.com/Kortivex/connected_roots/internal/connected_roots/session"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/Kortivex/connected_roots/pkg/logger/commons"
	"github.com/Kortivex/connected_roots/pkg/sdk"
	"github.com/Kortivex/connected_roots/pkg/sdk/sdk_models"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
	"net/http"
	"time"
)

const (
	tracingSensorHandlers = "http-handler.sensor"

	getCreateSensorHandlerName  = "http-handler.orchard.get-create-sensor"
	postCreateSensorHandlerName = "http-handler.orchard.post-create-sensor"
	getUpdateSensorHandlerName  = "http-handler.orchard.get-update-sensor"
	postUpdateSensorHandlerName = "http-handler.orchard.post-update-sensor"
	getViewSensorHandlerName    = "http-handler.orchard.get-view-sensor"
	getListSensorHandlerName    = "http-handler.orchard.get-list-sensors"
	getDeleteSensorHandlerName  = "http-handler.orchard.get-delete-sensor"
	postDeleteSensorHandlerName = "http-handler.orchard.post-delete-sensor"

	sensorIDParam = "sensor_id"
)

type Handlers struct {
	gorm   *gorm.DB
	logger *logger.Logger
	conf   *config.Config
	sdk    *sdk.ExternalAPI
	// Services
	sessionSvc *sessionServ.Service
}

func NewSensorsHandlers(appCtx *connected_roots.Context) *Handlers {
	loggerEmpty := appCtx.Logger.NewEmpty()
	log := loggerEmpty.WithTag(tracingSensorHandlers)

	return &Handlers{
		gorm:       appCtx.Gorm,
		logger:     log,
		conf:       appCtx.Conf,
		sdk:        appCtx.SDK,
		sessionSvc: sessionServ.New(appCtx.Conf, appCtx.Gorm, appCtx.Logger),
	}
}

func (h *Handlers) GetSensorCreateHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getCreateSensorHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(getCreateSensorHandlerName)

	sess, err := h.sessionSvc.Obtain(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	isAdminTech, err := h.sessionSvc.IsAdminTechnical(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}
	if !isAdminTech {
		return commons.NewErrorS(http.StatusUnauthorized, "forbidden", nil, ferrors.ErrUnauthorized)
	}

	orchards, _, err := h.sdk.ConnectedRootsService.SDK.ObtainOrchards(ctx, "10000", "", "", nil, nil, nil)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-sensors-create.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(ctx, c, h.sessionSvc),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonSensorCreatePageI18N(c)), map[string]interface{}{
			"orchards": orchards,
		}))
}

func (h *Handlers) PostSensorCreateHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), postCreateSensorHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	log := loggerNew.WithTag(postCreateSensorHandlerName)

	sess, err := h.sessionSvc.Obtain(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	isAdminTech, err := h.sessionSvc.IsAdminTechnical(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}
	if !isAdminTech {
		return commons.NewErrorS(http.StatusUnauthorized, "forbidden", nil, ferrors.ErrUnauthorized)
	}

	sensor := &sdk_models.SensorsBody{
		Name:      c.FormValue("name"),
		Location:  c.FormValue("location"),
		OrchardID: c.FormValue("orchard-id"),
	}

	log.Debug(fmt.Sprintf("sensor: %+v", sensor))

	_, err = h.sdk.ConnectedRootsService.SDK.SaveSensor(ctx, sensor)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-sensors-create.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(ctx, c, h.sessionSvc),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonSensorCreatePageI18N(c)), map[string]interface{}{
			"notification_type":    "success",
			"notification_title":   translator.T(c, translator.NotificationsAdminSensorsCreateSuccessTitle),
			"notification_message": translator.T(c, translator.NotificationsAdminSensorsCreateSuccessMessage),
		}))
}

func (h *Handlers) GetSensorUpdateHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getUpdateSensorHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(getUpdateSensorHandlerName)

	sensorId := c.Param(sensorIDParam)
	if sensorId == "" {
		err := ferrors.ErrPathParamInvalidValue
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	sess, err := h.sessionSvc.Obtain(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	isAdminTech, err := h.sessionSvc.IsAdminTechnical(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}
	if !isAdminTech {
		return commons.NewErrorS(http.StatusUnauthorized, "forbidden", nil, ferrors.ErrUnauthorized)
	}

	orchards, _, err := h.sdk.ConnectedRootsService.SDK.ObtainOrchards(ctx, "10000", "", "", nil, nil, nil)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	sensor, err := h.sdk.ConnectedRootsService.SDK.ObtainSensor(ctx, sensorId)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-sensors-update.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(ctx, c, h.sessionSvc),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonSensorUpdatePageI18N(c)), map[string]interface{}{
			"orchards": orchards,
			"sensor":   sensor,
		}))
}

func (h *Handlers) PostSensorUpdateHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), postUpdateSensorHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	log := loggerNew.WithTag(postUpdateSensorHandlerName)

	sensorId := c.Param(sensorIDParam)
	if sensorId == "" {
		err := ferrors.ErrPathParamInvalidValue
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	sess, err := h.sessionSvc.Obtain(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	isAdminTech, err := h.sessionSvc.IsAdminTechnical(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}
	if !isAdminTech {
		return commons.NewErrorS(http.StatusUnauthorized, "forbidden", nil, ferrors.ErrUnauthorized)
	}

	sensor, err := h.sdk.ConnectedRootsService.SDK.ObtainSensor(ctx, sensorId)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	sensor.Name = c.FormValue("name")
	sensor.Location = c.FormValue("location")
	sensor.CalibrationDate = time.Now()
	sensor.OrchardID = c.FormValue("orchard-id")

	log.Debug(fmt.Sprintf("sensor: %+v", sensor))

	orchards, _, err := h.sdk.ConnectedRootsService.SDK.ObtainOrchards(ctx, "10000", "", "", nil, nil, nil)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	_, err = h.sdk.ConnectedRootsService.SDK.UpdateSensor(ctx, sensor.ToSensorBody())
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-sensors-update.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(ctx, c, h.sessionSvc),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonSensorUpdatePageI18N(c)), map[string]interface{}{
			"orchards":             orchards,
			"sensor":               sensor,
			"notification_type":    "success",
			"notification_title":   translator.T(c, translator.NotificationsAdminSensorsUpdateSuccessTitle),
			"notification_message": translator.T(c, translator.NotificationsAdminSensorsUpdateSuccessMessage),
		}))
}

func (h *Handlers) GetSensorViewHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getViewSensorHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(getViewSensorHandlerName)

	sensorId := c.Param(sensorIDParam)
	if sensorId == "" {
		err := ferrors.ErrPathParamInvalidValue
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	sess, err := h.sessionSvc.Obtain(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	isAdminTech, err := h.sessionSvc.IsAdminTechnical(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}
	if !isAdminTech {
		return commons.NewErrorS(http.StatusUnauthorized, "forbidden", nil, ferrors.ErrUnauthorized)
	}

	sensor, err := h.sdk.ConnectedRootsService.SDK.ObtainSensor(ctx, sensorId)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-sensors-view.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(ctx, c, h.sessionSvc),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonSensorViewPageI18N(c)), map[string]interface{}{
			"sensor": sensor,
		}))
}

func (h *Handlers) GetSensorsListHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getListSensorHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	log := loggerNew.WithTag(getListSensorHandlerName)

	isAdminTech, err := h.sessionSvc.IsAdminTechnical(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}
	if !isAdminTech {
		return commons.NewErrorS(http.StatusUnauthorized, "forbidden", nil, ferrors.ErrUnauthorized)
	}

	message, err := h.sessionSvc.ObtainMessage(ctx, c, "message")
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	log.Debug(fmt.Sprintf("message: %s", message))

	notifications := map[string]interface{}{}
	if len(message) > 0 && message[0] == "success" {
		notifications = map[string]interface{}{
			"notification_type":    "success",
			"notification_title":   translator.T(c, translator.NotificationsAdminSensorsDeleteSuccessTitle),
			"notification_message": translator.T(c, translator.NotificationsAdminSensorsDeleteSuccessMessage),
		}
	}

	nextCursor := ""
	if c.QueryParam("next_cursor") != "" {
		nextCursor = c.QueryParam("next_cursor")
		log.Debug(fmt.Sprintf("next_cursor: %s", nextCursor))
	}
	prevCursor := ""
	if c.QueryParam("previous_cursor") != "" {
		prevCursor = c.QueryParam("previous_cursor")
		log.Debug(fmt.Sprintf("previous_cursor: %s", prevCursor))
	}

	sess, err := h.sessionSvc.Obtain(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	sensors, pagination, err := h.sdk.ConnectedRootsService.SDK.ObtainSensors(ctx, "20", nextCursor, prevCursor, nil, nil, nil, nil, nil)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-sensors-list.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(ctx, c, h.sessionSvc),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonSensorListPageI18N(c)), map[string]interface{}{
			"sensors":    sensors,
			"pagination": pagination,
		}), notifications))
}

func (h *Handlers) GetSensorDeleteHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getDeleteSensorHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(getDeleteSensorHandlerName)

	sensorId := c.Param(sensorIDParam)
	if sensorId == "" {
		err := ferrors.ErrPathParamInvalidValue
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	sess, err := h.sessionSvc.Obtain(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	isAdminTech, err := h.sessionSvc.IsAdminTechnical(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}
	if !isAdminTech {
		return commons.NewErrorS(http.StatusUnauthorized, "forbidden", nil, ferrors.ErrUnauthorized)
	}

	sensor, err := h.sdk.ConnectedRootsService.SDK.ObtainSensor(ctx, sensorId)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-sensors-delete.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(ctx, c, h.sessionSvc),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonSensorDeletePageI18N(c)), map[string]interface{}{
			"sensor": sensor,
		}))
}

func (h *Handlers) PostSensorDeleteHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), postDeleteSensorHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(postDeleteSensorHandlerName)

	sensorId := c.Param(sensorIDParam)
	if sensorId == "" {
		err := ferrors.ErrPathParamInvalidValue
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	isAdminTech, err := h.sessionSvc.IsAdminTechnical(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}
	if !isAdminTech {
		return commons.NewErrorS(http.StatusUnauthorized, "forbidden", nil, ferrors.ErrUnauthorized)
	}

	if err := h.sdk.ConnectedRootsService.SDK.DeleteSensor(ctx, sensorId); err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	if err := h.sessionSvc.SaveMessage(ctx, c, "message", "success"); err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Redirect(http.StatusFound, "/admin/sensors/list")
}
