package logout

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/i18n/translator"
	"github.com/labstack/echo/v4"
)

func CommonLogoutPageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":         translator.T(c, translator.PagesCommonSiteTitle),
		"title":              translator.T(c, translator.PagesLogoutTitle),
		"logout_label":       translator.T(c, translator.PagesLogoutLogoutLabel),
		"thanks_label":       translator.T(c, translator.PagesLogoutThanksLabel),
		"logout_information": translator.T(c, translator.PagesLogoutLogoutInformation),
		"logout_button":      translator.T(c, translator.PagesLogoutLogoutButton),
	}
}
