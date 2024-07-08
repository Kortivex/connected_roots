package bars

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/i18n/translator"
	"github.com/labstack/echo/v4"
)

func CommonNavBarI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"admin_nav_label":              translator.T(c, translator.NavbarLabelsAdminNavLabel),
		"role_nav_label":               translator.T(c, translator.NavbarLabelsRoleNavLabel),
		"role_nav_management_label":    translator.T(c, translator.NavbarLabelsRoleNavManagementLabel),
		"user_nav_label":               translator.T(c, translator.NavbarLabelsUserNavLabel),
		"user_nav_management_label":    translator.T(c, translator.NavbarLabelsUserNavManagementLabel),
		"orchard_nav_label":            translator.T(c, translator.NavbarLabelsOrchardNavLabel),
		"orchard_nav_management_label": translator.T(c, translator.NavbarLabelsOrchardNavManagementLabel),
		"collapsed_view_label":         translator.T(c, translator.NavbarLabelsCollapsedViewLabel),
	}
}

func CommonTopBarI18N(c echo.Context, name, surname string) map[string]interface{} {
	return map[string]interface{}{
		"user_session_name":    name,
		"user_session_surname": surname,
		"profile_label":        translator.T(c, translator.TopBarLabelsProfileLabel),
		"sign_out_label":       translator.T(c, translator.TopBarLabelsSignOutLabel),
	}
}
