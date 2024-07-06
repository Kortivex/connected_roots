package role

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
	tracingRoleHandlers = "http-handler.role"

	getCreateRoleHandlerName  = "http-handler.role.get-create-role"
	postCreateRoleHandlerName = "http-handler.role.post-create-role"
	getUpdateRoleHandlerName  = "http-handler.role.get-update-role"
	postUpdateRoleHandlerName = "http-handler.role.post-update-role"
	getViewRoleHandlerName    = "http-handler.role.get-view-role"
	getListRoleHandlerName    = "http-handler.role.get-list-roles"
	getDeleteRoleHandlerName  = "http-handler.role.get-delete-role"
	postDeleteRoleHandlerName = "http-handler.role.post-delete-role"

	roleIDParam = "role_id"
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

func (h *Handlers) GetRoleCreateHandler(c echo.Context) error {
	_, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getCreateRoleHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(getCreateRoleHandlerName)

	sess, err := h.sessionSvc.Obtain(c.Request().Context(), c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-roles-create.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(c),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonRoleCreatePageI18N(c)), map[string]interface{}{}))
}

func (h *Handlers) PostRoleCreateHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), postCreateRoleHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	log := loggerNew.WithTag(postCreateRoleHandlerName)

	sess, err := h.sessionSvc.Obtain(c.Request().Context(), c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	role := &sdk_models.RolesBody{
		Name:        c.FormValue("name"),
		Description: c.FormValue("description"),
		Protected:   c.FormValue("protected") == "on",
	}

	log.Debug(fmt.Sprintf("role: %+v", role))

	_, err = h.sdk.ConnectedRootsService.SDK.SaveRole(ctx, role)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-roles-create.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(c),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonRoleCreatePageI18N(c)), map[string]interface{}{
			"notification_type":    "success",
			"notification_title":   translator.T(c, translator.NotificationsAdminRolesCreateSuccessTitle),
			"notification_message": translator.T(c, translator.NotificationsAdminRolesCreateSuccessMessage),
		}))
}

func (h *Handlers) GetRoleUpdateHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getUpdateRoleHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(getUpdateRoleHandlerName)

	roleID := c.Param(roleIDParam)
	if roleID == "" {
		err := ferrors.ErrPathParamInvalidValue
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	sess, err := h.sessionSvc.Obtain(c.Request().Context(), c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	role, err := h.sdk.ConnectedRootsService.SDK.ObtainRole(ctx, roleID)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-roles-update.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(c),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonRoleUpdatePageI18N(c)), map[string]interface{}{
			"role": role,
		}))
}

func (h *Handlers) PostRoleUpdateHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), postUpdateRoleHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	log := loggerNew.WithTag(postUpdateRoleHandlerName)

	roleID := c.Param(roleIDParam)
	if roleID == "" {
		err := ferrors.ErrPathParamInvalidValue
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	sess, err := h.sessionSvc.Obtain(c.Request().Context(), c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	role := &sdk_models.RolesBody{
		ID:          roleID,
		Name:        c.FormValue("name"),
		Description: c.FormValue("description"),
		Protected:   c.FormValue("protected") == "on",
	}

	log.Debug(fmt.Sprintf("role: %+v", role))

	roleResp, err := h.sdk.ConnectedRootsService.SDK.UpdateRole(ctx, role)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-roles-update.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(c),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonRoleUpdatePageI18N(c)), map[string]interface{}{
			"role":                 roleResp,
			"notification_type":    "success",
			"notification_title":   translator.T(c, translator.NotificationsAdminRolesUpdateSuccessTitle),
			"notification_message": translator.T(c, translator.NotificationsAdminRolesUpdateSuccessMessage),
		}))
}

func (h *Handlers) GetRoleViewHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getViewRoleHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(getViewRoleHandlerName)

	roleID := c.Param(roleIDParam)
	if roleID == "" {
		err := ferrors.ErrPathParamInvalidValue
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	sess, err := h.sessionSvc.Obtain(c.Request().Context(), c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	role, err := h.sdk.ConnectedRootsService.SDK.ObtainRole(ctx, roleID)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-roles-view.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(c),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonRoleViewPageI18N(c)), map[string]interface{}{
			"role": role,
		}))
}

func (h *Handlers) GetRolesListHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getListRoleHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	log := loggerNew.WithTag(getListRoleHandlerName)

	message, err := h.sessionSvc.ObtainMessage(c.Request().Context(), c, "message")
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	log.Debug(fmt.Sprintf("message: %s", message))

	notifications := map[string]interface{}{}
	if len(message) > 0 && message[0] == "success" {
		notifications = map[string]interface{}{
			"notification_type":    "success",
			"notification_title":   translator.T(c, translator.NotificationsAdminRolesDeleteSuccessTitle),
			"notification_message": translator.T(c, translator.NotificationsAdminRolesDeleteSuccessMessage),
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

	sess, err := h.sessionSvc.Obtain(c.Request().Context(), c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	roles, pagination, err := h.sdk.ConnectedRootsService.SDK.ObtainRoles(ctx, "20", nextCursor, prevCursor, nil)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-roles-list.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(c),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonRoleListPageI18N(c)), map[string]interface{}{
			"roles":           roles,
			"protected_roles": h.conf.Roles.Protected,
			"pagination":      pagination,
		}), notifications))
}

func (h *Handlers) GetRoleDeleteHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getDeleteRoleHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(getDeleteRoleHandlerName)

	roleID := c.Param(roleIDParam)
	if roleID == "" {
		err := ferrors.ErrPathParamInvalidValue
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	sess, err := h.sessionSvc.Obtain(c.Request().Context(), c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	role, err := h.sdk.ConnectedRootsService.SDK.ObtainRole(ctx, roleID)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-roles-delete.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(c),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonRoleDeletePageI18N(c)), map[string]interface{}{
			"role": role,
		}))
}

func (h *Handlers) PostRoleDeleteHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), postDeleteRoleHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(postDeleteRoleHandlerName)

	roleID := c.Param(roleIDParam)
	if roleID == "" {
		err := ferrors.ErrPathParamInvalidValue
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	if err := h.sdk.ConnectedRootsService.SDK.DeleteRole(ctx, roleID); err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	if err := h.sessionSvc.SaveMessage(c.Request().Context(), c, "message", "success"); err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Redirect(http.StatusFound, "/admin/roles/list")
}
