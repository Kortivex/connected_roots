package translator

import (
	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func T(c echo.Context, id string) string {
	if lz, ok := c.Get("localizer").(*i18n.Localizer); ok {
		return lz.MustLocalize(&i18n.LocalizeConfig{MessageID: id})
	}
	return ""
}

func AddDataKeys(base, additional map[string]interface{}) map[string]interface{} {
	for k, v := range base {
		additional[k] = v
	}
	return additional
}
