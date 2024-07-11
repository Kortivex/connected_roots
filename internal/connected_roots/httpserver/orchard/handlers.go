package orchard

import (
	"fmt"
	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/internal/connected_roots/httpserver/errors"
	"github.com/Kortivex/connected_roots/internal/connected_roots/orchard"
	"github.com/Kortivex/connected_roots/internal/connected_roots/user"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/Kortivex/connected_roots/pkg/utils"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
	"net/http"
)

const (
	tracingOrchardsHandlers = "http-handler.orchard"

	tracingPostOrchardsHandlers     = "http-handler.orchard.post-orchard"
	tracingPutOrchardsHandlers      = "http-handler.orchard.put-orchard"
	tracingGetOrchardsHandlers      = "http-handler.orchard.get-orchard"
	tracingListOrchardsHandlers     = "http-handler.orchard.list-orchards"
	tracingDeleteOrchardsHandlers   = "http-handler.orchard.delete-orchard"
	tracingGetUserOrchardHandlers   = "http-handler.orchard.get-user-orchard"
	tracingListUserOrchardsHandlers = "http-handler.orchard.list-user-orchards"

	orchardIDParam = "orchard_id"
	userIDParam    = "user_id"
)

type OrchardsHandlers struct {
	gorm   *gorm.DB
	logger *logger.Logger
	conf   *config.Config
	// Services.
	orchardSvc *orchard.Service
	userSvc    *user.Service
}

func NewOrchardsHandlers(appCtx *connected_roots.Context) *OrchardsHandlers {
	loggerEmpty := appCtx.Logger.NewEmpty()
	log := loggerEmpty.WithTag(tracingOrchardsHandlers)

	return &OrchardsHandlers{
		gorm:       appCtx.Gorm,
		logger:     log,
		conf:       appCtx.Conf,
		orchardSvc: orchard.New(appCtx.Conf, appCtx.Gorm, appCtx.Logger),
		userSvc:    user.New(appCtx.Conf, appCtx.Gorm, appCtx.Logger),
	}
}

func (h *OrchardsHandlers) PostOrchardHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingPostOrchardsHandlers)
	defer span.End()

	orchardBody := connected_roots.Orchards{}
	if err := c.Bind(&orchardBody); err != nil {
		err = fmt.Errorf("%s: %w", tracingPostOrchardsHandlers, errors.ErrInvalidPayload)
		return errors.NewErrorResponse(c, err)
	}

	orchardRes, err := h.orchardSvc.Save(ctx, &orchardBody)
	if err != nil {
		err = fmt.Errorf("%s: %w", tracingPostOrchardsHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated, orchardRes)
}

func (h *OrchardsHandlers) PutOrchardHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingPutOrchardsHandlers)
	defer span.End()

	orchardID := c.Param(orchardIDParam)
	if orchardID == "" {
		return errors.NewErrorResponse(c, errors.ErrPathParamInvalidValue)
	}

	orchardBody := connected_roots.Orchards{}
	if err := c.Bind(&orchardBody); err != nil {
		err = fmt.Errorf("%s: %w", tracingPutOrchardsHandlers, errors.ErrInvalidPayload)
		return errors.NewErrorResponse(c, err)
	}

	orchardBody.ID = orchardID

	orchardRes, err := h.orchardSvc.Update(ctx, &orchardBody)
	if err != nil {
		err = fmt.Errorf("%s: %w", tracingPutOrchardsHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, orchardRes)
}

func (h *OrchardsHandlers) GetOrchardHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingGetOrchardsHandlers)
	defer span.End()

	orchardID := c.Param(orchardIDParam)
	if orchardID == "" {
		return errors.NewErrorResponse(c, errors.ErrPathParamInvalidValue)
	}

	orchardRes, err := h.orchardSvc.Obtain(ctx, orchardID)
	if err != nil {
		err = fmt.Errorf("%s: %w", tracingGetOrchardsHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, orchardRes)
}

func (h *OrchardsHandlers) ListOrchardsHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingListOrchardsHandlers)
	defer span.End()

	filters := connected_roots.OrchardPaginationFilters{}
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &filters); err != nil {
		err = fmt.Errorf("%s: %w", tracingListOrchardsHandlers, errors.ErrQueryParamInvalidValue)
		return errors.NewErrorResponse(c, err)
	}

	orchardRes, err := h.orchardSvc.ObtainAll(ctx, &filters)
	if err != nil {
		err = fmt.Errorf("%s: %w", tracingListOrchardsHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, orchardRes)
}

func (h *OrchardsHandlers) DeleteOrchardHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingDeleteOrchardsHandlers)
	defer span.End()

	orchardID := c.Param(orchardIDParam)
	if orchardID == "" {
		return errors.NewErrorResponse(c, errors.ErrPathParamInvalidValue)
	}

	if err := h.orchardSvc.Remove(ctx, orchardID); err != nil {
		err = fmt.Errorf("%s: %w", tracingDeleteOrchardsHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *OrchardsHandlers) GetUserOrchardHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingGetUserOrchardHandlers)
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

	orchardID := c.Param(orchardIDParam)
	if orchardID == "" {
		return errors.NewErrorResponse(c, errors.ErrPathParamInvalidValue)
	}

	orchardRes, err := h.orchardSvc.Obtain(ctx, orchardID)
	if err != nil {
		err = fmt.Errorf("%s: %w", tracingGetUserOrchardHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, orchardRes)
}

func (h *OrchardsHandlers) ListUserOrchardsHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingListUserOrchardsHandlers)
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

	filters := connected_roots.OrchardPaginationFilters{}
	if err = (&echo.DefaultBinder{}).BindQueryParams(c, &filters); err != nil {
		err = fmt.Errorf("%s: %w", tracingListUserOrchardsHandlers, errors.ErrQueryParamInvalidValue)
		return errors.NewErrorResponse(c, err)
	}

	filters.UserID = []string{userRes.ID}

	orchardRes, err := h.orchardSvc.ObtainAll(ctx, &filters)
	if err != nil {
		err = fmt.Errorf("%s: %w", tracingListUserOrchardsHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, orchardRes)
}
