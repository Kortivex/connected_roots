package httpserver

import (
	"context"

	"github.com/labstack/echo/v4"
)

// MiddlewareRequestIDToContext set the request ID to a context key interface.
func MiddlewareRequestIDToContext(contextKeys ...interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			for _, ck := range contextKeys {
				ctx := context.WithValue(c.Request().Context(), ck, GetRequestID(c))

				c.SetRequest(c.Request().WithContext(ctx))
			}

			return next(c)
		}
	}
}
