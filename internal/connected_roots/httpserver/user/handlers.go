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

	tracingPostUsersHandlers     = "http-handler.user.post-user"
	tracingPutUsersHandlers      = "http-handler.user.put-user"
	tracingPatchUsersHandlers    = "http-handler.user.patch-user"
	tracingGetUsersHandlers      = "http-handler.user.get-user"
	tracingListUsersHandlers     = "http-handler.user.list-users"
	tracingDeleteUsersHandlers   = "http-handler.user.delete-user"
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

func (h *UsersHandlers) PostUserHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingPostUsersHandlers)
	defer span.End()

	userBody := connected_roots.Users{}
	if err := c.Bind(&userBody); err != nil {
		err = fmt.Errorf("%s: %w", tracingPostUsersHandlers, errors.ErrInvalidPayload)
		return errors.NewErrorResponse(c, err)
	}

	userRes, err := h.userSvc.Save(ctx, &userBody)
	if err != nil {
		err = fmt.Errorf("%s: %w", tracingPostUsersHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated, userRes)
}

func (h *UsersHandlers) PutUserHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingPutUsersHandlers)
	defer span.End()

	userID := c.Param(userIDParam)
	var userRes *connected_roots.Users
	var err error

	userBody := &connected_roots.Users{}
	if err = c.Bind(userBody); err != nil {
		err = fmt.Errorf("%s: %w", tracingPutUsersHandlers, errors.ErrBodyBadRequestWrongBody)
		return errors.NewErrorResponse(c, err)
	}

	if utils.IsValidEmail(userID) {
		userRes, err = h.userSvc.ObtainFromEmail(ctx, userID)
		if err != nil {
			return errors.NewErrorResponse(c, err)
		}
	} else {
		userRes, err = h.userSvc.ObtainFromID(ctx, userID)
		if err != nil {
			return errors.NewErrorResponse(c, err)
		}
	}

	userRes, err = h.userSvc.Update(ctx, userRes)
	if err != nil {
		err = fmt.Errorf("%s: %w", tracingPutUsersHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, userRes)
}

func (h *UsersHandlers) PatchUserPartiallyHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingPatchUsersHandlers)
	defer span.End()

	userID := c.Param(userIDParam)
	var userRes *connected_roots.Users
	var err error

	userBody := &connected_roots.Users{}
	if err = c.Bind(userBody); err != nil {
		err = fmt.Errorf("%s: %w", tracingPatchUsersHandlers, errors.ErrBodyBadRequestWrongBody)
		return errors.NewErrorResponse(c, err)
	}

	if utils.IsValidEmail(userID) {
		userRes, err = h.userSvc.ObtainFromEmail(ctx, userID)
		if err != nil {
			return errors.NewErrorResponse(c, err)
		}
	} else {
		userRes, err = h.userSvc.ObtainFromID(ctx, userID)
		if err != nil {
			return errors.NewErrorResponse(c, err)
		}
	}

	userRes.Name = userBody.Name
	userRes.Surname = userBody.Surname
	userRes.Telephone = userBody.Telephone

	userRes, err = h.userSvc.Update(ctx, userRes)
	if err != nil {
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated, userRes)
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

func (h *UsersHandlers) ListUsersHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingListUsersHandlers)
	defer span.End()

	filters := connected_roots.UserPaginationFilters{}
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &filters); err != nil {
		err = fmt.Errorf("%s: %w", tracingListUsersHandlers, errors.ErrQueryParamInvalidValue)
		return errors.NewErrorResponse(c, err)
	}

	rolesRes, err := h.userSvc.ObtainAll(ctx, &filters)
	if err != nil {
		err = fmt.Errorf("%s: %w", tracingListUsersHandlers, err)
		return errors.NewErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, rolesRes)
}

func (h *UsersHandlers) DeleteUserHandler(c echo.Context) error {
	ctx, span := otel.Tracer(h.conf.App.Name).Start(c.Request().Context(), tracingDeleteUsersHandlers)
	defer span.End()

	userID := c.Param(userIDParam)
	var userRes *connected_roots.Users
	var err error

	if utils.IsValidEmail(userID) {
		userRes, err = h.userSvc.ObtainFromEmail(ctx, userID)
		if err != nil {
			return errors.NewErrorResponse(c, err)
		}
		if err = h.userSvc.RemoveByEmail(ctx, userRes.Email); err != nil {
			return errors.NewErrorResponse(c, err)
		}

		return c.NoContent(http.StatusNoContent)
	}

	userRes, err = h.userSvc.ObtainFromID(ctx, userID)
	if err != nil {
		return errors.NewErrorResponse(c, err)
	}
	if err = h.userSvc.RemoveByID(ctx, userRes.ID); err != nil {
		return errors.NewErrorResponse(c, err)
	}

	return c.NoContent(http.StatusNoContent)
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
