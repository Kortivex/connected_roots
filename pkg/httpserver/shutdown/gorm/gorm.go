// Package gorm is used to shut down Gorm.
package gorm

import (
	"github.com/Kortivex/connected_roots/pkg/httpserver"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Shutdown is a function to shut down gorm when on a server shutdown occurs.
func Shutdown(db *gorm.DB) httpserver.ServerShutdownFunc {
	return func(e *echo.Echo) {
		sqlDB, err := db.DB()
		if err != nil {
			return
		}

		_ = sqlDB.Close()
	}
}
