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

	tracingPostRolesHandlers     = "http-handler.role.post-role"
	tracingPutRolesHandlers      = "http-handler.role.put-role"
	tracingGetRolesHandlers      = "http-handler.role.get-role"
	tracingListRolesHandlers     = "http-handler.role.list-roles"
	tracingDeleteRolesHandlers   = "http-handler.role.delete-role"
	tracingGetCountRolesHandlers = "http-handler.role.get-count-roles"

	roleIDParam = "role_id"
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

func (h *RolesHandlers) PutRolesHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingPutRolesHandlers)
	defer span.End()

	roleID := c.Param(roleIDParam)
	if roleID == "" {
		return errors.NewErrorResponse(c, errors.ErrPathParamInvalidValue)
	}

	roleBody := connected_roots.Roles{}
	if err := c.Bind(&roleBody); err != nil {
		err = fmt.Errorf("%s: %w", tracingPutRolesHandlers, errors.ErrInvalidPayload)
		return errors.NewErrorResponse(c, err)
	}

	roleBody.ID = roleID

	rolesRes, err := h.roleSvc.Update(ctx, &roleBody)
	if err != nil {
		err = fmt.Errorf("%s: %w", tracingPutRolesHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, rolesRes)
}

func (h *RolesHandlers) GetRolesHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingGetRolesHandlers)
	defer span.End()

	roleID := c.Param(roleIDParam)
	if roleID == "" {
		return errors.NewErrorResponse(c, errors.ErrPathParamInvalidValue)
	}

	rolesRes, err := h.roleSvc.Obtain(ctx, roleID)
	if err != nil {
		err = fmt.Errorf("%s: %w", tracingGetRolesHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, rolesRes)
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

func (h *RolesHandlers) DeleteRolesHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingDeleteRolesHandlers)
	defer span.End()

	roleID := c.Param(roleIDParam)
	if roleID == "" {
		return errors.NewErrorResponse(c, errors.ErrPathParamInvalidValue)
	}

	if err := h.roleSvc.Remove(ctx, roleID); err != nil {
		err = fmt.Errorf("%s: %w", tracingDeleteRolesHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *RolesHandlers) GetCountRolesHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingGetCountRolesHandlers)
	defer span.End()

	total, err := h.roleSvc.CountAll(ctx)
	if err != nil {
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, total)
}
