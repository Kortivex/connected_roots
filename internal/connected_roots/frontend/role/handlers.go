package role

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/bars"
	sessionServ "github.com/Kortivex/connected_roots/internal/connected_roots/session"
	"github.com/Kortivex/connected_roots/pkg/logger/commons"
	"net/http"

	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/i18n/translator"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/Kortivex/connected_roots/pkg/sdk"
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

func (h *Handlers) ListRolesHandler(c echo.Context) error {
	_, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), listRoleHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(listRoleHandlerName)

	sess, err := h.sessionSvc.Obtain(c.Request().Context(), c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusInternalServerError, "admin-roles-list.gohtml", translator.AddDataKeys(
		translator.AddDataKeys(bars.CommonNavBarI18N(c), bars.CommonTopBarI18N(c, sess.Name, sess.Surname)), map[string]interface{}{
			"site_title": translator.T(c, translator.PagesCommonSiteTitle),
			"title":      translator.T(c, translator.PagesAdminRolesTitle),
			"active":     "roles-management",
		}))
}
