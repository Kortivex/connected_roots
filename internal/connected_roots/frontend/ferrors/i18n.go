package ferrors

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/i18n/translator"
	"github.com/labstack/echo/v4"
)

func CommonErrorsPageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"title": translator.T(c, translator.PagesErrorsTitle),
	}
}
