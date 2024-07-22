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
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
	"net/http"
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

/*func (h *Handlers) GetSensorCreateHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getCreateOrchardHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(getCreateOrchardHandlerName)

	sess, err := h.sessionSvc.Obtain(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	isAdmin, err := h.sessionSvc.IsAdmin(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}
	if !isAdmin {
		return commons.NewErrorS(http.StatusUnauthorized, "forbidden", nil, ferrors.ErrUnauthorized)
	}

	users, _, err := h.sdk.ConnectedRootsService.SDK.ObtainUsers(ctx, "10000", "", "", nil, nil, nil)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	cropTypes, _, err := h.sdk.ConnectedRootsService.SDK.ObtainCropTypes(ctx, "10000", "", "", nil, nil, nil, nil)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-orchards-create.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(ctx, c, h.sessionSvc),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonOrchardCreatePageI18N(c)), map[string]interface{}{
			"users":      users,
			"crop_types": cropTypes,
		}))
}

func (h *Handlers) PostSensorCreateHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), postCreateOrchardHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	log := loggerNew.WithTag(postCreateOrchardHandlerName)

	sess, err := h.sessionSvc.Obtain(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	isAdmin, err := h.sessionSvc.IsAdmin(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}
	if !isAdmin {
		return commons.NewErrorS(http.StatusUnauthorized, "forbidden", nil, ferrors.ErrUnauthorized)
	}

	size, err := strconv.ParseFloat(c.FormValue("size"), 64)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	file, err := c.FormFile("file")
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	fileName, err := hashing.GenUniqueFileName(file.Filename)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	pathImage := "images/orchards/" + fileName
	if err = uploads.SaveUploadedImage(file, pathImage, 800, 800); err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	orchard := &sdk_models.OrchardsBody{
		Name:       c.FormValue("name"),
		Location:   c.FormValue("location"),
		Size:       size,
		Soil:       c.FormValue("soil"),
		Fertilizer: c.FormValue("fertilizer"),
		Composting: c.FormValue("composting"),
		ImageURL:   pathImage,
		UserID:     c.FormValue("user-id"),
		CropTypeID: c.FormValue("crop-type-id"),
	}

	log.Debug(fmt.Sprintf("orchard: %+v", orchard))

	_, err = h.sdk.ConnectedRootsService.SDK.SaveOrchard(ctx, orchard)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-orchards-create.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(ctx, c, h.sessionSvc),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonOrchardCreatePageI18N(c)), map[string]interface{}{
			"notification_type":    "success",
			"notification_title":   translator.T(c, translator.NotificationsAdminOrchardsCreateSuccessTitle),
			"notification_message": translator.T(c, translator.NotificationsAdminOrchardsCreateSuccessMessage),
		}))
}

func (h *Handlers) GetSensorUpdateHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getUpdateOrchardHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(getUpdateOrchardHandlerName)

	orchardID := c.Param(orchardIDParam)
	if orchardID == "" {
		err := ferrors.ErrPathParamInvalidValue
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	sess, err := h.sessionSvc.Obtain(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	isAdmin, err := h.sessionSvc.IsAdmin(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}
	if !isAdmin {
		return commons.NewErrorS(http.StatusUnauthorized, "forbidden", nil, ferrors.ErrUnauthorized)
	}

	users, _, err := h.sdk.ConnectedRootsService.SDK.ObtainUsers(ctx, "10000", "", "", nil, nil, nil)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	cropTypes, _, err := h.sdk.ConnectedRootsService.SDK.ObtainCropTypes(ctx, "10000", "", "", nil, nil, nil, nil)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	orchard, err := h.sdk.ConnectedRootsService.SDK.ObtainOrchard(ctx, orchardID)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-orchards-update.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(ctx, c, h.sessionSvc),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonOrchardUpdatePageI18N(c)), map[string]interface{}{
			"users":      users,
			"crop_types": cropTypes,
			"orchard":    orchard,
		}))
}

func (h *Handlers) PostSensorUpdateHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), postUpdateOrchardHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	log := loggerNew.WithTag(postUpdateOrchardHandlerName)

	orchardID := c.Param(orchardIDParam)
	if orchardID == "" {
		err := ferrors.ErrPathParamInvalidValue
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	sess, err := h.sessionSvc.Obtain(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	isAdmin, err := h.sessionSvc.IsAdmin(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}
	if !isAdmin {
		return commons.NewErrorS(http.StatusUnauthorized, "forbidden", nil, ferrors.ErrUnauthorized)
	}

	file, err := c.FormFile("file")
	if err != nil {
		if !errors.Is(err, http.ErrMissingFile) {
			return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
		}
	}

	pathImage := ""
	if file != nil {
		fileName, err := hashing.GenUniqueFileName(file.Filename)
		if err != nil {
			return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
		}
		pathImage = "images/orchards/" + fileName
	}

	orchard, err := h.sdk.ConnectedRootsService.SDK.ObtainOrchard(ctx, orchardID)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	oldPathImage := orchard.ImageURL

	size, err := strconv.ParseFloat(c.FormValue("size"), 64)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	orchard.Name = c.FormValue("name")
	orchard.Location = c.FormValue("location")
	orchard.Size = size
	orchard.Soil = c.FormValue("soil")
	orchard.Fertilizer = c.FormValue("fertilizer")
	orchard.Composting = c.FormValue("composting")
	orchard.UserID = c.FormValue("user-id")
	orchard.CropTypeID = c.FormValue("crop-type-id")
	if pathImage != "" {
		orchard.ImageURL = pathImage
	}

	log.Debug(fmt.Sprintf("orchard: %+v", orchard))

	users, _, err := h.sdk.ConnectedRootsService.SDK.ObtainUsers(ctx, "10000", "", "", nil, nil, nil)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	cropTypes, _, err := h.sdk.ConnectedRootsService.SDK.ObtainCropTypes(ctx, "10000", "", "", nil, nil, nil, nil)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	_, err = h.sdk.ConnectedRootsService.SDK.UpdateOrchard(ctx, orchard.ToOrchardBody())
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	if file != nil {
		if err = uploads.SaveUploadedImage(file, pathImage, 800, 800); err != nil {
			return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
		}

		if err = uploads.DeleteUploadedImage(oldPathImage); err != nil {
			return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
		}
	}

	return c.Render(http.StatusOK, "admin-orchards-update.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(ctx, c, h.sessionSvc),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonOrchardUpdatePageI18N(c)), map[string]interface{}{
			"users":                users,
			"crop_types":           cropTypes,
			"orchard":              orchard,
			"notification_type":    "success",
			"notification_title":   translator.T(c, translator.NotificationsAdminOrchardsUpdateSuccessTitle),
			"notification_message": translator.T(c, translator.NotificationsAdminOrchardsUpdateSuccessMessage),
		}))
}
*/

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
