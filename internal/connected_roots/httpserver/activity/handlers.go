package activity

import (
	"fmt"
	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/activity"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/internal/connected_roots/httpserver/errors"
	"github.com/Kortivex/connected_roots/internal/connected_roots/user"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/Kortivex/connected_roots/pkg/utils"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
	"net/http"
)

const (
	tracingActivitiesHandlers = "http-handler.activity"

	tracingPostActivitiesHandlers         = "http-handler.activity.post-activity"
	tracingPutActivitiesHandlers          = "http-handler.activity.put-activity"
	tracingGetActivitiesHandlers          = "http-handler.activity.get-activity"
	tracingListActivitiesHandlers         = "http-handler.activity.list-activities"
	tracingDeleteActivitiesHandlers       = "http-handler.activity.delete-activity"
	tracingGetCountActivitiesHandlers     = "http-handler.activity.get-count-activities"
	tracingGetCountUserActivitiesHandlers = "http-handler.activity.get-count-user-activities"

	activityIDParam = "activity_id"
	userIDParam     = "user_id"
)

type ActivitiesHandlers struct {
	gorm   *gorm.DB
	logger *logger.Logger
	conf   *config.Config
	// Services.
	activitySvc *activity.Service
	userSvc     *user.Service
}

func NewActivitiesHandlers(appCtx *connected_roots.Context) *ActivitiesHandlers {
	loggerEmpty := appCtx.Logger.NewEmpty()
	log := loggerEmpty.WithTag(tracingActivitiesHandlers)

	return &ActivitiesHandlers{
		gorm:        appCtx.Gorm,
		logger:      log,
		conf:        appCtx.Conf,
		activitySvc: activity.New(appCtx.Conf, appCtx.Gorm, appCtx.Logger),
		userSvc:     user.New(appCtx.Conf, appCtx.Gorm, appCtx.Logger),
	}
}

func (h *ActivitiesHandlers) PostActivityHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingPostActivitiesHandlers)
	defer span.End()

	userID := c.Param(userIDParam)
	if userID == "" {
		return errors.NewErrorResponse(c, errors.ErrPathParamInvalidValue)
	}
	var userRes *connected_roots.Users
	var err error

	if utils.IsValidEmail(userID) {
		userRes, err = h.userSvc.ObtainFromEmail(ctx, userID)
		if err != nil {
			return errors.NewErrorResponse(c, err)
		}

		return c.JSON(http.StatusOK, userRes)
	}

	userRes, err = h.userSvc.ObtainFromID(ctx, userID)
	if err != nil {
		return errors.NewErrorResponse(c, err)
	}

	activityBody := connected_roots.Activities{}
	if err = c.Bind(&activityBody); err != nil {
		err = fmt.Errorf("%s: %w", tracingPostActivitiesHandlers, errors.ErrInvalidPayload)
		return errors.NewErrorResponse(c, err)
	}

	activityRes, err := h.activitySvc.Save(ctx, &activityBody)
	if err != nil {
		err = fmt.Errorf("%s: %w", tracingPostActivitiesHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated, activityRes)
}

func (h *ActivitiesHandlers) PutActivityHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingPutActivitiesHandlers)
	defer span.End()

	userID := c.Param(userIDParam)
	if userID == "" {
		return errors.NewErrorResponse(c, errors.ErrPathParamInvalidValue)
	}
	var userRes *connected_roots.Users
	var err error

	if utils.IsValidEmail(userID) {
		userRes, err = h.userSvc.ObtainFromEmail(ctx, userID)
		if err != nil {
			return errors.NewErrorResponse(c, err)
		}

		return c.JSON(http.StatusOK, userRes)
	}

	userRes, err = h.userSvc.ObtainFromID(ctx, userID)
	if err != nil {
		return errors.NewErrorResponse(c, err)
	}

	activityID := c.Param(activityIDParam)
	if activityID == "" {
		return errors.NewErrorResponse(c, errors.ErrPathParamInvalidValue)
	}

	activityBody := connected_roots.Activities{}
	if err = c.Bind(&activityBody); err != nil {
		err = fmt.Errorf("%s: %w", tracingPutActivitiesHandlers, errors.ErrInvalidPayload)
		return errors.NewErrorResponse(c, err)
	}

	activityBody.ID = activityID

	activityRes, err := h.activitySvc.Update(ctx, &activityBody)
	if err != nil {
		err = fmt.Errorf("%s: %w", tracingPutActivitiesHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, activityRes)
}

func (h *ActivitiesHandlers) GetActivityHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingGetActivitiesHandlers)
	defer span.End()

	userID := c.Param(userIDParam)
	if userID == "" {
		return errors.NewErrorResponse(c, errors.ErrPathParamInvalidValue)
	}
	var userRes *connected_roots.Users
	var err error

	if utils.IsValidEmail(userID) {
		userRes, err = h.userSvc.ObtainFromEmail(ctx, userID)
		if err != nil {
			return errors.NewErrorResponse(c, err)
		}

		return c.JSON(http.StatusOK, userRes)
	}

	userRes, err = h.userSvc.ObtainFromID(ctx, userID)
	if err != nil {
		return errors.NewErrorResponse(c, err)
	}

	activityID := c.Param(activityIDParam)
	if activityID == "" {
		return errors.NewErrorResponse(c, errors.ErrPathParamInvalidValue)
	}

	activityRes, err := h.activitySvc.Obtain(ctx, activityID)
	if err != nil {
		err = fmt.Errorf("%s: %w", tracingGetActivitiesHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, activityRes)
}

func (h *ActivitiesHandlers) ListActivitiesHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingListActivitiesHandlers)
	defer span.End()

	userID := c.Param(userIDParam)
	if userID == "" {
		return errors.NewErrorResponse(c, errors.ErrPathParamInvalidValue)
	}
	var userRes *connected_roots.Users
	var err error

	if utils.IsValidEmail(userID) {
		userRes, err = h.userSvc.ObtainFromEmail(ctx, userID)
		if err != nil {
			return errors.NewErrorResponse(c, err)
		}

		return c.JSON(http.StatusOK, userRes)
	}

	userRes, err = h.userSvc.ObtainFromID(ctx, userID)
	if err != nil {
		return errors.NewErrorResponse(c, err)
	}

	filters := connected_roots.ActivityPaginationFilters{}
	if err = (&echo.DefaultBinder{}).BindQueryParams(c, &filters); err != nil {
		err = fmt.Errorf("%s: %w", tracingListActivitiesHandlers, errors.ErrQueryParamInvalidValue)
		return errors.NewErrorResponse(c, err)
	}

	activityRes, err := h.activitySvc.ObtainAll(ctx, &filters)
	if err != nil {
		err = fmt.Errorf("%s: %w", tracingListActivitiesHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, activityRes)
}

func (h *ActivitiesHandlers) DeleteActivityHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingDeleteActivitiesHandlers)
	defer span.End()

	userID := c.Param(userIDParam)
	if userID == "" {
		return errors.NewErrorResponse(c, errors.ErrPathParamInvalidValue)
	}
	var userRes *connected_roots.Users
	var err error

	if utils.IsValidEmail(userID) {
		userRes, err = h.userSvc.ObtainFromEmail(ctx, userID)
		if err != nil {
			return errors.NewErrorResponse(c, err)
		}

		return c.JSON(http.StatusOK, userRes)
	}

	userRes, err = h.userSvc.ObtainFromID(ctx, userID)
	if err != nil {
		return errors.NewErrorResponse(c, err)
	}

	activityID := c.Param(activityIDParam)
	if activityID == "" {
		return errors.NewErrorResponse(c, errors.ErrPathParamInvalidValue)
	}

	if err = h.activitySvc.Remove(ctx, activityID); err != nil {
		err = fmt.Errorf("%s: %w", tracingDeleteActivitiesHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *ActivitiesHandlers) GetCountActivitiesHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingGetCountActivitiesHandlers)
	defer span.End()

	total, err := h.activitySvc.CountAll(ctx)
	if err != nil {
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, total)
}

func (h *ActivitiesHandlers) GetCountUserActivitiesHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingGetCountUserActivitiesHandlers)
	defer span.End()

	userID := c.Param(userIDParam)
	if userID == "" {
		return errors.NewErrorResponse(c, errors.ErrPathParamInvalidValue)
	}

	total, err := h.activitySvc.CountAllByUser(ctx, userID)
	if err != nil {
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, total)
}
