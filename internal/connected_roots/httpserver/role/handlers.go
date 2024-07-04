package role

import (
	"fmt"
	"github.com/Kortivex/connected_roots/internal/connected_roots/httpserver/errors"
	"github.com/Kortivex/connected_roots/internal/connected_roots/role"
	"net/http"

	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

const (
	tracingRolesHandlers = "http-handler.role"

	tracingPostRolesHandlers = "http-handler.role.post-role"
	tracingListRolesHandlers = "http-handler.role.list-roles"
)

type RolesHandlers struct {
	gorm   *gorm.DB
	logger *logger.Logger
	conf   *config.Config
	// Services.
	roleSvc *role.Service
}

func NewRolesHandlers(appCtx *connected_roots.Context) *RolesHandlers {
	loggerEmpty := appCtx.Logger.NewEmpty()
	log := loggerEmpty.WithTag(tracingRolesHandlers)

	return &RolesHandlers{
		gorm:    appCtx.Gorm,
		logger:  log,
		conf:    appCtx.Conf,
		roleSvc: role.New(appCtx.Conf, appCtx.Gorm, appCtx.Logger),
	}
}

func (h *RolesHandlers) PostRolesHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingPostRolesHandlers)
	defer span.End()

	roleBody := connected_roots.Roles{}
	if err := c.Bind(&roleBody); err != nil {
		err = fmt.Errorf("%s: %w", tracingPostRolesHandlers, errors.ErrInvalidPayload)
		return errors.NewErrorResponse(c, err)
	}

	rolesRes, err := h.roleSvc.Save(ctx, &roleBody)
	if err != nil {
		err = fmt.Errorf("%s: %w", tracingPostRolesHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated, rolesRes)
}

func (h *RolesHandlers) ListRolesHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingListRolesHandlers)
	defer span.End()

	filters := connected_roots.RolePaginationFilters{}
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &filters); err != nil {
		err = fmt.Errorf("%s: %w", tracingListRolesHandlers, errors.ErrQueryParamInvalidValue)
		return errors.NewErrorResponse(c, err)
	}

	rolesRes, err := h.roleSvc.ObtainAll(ctx, &filters)
	if err != nil {
		err = fmt.Errorf("%s: %w", tracingListRolesHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, rolesRes)
}
