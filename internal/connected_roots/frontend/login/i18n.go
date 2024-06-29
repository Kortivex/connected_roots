package login

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/i18n/translator"
	"github.com/labstack/echo/v4"
)

func CommonLoginPageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":           translator.T(c, translator.PagesCommonSiteTitle),
		"title":                translator.T(c, translator.PagesLoginTitle),
		"sign_in_label":        translator.T(c, translator.PagesLoginSignInLabel),
		"email_label":          translator.T(c, translator.PagesLoginEmailLabel),
		"email_placeholder":    translator.T(c, translator.PagesLoginEmailPlaceholder),
		"password_label":       translator.T(c, translator.PagesLoginPasswordLabel),
		"password_placeholder": translator.T(c, translator.PagesLoginPasswordPlaceholder),
		"sign_in_button":       translator.T(c, translator.PagesLoginSignInButton),
	}
}
