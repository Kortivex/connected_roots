package navbar

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/i18n/translator"
	"github.com/labstack/echo/v4"
)

func CommonNavBarI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"admin_nav_label":           translator.T(c, translator.NavbarLabelsAdminNavLabel),
		"role_nav_label":            translator.T(c, translator.NavbarLabelsRoleNavLabel),
		"role_nav_management_label": translator.T(c, translator.NavbarLabelsRoleNavManagementLabel),
		"user_nav_label":            translator.T(c, translator.NavbarLabelsUserNavLabel),
		"user_nav_management_label": translator.T(c, translator.NavbarLabelsUserNavManagementLabel),
		"collapsed_view_label":      translator.T(c, translator.NavbarLabelsCollapsedViewLabel),
	}
}
