package login

import (
	"fmt"
	"net/http"

	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/i18n/translator"

	"github.com/Kortivex/connected_roots/pkg/sdk/sdk_models"

	"github.com/Kortivex/connected_roots/pkg/sdk"

	"github.com/Kortivex/connected_roots/pkg/logger/commons"

	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	sessionServ "github.com/Kortivex/connected_roots/internal/connected_roots/session"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

const (
	tracingLoginHandlers = "http-handler.login"

	getLoginHandlerName  = "http-handler.login.get-login"
	postLoginHandlerName = "http-handler.login.post-login"
)

type Handlers struct {
	gorm   *gorm.DB
	logger *logger.Logger
	conf   *config.Config
	sdk    *sdk.ExternalAPI
	// Services
	sessionSvc *sessionServ.Service
}

func NewLoginHandlers(appCtx *connected_roots.Context) *Handlers {
	loggerEmpty := appCtx.Logger.NewEmpty()
	log := loggerEmpty.WithTag(tracingLoginHandlers)

	return &Handlers{
		gorm:       appCtx.Gorm,
		logger:     log,
		conf:       appCtx.Conf,
		sdk:        appCtx.SDK,
		sessionSvc: sessionServ.New(appCtx.Conf, appCtx.Gorm, appCtx.Logger),
	}
}

func (h *Handlers) GetLoginHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), getLoginHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	log := loggerNew.WithTag(getLoginHandlerName)

	sn, ok, err := h.sessionSvc.IsValid(ctx, c)
	if err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	if ok {
		log.Debug(fmt.Sprintf("user is already logged with email %s", sn.Email))
		return c.Redirect(http.StatusFound, "/")
	}

	return c.Render(http.StatusOK, "sign-in.gohtml", CommonLoginPageI18N(c))
}

func (h *Handlers) PostLoginHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), postLoginHandlerName)
	defer span.End()

	loggerNew := h.logger.New()
	log := loggerNew.WithTag(postLoginHandlerName)

	if err := c.Request().ParseForm(); err != nil {
		return c.Render(http.StatusOK, "sign-in.gohtml", translator.AddDataKeys(CommonLoginPageI18N(c), map[string]interface{}{
			"notification_type":    "error",
			"notification_title":   translator.T(c, translator.ErrorsNotificationsErrorTitle),
			"notification_message": translator.T(c, translator.ErrorsCommonParseForm),
		}))
	}
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" {
		return c.Render(http.StatusOK, "sign-in.gohtml", translator.AddDataKeys(CommonLoginPageI18N(c), map[string]interface{}{
			"notification_type":    "error",
			"notification_title":   translator.T(c, translator.ErrorsNotificationsErrorTitle),
			"notification_message": translator.T(c, translator.ErrorsLoginEmptyEmailOrPassword),
		}))
	}

	log.Debug(fmt.Sprintf("user is logged with email %s", email))

	if isValid, err := h.sdk.ConnectedRootsService.SDK.AuthenticateUser(ctx, email, &sdk_models.UsersAuthenticationBody{
		Email:    email,
		Password: password,
	}); err != nil || !isValid.Valid {
		log.Debug(fmt.Sprintf("credentials are not valid for: %s", email))

		return c.Render(http.StatusOK, "sign-in.gohtml", translator.AddDataKeys(CommonLoginPageI18N(c), map[string]interface{}{
			"notification_type":    "error",
			"notification_title":   translator.T(c, translator.ErrorsNotificationsErrorTitle),
			"notification_message": translator.T(c, translator.ErrorsLoginInvalidCredentials),
		}))
	}

	log.Debug(fmt.Sprintf("credentials are valid for: %s", email))

	user, err := h.sdk.ConnectedRootsService.SDK.ObtainUser(ctx, email)
	if err != nil {
		return c.Render(http.StatusOK, "sign-in.gohtml", translator.AddDataKeys(CommonLoginPageI18N(c), map[string]interface{}{
			"notification_type":    "error",
			"notification_title":   translator.T(c, translator.ErrorsNotificationsErrorTitle),
			"notification_message": translator.T(c, translator.ErrorsLoginInvalidCredentials),
		}))
	}

	sess := &connected_roots.Session{
		Email:    user.Email,
		UserID:   user.ID,
		Name:     user.Name,
		Surname:  user.Surname,
		Language: user.Language,
		Role:     user.Role.Name}

	if _, err = h.sessionSvc.Save(ctx, c, sess); err != nil {
		return commons.NewErrorS(http.StatusInternalServerError, err.Error(), nil, err)
	}

	return c.Redirect(http.StatusFound, "/")
}
