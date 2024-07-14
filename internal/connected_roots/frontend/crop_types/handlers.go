package crop_types

import (
	"fmt"
	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/bars"
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/ferrors"
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/i18n/translator"
	sessionServ "github.com/Kortivex/connected_roots/internal/connected_roots/session"
	"github.com/Kortivex/connected_roots/pkg/hashing"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/Kortivex/connected_roots/pkg/logger/commons"
	"github.com/Kortivex/connected_roots/pkg/sdk"
	"github.com/Kortivex/connected_roots/pkg/sdk/sdk_models"
	"github.com/Kortivex/connected_roots/pkg/uploads"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
	"net/http"
)

const (
	tracingCropTypesHandlers = "http-handler.crop-types"

	getCreateCropTypeHandlerName  = "http-handler.crop-types.get-create-crop-type"
	postCreateCropTypeHandlerName = "http-handler.crop-types.post-create-crop-type"
	getUpdateCropTypeHandlerName  = "http-handler.crop-types.get-update-crop-type"
	postUpdateCropTypeHandlerName = "http-handler.crop-types.post-update-crop-type"
	getViewCropTypeHandlerName    = "http-handler.crop-types.get-view-crop-type"
	getListCropTypeHandlerName    = "http-handler.crop-types.get-list-crop-types"
	getDeleteCropTypeHandlerName  = "http-handler.crop-types.get-delete-crop-type"
	postDeleteCropTypeHandlerName = "http-handler.crop-types.post-delete-crop-type"

	cropTypeIDParam = "crop_type_id"
)

type Handlers struct {
	gorm   *gorm.DB
	logger *logger.Logger
	conf   *config.Config
	sdk    *sdk.ExternalAPI
	// Services
	sessionSvc *sessionServ.Service
}

func NewCropTypesHandlers(appCtx *connected_roots.Context) *Handlers {
	loggerEmpty := appCtx.Logger.NewEmpty()
	log := loggerEmpty.WithTag(tracingCropTypesHandlers)

	return &Handlers{
		gorm:       appCtx.Gorm,
		logger:     log,
		conf:       appCtx.Conf,
		sdk:        appCtx.SDK,
		sessionSvc: sessionServ.New(appCtx.Conf, appCtx.Gorm, appCtx.Logger),
	}
}

func (h *Handlers) GetCropTypeCreateHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getCreateCropTypeHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(getCreateCropTypeHandlerName)

	sess, err := h.sessionSvc.Obtain(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-crop-types-create.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(c),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonCropTypeCreatePageI18N(c)), map[string]interface{}{}))
}

func (h *Handlers) PostCropTypeCreateHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), postCreateCropTypeHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	log := loggerNew.WithTag(postCreateCropTypeHandlerName)

	sess, err := h.sessionSvc.Obtain(ctx, c)
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

	pathImage := "images/crop_types/" + fileName
	if err = uploads.SaveUploadedImage(file, pathImage, 800, 800); err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	cropType := &sdk_models.CropTypesBody{
		Name:           c.FormValue("name"),
		ScientificName: c.FormValue("scientific-name"),
		LifeCycle:      c.FormValue("life-cycle"),
		PlantingSeason: c.FormValue("planting-season"),
		HarvestSeason:  c.FormValue("harvest-season"),
		Irrigation:     c.FormValue("irrigation"),
		Description:    c.FormValue("description"),
		ImageURL:       pathImage,
	}

	log.Debug(fmt.Sprintf("crop_type: %+v", cropType))

	_, err = h.sdk.ConnectedRootsService.SDK.SaveCropType(ctx, cropType)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-crop-types-create.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(c),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonCropTypeCreatePageI18N(c)), map[string]interface{}{
			"notification_type":    "success",
			"notification_title":   translator.T(c, translator.NotificationsAdminCropTypesCreateSuccessTitle),
			"notification_message": translator.T(c, translator.NotificationsAdminCropTypesCreateSuccessMessage),
		}))
}

func (h *Handlers) GetCropTypeViewHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getViewCropTypeHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(getViewCropTypeHandlerName)

	cropTypeId := c.Param(cropTypeIDParam)
	if cropTypeId == "" {
		err := ferrors.ErrPathParamInvalidValue
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	sess, err := h.sessionSvc.Obtain(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	cropType, err := h.sdk.ConnectedRootsService.SDK.ObtainCropType(ctx, cropTypeId)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-crop-types-view.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(c),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonCropTypeViewPageI18N(c)), map[string]interface{}{
			"crop_type": cropType,
		}))
}

func (h *Handlers) GetCropTypesListHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getListCropTypeHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	log := loggerNew.WithTag(getListCropTypeHandlerName)

	message, err := h.sessionSvc.ObtainMessage(ctx, c, "message")
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	log.Debug(fmt.Sprintf("message: %s", message))

	notifications := map[string]interface{}{}
	if len(message) > 0 && message[0] == "success" {
		notifications = map[string]interface{}{
			"notification_type":    "success",
			"notification_title":   translator.T(c, translator.NotificationsAdminCropTypesDeleteSuccessTitle),
			"notification_message": translator.T(c, translator.NotificationsAdminCropTypesDeleteSuccessMessage),
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

	cropTypes, pagination, err := h.sdk.ConnectedRootsService.SDK.ObtainCropTypes(ctx, "20", nextCursor, prevCursor, nil, nil, nil, nil)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-crop-types-list.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(c),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonCropTypeListPageI18N(c)), map[string]interface{}{
			"crop_types": cropTypes,
			"pagination": pagination,
		}), notifications))
}

func (h *Handlers) GetCropTypeDeleteHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getDeleteCropTypeHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(getDeleteCropTypeHandlerName)

	cropTypeId := c.Param(cropTypeIDParam)
	if cropTypeId == "" {
		err := ferrors.ErrPathParamInvalidValue
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	sess, err := h.sessionSvc.Obtain(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	cropType, err := h.sdk.ConnectedRootsService.SDK.ObtainCropType(ctx, cropTypeId)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-crop-types-delete.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(c),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonCropTypeDeletePageI18N(c)), map[string]interface{}{
			"crop_type": cropType,
		}))
}

func (h *Handlers) PostCropTypeDeleteHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), postDeleteCropTypeHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(postDeleteCropTypeHandlerName)

	cropTypeId := c.Param(cropTypeIDParam)
	if cropTypeId == "" {
		err := ferrors.ErrPathParamInvalidValue
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	cropType, err := h.sdk.ConnectedRootsService.SDK.ObtainCropType(ctx, cropTypeId)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	if err = h.sdk.ConnectedRootsService.SDK.DeleteCropType(ctx, cropTypeId); err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	if err = uploads.DeleteUploadedImage(cropType.ImageURL); err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	if err = h.sessionSvc.SaveMessage(ctx, c, "message", "success"); err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Redirect(http.StatusFound, "/admin/crop-types/list")
}
