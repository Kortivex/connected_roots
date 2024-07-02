package user

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/i18n/translator"
	"github.com/labstack/echo/v4"
)

func CommonUserProfilePageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":          translator.T(c, translator.PagesCommonSiteTitle),
		"title":               translator.T(c, translator.PagesProfileTitle),
		"breadcrumb_user":     translator.T(c, translator.BreadcrumbUser),
		"breadcrumb_profile":  translator.T(c, translator.BreadcrumbProfile),
		"button_edit_profile": translator.T(c, translator.PagesProfileButtonEditProfile),
		"joined_at_label":     translator.T(c, translator.PagesProfileJoinedAtLabel),
		"data_label":          translator.T(c, translator.PagesProfileDataLabel),
		"email_label":         translator.T(c, translator.PagesProfileEmailLabel),
		"phone_label":         translator.T(c, translator.PagesProfilePhoneLabel),
	}
}

func CommonEditUserProfilePageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":              translator.T(c, translator.PagesCommonSiteTitle),
		"title":                   translator.T(c, translator.PagesProfileEditTitle),
		"breadcrumb_user":         translator.T(c, translator.BreadcrumbUser),
		"breadcrumb_profile":      translator.T(c, translator.BreadcrumbProfile),
		"breadcrumb_edit_profile": translator.T(c, translator.BreadcrumbEditProfile),
		"name_label":              translator.T(c, translator.PagesProfileEditNameLabel),
		"name_placeholder":        translator.T(c, translator.PagesProfileEditNamePlaceholder),
		"surname_label":           translator.T(c, translator.PagesProfileEditSurnameLabel),
		"surname_placeholder":     translator.T(c, translator.PagesProfileEditSurnamePlaceholder),
		"email_label":             translator.T(c, translator.PagesProfileEditEmailLabel),
		"phone_label":             translator.T(c, translator.PagesProfileEditPhoneLabel),
		"button_submit":           translator.T(c, translator.PagesProfileEditButtonSubmit),
	}
}
