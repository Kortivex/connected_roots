package httpserver

import (
	"github.com/labstack/echo/v4"
)

// HTTPErrorHandlerMiddleware is a middleware for an HTTPErrorHandler.
type HTTPErrorHandlerMiddleware func(h echo.HTTPErrorHandler) echo.HTTPErrorHandler

// NOOPErrorHandler is an empty error handler that do not assigns a value to an error.
// This allows to use your own logic to perform this operation in a middleware:
//
//	err = c.JSON(status, err)
//
// Assigning to an error should be done only once in all the error handlers chain so
// use the NOOPErrorHandler handler as a HTTPErrorHandler when one of your
// error middlewares already assigns a value to the err.
func NOOPErrorHandler(err error, c echo.Context) {}
