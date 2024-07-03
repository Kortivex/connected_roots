package role

import (
	"fmt"
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
	tracingRoleHandlers = "http-handler.role"

	listRoleHandlerName = "http-handler.role.list-roles"
)

type Handlers struct {
	gorm   *gorm.DB
	logger *logger.Logger
	conf   *config.Config
	sdk    *sdk.ExternalAPI
	// Services
	sessionSvc *sessionServ.Service
}

func NewRolesHandlers(appCtx *connected_roots.Context) *Handlers {
	loggerEmpty := appCtx.Logger.NewEmpty()
	log := loggerEmpty.WithTag(tracingRoleHandlers)

	return &Handlers{
		gorm:       appCtx.Gorm,
		logger:     log,
		conf:       appCtx.Conf,
		sdk:        appCtx.SDK,
		sessionSvc: sessionServ.New(appCtx.Conf, appCtx.Gorm, appCtx.Logger),
	}
}

func (h *Handlers) GetRolesListHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), listRoleHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	log := loggerNew.WithTag(listRoleHandlerName)

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

	sess, err := h.sessionSvc.Obtain(c.Request().Context(), c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	roles, pagination, err := h.sdk.ConnectedRootsService.SDK.ObtainRoles(ctx, "20", nextCursor, prevCursor, nil)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-roles-list.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(c),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonRoleListPageI18N(c)), map[string]interface{}{
			"active":     "roles-management",
			"roles":      roles,
			"pagination": pagination,
		}))
}
