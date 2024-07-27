package home

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/bars"
	sessionServ "github.com/Kortivex/connected_roots/internal/connected_roots/session"
	"github.com/Kortivex/connected_roots/pkg/logger/commons"
	"github.com/Kortivex/connected_roots/pkg/sdk"
	"net/http"

	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/i18n/translator"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

const (
	tracingHomeHandlers = "http-handler.home"

	getHomeHandlerName = "http-handler.home.get-login"
)

type Handlers struct {
	gorm   *gorm.DB
	logger *logger.Logger
	conf   *config.Config
	sdk    *sdk.ExternalAPI
	// Services
	sessionSvc *sessionServ.Service
}

func NewHomeHandlers(appCtx *connected_roots.Context) *Handlers {
	loggerEmpty := appCtx.Logger.NewEmpty()
	log := loggerEmpty.WithTag(tracingHomeHandlers)

	return &Handlers{
		gorm:       appCtx.Gorm,
		logger:     log,
		conf:       appCtx.Conf,
		sdk:        appCtx.SDK,
		sessionSvc: sessionServ.New(appCtx.Conf, appCtx.Gorm, appCtx.Logger),
	}
}

func (h *Handlers) GetHomeHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getHomeHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	log := loggerNew.WithTag(getHomeHandlerName)

	log.Debug("home")

	sess, err := h.sessionSvc.Obtain(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	isAdminTech, _ := h.sessionSvc.IsAdminTechnical(ctx, c)
	if isAdminTech {
		totalSessions, err := h.sessionSvc.CountAll(ctx)
		if err != nil {
			return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
		}
		totalUsers, err := h.sdk.ConnectedRootsService.SDK.CountUsers(ctx)
		if err != nil {
			return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
		}
		totalOrchards, err := h.sdk.ConnectedRootsService.SDK.CountOrchards(ctx)
		if err != nil {
			return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
		}
		totalSensors, err := h.sdk.ConnectedRootsService.SDK.CountSensors(ctx)
		if err != nil {
			return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
		}
		totalCropTypes, err := h.sdk.ConnectedRootsService.SDK.CountCropTypes(ctx)
		if err != nil {
			return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
		}
		totalActivities, err := h.sdk.ConnectedRootsService.SDK.CountActivities(ctx)
		if err != nil {
			return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
		}

		return c.Render(http.StatusOK, "admin-home.gohtml",
			translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(ctx, c, h.sessionSvc),
				bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
				CommonHomePageI18N(c)), map[string]interface{}{
				"total_sessions":   totalSessions,
				"total_users":      totalUsers,
				"total_orchards":   totalOrchards,
				"total_sensors":    totalSensors,
				"total_crop_types": totalCropTypes,
				"total_activities": totalActivities,
			}))
	}

	totalOrchards, err := h.sdk.ConnectedRootsService.SDK.CountUserOrchards(ctx, sess.UserID)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}
	totalSensors, err := h.sdk.ConnectedRootsService.SDK.CountUserSensors(ctx, sess.UserID)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}
	totalActivities, err := h.sdk.ConnectedRootsService.SDK.CountUserActivities(ctx, sess.UserID)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "user-home.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(ctx, c, h.sessionSvc),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonHomePageI18N(c)), map[string]interface{}{
			"total_orchards":   totalOrchards,
			"total_sensors":    totalSensors,
			"total_activities": totalActivities,
		}))
}
