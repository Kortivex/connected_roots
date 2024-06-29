package frontend

import (
	"errors"
	"net/http"

	"github.com/Kortivex/connected_roots/pkg/logger/commons"

	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func (s *Service) SessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(s.conf.Cookie.Name)
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				return commons.NewErrorS(http.StatusUnauthorized, "no session cookie found", nil, err)
			}

			return commons.NewErrorS(http.StatusBadRequest, "error reading cookie", nil, err)
		}

		c.Set(s.conf.Cookie.Name, cookie.Value)
		return next(c)
	}
}

func (s *Service) I18n() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			accept := c.Request().Header.Get("Accept-Language")
			localizer := i18n.NewLocalizer(s.i18n, accept)
			c.Set("localizer", localizer)
			return handlerFunc(c)
		}
	}
}
