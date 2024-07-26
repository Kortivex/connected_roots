package activity

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
)

const (
	tracingActivityHandlers = "http-handler.activity"

	getCreateActivityHandlerName  = "http-handler.activity.get-create-activity"
	postCreateActivityHandlerName = "http-handler.activity.post-create-activity"
	getUpdateActivityHandlerName  = "http-handler.activity.get-update-activity"
	postUpdateActivityHandlerName = "http-handler.activity.post-update-activity"
	getViewActivityHandlerName    = "http-handler.activity.get-view-activity"
	getListActivityHandlerName    = "http-handler.activity.get-list-activity"
	getDeleteActivityHandlerName  = "http-handler.activity.get-delete-activity"
	postDeleteActivityHandlerName = "http-handler.activity.post-delete-activity"

	activityIDParam = "activity_id"
)

type Handlers struct {
	gorm   *gorm.DB
	logger *logger.Logger
	conf   *config.Config
	sdk    *sdk.ExternalAPI
	// Services
	sessionSvc *sessionServ.Service
}

func NewActivitiesHandlers(appCtx *connected_roots.Context) *Handlers {
	loggerEmpty := appCtx.Logger.NewEmpty()
	log := loggerEmpty.WithTag(tracingActivityHandlers)

	return &Handlers{
		gorm:       appCtx.Gorm,
		logger:     log,
		conf:       appCtx.Conf,
		sdk:        appCtx.SDK,
		sessionSvc: sessionServ.New(appCtx.Conf, appCtx.Gorm, appCtx.Logger),
	}
}

func (h *Handlers) GetActivityCreateHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getCreateActivityHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(getCreateActivityHandlerName)

	sess, err := h.sessionSvc.Obtain(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	isUser, err := h.sessionSvc.IsUser(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}
	if !isUser {
		return commons.NewErrorS(http.StatusUnauthorized, "forbidden", nil, ferrors.ErrUnauthorized)
	}

	orchards, _, err := h.sdk.ConnectedRootsService.SDK.ObtainOrchards(ctx, "100000", "", "", nil, nil, nil)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "user-activities-create.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(ctx, c, h.sessionSvc),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonActivityCreatePageI18N(c)), map[string]interface{}{
			"orchards": orchards,
		}))
}

func (h *Handlers) PostActivityCreateHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), postCreateActivityHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	log := loggerNew.WithTag(postCreateActivityHandlerName)

	sess, err := h.sessionSvc.Obtain(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	isUser, err := h.sessionSvc.IsUser(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}
	if !isUser {
		return commons.NewErrorS(http.StatusUnauthorized, "forbidden", nil, ferrors.ErrUnauthorized)
	}

	activity := &sdk_models.ActivitiesBody{
		Name:        c.FormValue("name"),
		Description: c.FormValue("description"),
		OrchardID:   c.FormValue("orchard-id"),
	}

	log.Debug(fmt.Sprintf("activity: %+v", activity))

	_, err = h.sdk.ConnectedRootsService.SDK.SaveActivity(ctx, sess.UserID, activity)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "user-activities-create.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(ctx, c, h.sessionSvc),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonActivityCreatePageI18N(c)), map[string]interface{}{
			"notification_type":    "success",
			"notification_title":   translator.T(c, translator.NotificationsUserActivitiesCreateSuccessTitle),
			"notification_message": translator.T(c, translator.NotificationsUserActivitiesCreateSuccessMessage),
		}))
}

func (h *Handlers) GetActivityUpdateHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getUpdateActivityHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(getUpdateActivityHandlerName)

	activityId := c.Param(activityIDParam)
	if activityId == "" {
		err := ferrors.ErrPathParamInvalidValue
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	sess, err := h.sessionSvc.Obtain(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	isUser, err := h.sessionSvc.IsUser(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}
	if !isUser {
		return commons.NewErrorS(http.StatusUnauthorized, "forbidden", nil, ferrors.ErrUnauthorized)
	}

	activity, err := h.sdk.ConnectedRootsService.SDK.ObtainActivity(ctx, sess.UserID, activityId)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	orchards, _, err := h.sdk.ConnectedRootsService.SDK.ObtainOrchards(ctx, "100000", "", "", nil, nil, nil)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "user-activities-update.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(ctx, c, h.sessionSvc),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonActivityUpdatePageI18N(c)), map[string]interface{}{
			"activity": activity,
			"orchards": orchards,
		}))
}

func (h *Handlers) PostActivityUpdateHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), postUpdateActivityHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	log := loggerNew.WithTag(postUpdateActivityHandlerName)

	activityId := c.Param(activityIDParam)
	if activityId == "" {
		err := ferrors.ErrPathParamInvalidValue
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	sess, err := h.sessionSvc.Obtain(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	isUser, err := h.sessionSvc.IsUser(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}
	if !isUser {
		return commons.NewErrorS(http.StatusUnauthorized, "forbidden", nil, ferrors.ErrUnauthorized)
	}

	activity, err := h.sdk.ConnectedRootsService.SDK.ObtainActivity(ctx, sess.UserID, activityId)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	activity.ID = activityId
	activity.Name = c.FormValue("name")
	activity.Description = c.FormValue("description")
	activity.OrchardID = c.FormValue("orchard-id")

	log.Debug(fmt.Sprintf("activity: %+v", activity))

	if _, err = h.sdk.ConnectedRootsService.SDK.UpdateActivity(ctx, sess.UserID, activity.ToActivityBody()); err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	orchards, _, err := h.sdk.ConnectedRootsService.SDK.ObtainOrchards(ctx, "100000", "", "", nil, nil, nil)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "user-activities-update.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(ctx, c, h.sessionSvc),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonActivityUpdatePageI18N(c)), map[string]interface{}{
			"activity":             activity,
			"orchards":             orchards,
			"notification_type":    "success",
			"notification_title":   translator.T(c, translator.NotificationsUserActivitiesUpdateSuccessTitle),
			"notification_message": translator.T(c, translator.NotificationsUserActivitiesUpdateSuccessMessage),
		}))
}

func (h *Handlers) GetActivityViewHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getViewActivityHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(getViewActivityHandlerName)

	activityId := c.Param(activityIDParam)
	if activityId == "" {
		err := ferrors.ErrPathParamInvalidValue
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	sess, err := h.sessionSvc.Obtain(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	isUser, err := h.sessionSvc.IsUser(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}
	if !isUser {
		return commons.NewErrorS(http.StatusUnauthorized, "forbidden", nil, ferrors.ErrUnauthorized)
	}

	activity, err := h.sdk.ConnectedRootsService.SDK.ObtainActivity(ctx, sess.UserID, activityId)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "user-activities-view.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(ctx, c, h.sessionSvc),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonActivityViewPageI18N(c)), map[string]interface{}{
			"activity": activity,
		}))
}

func (h *Handlers) GetActivitiesListHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getListActivityHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	log := loggerNew.WithTag(getListActivityHandlerName)

	message, err := h.sessionSvc.ObtainMessage(ctx, c, "message")
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	isUser, err := h.sessionSvc.IsUser(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}
	if !isUser {
		return commons.NewErrorS(http.StatusUnauthorized, "forbidden", nil, ferrors.ErrUnauthorized)
	}

	log.Debug(fmt.Sprintf("message: %s", message))

	notifications := map[string]interface{}{}
	if len(message) > 0 && message[0] == "success" {
		notifications = map[string]interface{}{
			"notification_type":    "success",
			"notification_title":   translator.T(c, translator.NotificationsUserActivitiesDeleteSuccessTitle),
			"notification_message": translator.T(c, translator.NotificationsUserActivitiesDeleteSuccessMessage),
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

	activities, pagination, err := h.sdk.ConnectedRootsService.SDK.ObtainActivities(ctx, sess.UserID, "20", nextCursor, prevCursor, nil, nil)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "user-activities-list.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(ctx, c, h.sessionSvc),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonActivityListPageI18N(c)), map[string]interface{}{
			"activities": activities,
			"pagination": pagination,
		}), notifications))
}

func (h *Handlers) GetActivityDeleteHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getDeleteActivityHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(getDeleteActivityHandlerName)

	activityId := c.Param(activityIDParam)
	if activityId == "" {
		err := ferrors.ErrPathParamInvalidValue
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	sess, err := h.sessionSvc.Obtain(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	isUser, err := h.sessionSvc.IsUser(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}
	if !isUser {
		return commons.NewErrorS(http.StatusUnauthorized, "forbidden", nil, ferrors.ErrUnauthorized)
	}

	activity, err := h.sdk.ConnectedRootsService.SDK.ObtainActivity(ctx, sess.UserID, activityId)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "user-activities-delete.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(ctx, c, h.sessionSvc),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonActivityDeletePageI18N(c)), map[string]interface{}{
			"activity": activity,
		}))
}

func (h *Handlers) PostActivityDeleteHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), postDeleteActivityHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(postDeleteActivityHandlerName)

	activityId := c.Param(activityIDParam)
	if activityId == "" {
		err := ferrors.ErrPathParamInvalidValue
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	isUser, err := h.sessionSvc.IsUser(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}
	if !isUser {
		return commons.NewErrorS(http.StatusUnauthorized, "forbidden", nil, ferrors.ErrUnauthorized)
	}

	sess, err := h.sessionSvc.Obtain(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	if err = h.sdk.ConnectedRootsService.SDK.DeleteActivity(ctx, sess.UserID, activityId); err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	if err = h.sessionSvc.SaveMessage(ctx, c, "message", "success"); err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Redirect(http.StatusFound, "/users/activities/list")
}
