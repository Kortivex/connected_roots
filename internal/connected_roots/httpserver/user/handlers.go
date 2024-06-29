package user

import (
	"fmt"
	"net/http"

	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/internal/connected_roots/httpserver/errors"
	"github.com/Kortivex/connected_roots/internal/connected_roots/user"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/Kortivex/connected_roots/pkg/utils"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

const (
	tracingUsersHandlers = "http-handler.user"

	tracingGetUsersHandlers      = "http-handler.user.get-user"
	tracingPostUsersAuthHandlers = "http-handler.user.post-user-auth"

	userIDParam = "user_id"
)

type UsersHandlers struct {
	gorm   *gorm.DB
	logger *logger.Logger
	conf   *config.Config
	// Services.
	userSvc *user.Service
}

func NewUsersHandlers(appCtx *connected_roots.Context) *UsersHandlers {
	loggerEmpty := appCtx.Logger.NewEmpty()
	log := loggerEmpty.WithTag(tracingUsersHandlers)

	return &UsersHandlers{
		gorm:    appCtx.Gorm,
		logger:  log,
		conf:    appCtx.Conf,
		userSvc: user.New(appCtx.Conf, appCtx.Gorm, appCtx.Logger),
	}
}

func (h *UsersHandlers) GetUserHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingGetUsersHandlers)
	defer span.End()

	userID := c.Param(userIDParam)
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

	return c.JSON(http.StatusOK, userRes)
}

func (h *UsersHandlers) PostUserAuthHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingPostUsersAuthHandlers)
	defer span.End()

	userID := c.Param(userIDParam)
	var userRes *connected_roots.Users
	var err error

	userBody := &connected_roots.UsersAuthentication{}
	if err = c.Bind(userBody); err != nil {
		err = fmt.Errorf("%s: %w", tracingPostUsersAuthHandlers, errors.ErrBodyBadRequestWrongBody)
		return errors.NewErrorResponse(c, err)
	}

	if !utils.IsValidEmail(userID) {
		userRes, err = h.userSvc.ObtainFromID(ctx, userID)
		if err != nil {
			return errors.NewErrorResponse(c, err)
		}
		userBody.Email = userRes.Email
	}

	ok, err := h.userSvc.IsValidPassword(ctx, userBody.Email, userBody.Password)
	if err != nil {
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, &connected_roots.UsersAuthentication{Valid: ok})
}
