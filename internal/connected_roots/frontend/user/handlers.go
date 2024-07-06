package user

import (
	"fmt"
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/bars"
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/ferrors"
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
	tracingUserHandlers = "http-handler.user"

	getViewUserHandlerName     = "http-handler.user.get-view-user"
	getListUsersHandlerName    = "http-handler.user.get-list-users"
	getDeleteUsersHandlerName  = "http-handler.user.get-delete-user"
	postDeleteUsersHandlerName = "http-handler.user.post-delete-user"

	getUserProfileHandlerName      = "http-handler.user.get-user-profile"
	getEditUserProfileHandlerName  = "http-handler.user.get-edit-user-profile"
	postEditUserProfileHandlerName = "http-handler.user.post-edit-user-profile"

	userIDParam = "user_id"
)

type Handlers struct {
	gorm   *gorm.DB
	logger *logger.Logger
	conf   *config.Config
	sdk    *sdk.ExternalAPI
	// Services
	sessionSvc *sessionServ.Service
}

func NewUsersHandlers(appCtx *connected_roots.Context) *Handlers {
	loggerEmpty := appCtx.Logger.NewEmpty()
	log := loggerEmpty.WithTag(tracingUserHandlers)

	return &Handlers{
		gorm:       appCtx.Gorm,
		logger:     log,
		conf:       appCtx.Conf,
		sdk:        appCtx.SDK,
		sessionSvc: sessionServ.New(appCtx.Conf, appCtx.Gorm, appCtx.Logger),
	}
}

func (h *Handlers) GetUserViewHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getViewUserHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(getViewUserHandlerName)

	userID := c.Param(userIDParam)
	if userID == "" {
		err := ferrors.ErrPathParamInvalidValue
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	sess, err := h.sessionSvc.Obtain(c.Request().Context(), c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	user, err := h.sdk.ConnectedRootsService.SDK.ObtainUser(ctx, userID)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-users-view.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(c),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonUserViewPageI18N(c)), map[string]interface{}{
			"user": user,
		}))
}

func (h *Handlers) GetUsersListHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getListUsersHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	log := loggerNew.WithTag(getListUsersHandlerName)

	message, err := h.sessionSvc.ObtainMessage(c.Request().Context(), c, "message")
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	log.Debug(fmt.Sprintf("message: %s", message))

	notifications := map[string]interface{}{}
	if len(message) > 0 && message[0] == "success" {
		notifications = map[string]interface{}{
			"notification_type":    "success",
			"notification_title":   translator.T(c, translator.NotificationsAdminUsersDeleteSuccessTitle),
			"notification_message": translator.T(c, translator.NotificationsAdminUsersDeleteSuccessMessage),
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

	users, pagination, err := h.sdk.ConnectedRootsService.SDK.ObtainUsers(ctx, "20", nextCursor, prevCursor, nil, nil, nil)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-users-list.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(c),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonUserListPageI18N(c)), map[string]interface{}{
			"users":           users,
			"protected_roles": h.conf.Roles.Protected,
			"pagination":      pagination,
		}), notifications))
}

func (h *Handlers) GetUserDeleteHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getDeleteUsersHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(getDeleteUsersHandlerName)

	userID := c.Param(userIDParam)
	if userID == "" {
		err := ferrors.ErrPathParamInvalidValue
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	sess, err := h.sessionSvc.Obtain(c.Request().Context(), c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	user, err := h.sdk.ConnectedRootsService.SDK.ObtainUser(ctx, userID)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "admin-users-delete.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(c),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonUserDeletePageI18N(c)), map[string]interface{}{
			"user": user,
		}))
}

func (h *Handlers) PostUserDeleteHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), postDeleteUsersHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(postDeleteUsersHandlerName)

	userID := c.Param(userIDParam)
	if userID == "" {
		err := ferrors.ErrPathParamInvalidValue
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	if err := h.sdk.ConnectedRootsService.SDK.DeleteUser(ctx, userID); err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	if err := h.sessionSvc.SaveMessage(c.Request().Context(), c, "message", "success"); err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Redirect(http.StatusFound, "/admin/users/list")
}

func (h *Handlers) GetUserProfileHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getUserProfileHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(getUserProfileHandlerName)

	sess, err := h.sessionSvc.Obtain(c.Request().Context(), c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	user, err := h.sdk.ConnectedRootsService.SDK.ObtainUser(ctx, sess.Email)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}
	y, m, d := user.CreatedAt.Date()

	return c.Render(http.StatusOK, "profile.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(c),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonUserProfilePageI18N(c)), map[string]interface{}{
			"user_name":       user.Name,
			"user_surname":    user.Surname,
			"user_email":      user.Email,
			"user_telephone":  user.Telephone,
			"user_created_at": fmt.Sprintf("%d/%d/%d", d, int(m), y),
		}))
}

func (h *Handlers) GetEditUserProfileHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getEditUserProfileHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	_ = loggerNew.WithTag(getEditUserProfileHandlerName)

	sess, err := h.sessionSvc.Obtain(c.Request().Context(), c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	user, err := h.sdk.ConnectedRootsService.SDK.ObtainUser(ctx, sess.Email)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "edit-profile.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(c),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonEditUserProfilePageI18N(c)), map[string]interface{}{
			"user_name":      user.Name,
			"user_surname":   user.Surname,
			"user_email":     user.Email,
			"user_telephone": user.Telephone,
		}))
}

func (h *Handlers) PostEditUserProfileHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), postEditUserProfileHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	log := loggerNew.WithTag(postEditUserProfileHandlerName)

	sess, err := h.sessionSvc.Obtain(c.Request().Context(), c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	user, err := h.sdk.ConnectedRootsService.SDK.ObtainUser(ctx, sess.Email)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	if err = c.Request().ParseForm(); err != nil {
		return c.Render(http.StatusOK, "edit-profile.gohtml",
			translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(c),
				bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
				CommonEditUserProfilePageI18N(c)), map[string]interface{}{
				"user_name":            user.Name,
				"user_surname":         user.Surname,
				"user_email":           user.Email,
				"user_telephone":       user.Telephone,
				"notification_type":    "error",
				"notification_title":   translator.T(c, translator.ErrorsNotificationsErrorTitle),
				"notification_message": translator.T(c, translator.ErrorsCommonParseForm),
			}))
	}

	name := c.FormValue("name")
	surname := c.FormValue("surname")
	phone := c.FormValue("phone")

	if name == "" || surname == "" || phone == "" {
		return c.Render(http.StatusOK, "edit-profile.gohtml",
			translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(c),
				bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
				CommonEditUserProfilePageI18N(c)), map[string]interface{}{
				"user_name":            user.Name,
				"user_surname":         user.Surname,
				"user_email":           user.Email,
				"user_telephone":       user.Telephone,
				"notification_type":    "error",
				"notification_title":   translator.T(c, translator.ErrorsNotificationsErrorTitle),
				"notification_message": translator.T(c, translator.ErrorsLoginEmptyEmailOrPassword),
			}))
	}

	log.Debug(fmt.Sprintf("new user name: %s", name))
	log.Debug(fmt.Sprintf("new user surname %s", surname))
	log.Debug(fmt.Sprintf("new user phone: %s", phone))

	user.Name = name
	user.Surname = surname
	user.Telephone = phone

	user, err = h.sdk.ConnectedRootsService.SDK.UpdatePartiallyUser(ctx, user.ToUsersBody())
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	sess.Name = user.Name
	sess.Surname = user.Surname
	if _, err = h.sessionSvc.Save(ctx, c, sess); err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Render(http.StatusOK, "edit-profile.gohtml",
		translator.AddDataKeys(translator.AddDataKeys(translator.AddDataKeys(bars.CommonNavBarI18N(c),
			bars.CommonTopBarI18N(c, sess.Name, sess.Surname)),
			CommonEditUserProfilePageI18N(c)), map[string]interface{}{
			"user_name":            user.Name,
			"user_surname":         user.Surname,
			"user_email":           user.Email,
			"user_telephone":       user.Telephone,
			"notification_type":    "success",
			"notification_title":   translator.T(c, translator.NotificationsEditProfileSuccessTitle),
			"notification_message": translator.T(c, translator.NotificationsEditProfileSuccessMessage),
		}))
}
