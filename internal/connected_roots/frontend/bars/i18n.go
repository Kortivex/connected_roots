package bars

import (
	"context"
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/i18n/translator"
	"github.com/Kortivex/connected_roots/internal/connected_roots/session"
	"github.com/labstack/echo/v4"
)

func CommonNavBarI18N(ctx context.Context, c echo.Context, sessionSvc *session.Service) map[string]interface{} {
	isAdmin, _ := sessionSvc.IsAdmin(ctx, c)
	isTech, _ := sessionSvc.IsTechnical(ctx, c)
	isUser, _ := sessionSvc.IsUser(ctx, c)

	return map[string]interface{}{
		"admin_role":                      isAdmin,
		"tech_role":                       isTech,
		"user_role":                       isUser,
		"admin_nav_label":                 translator.T(c, translator.NavbarLabelsAdminNavLabel),
		"role_nav_label":                  translator.T(c, translator.NavbarLabelsRoleNavLabel),
		"role_nav_management_label":       translator.T(c, translator.NavbarLabelsRoleNavManagementLabel),
		"user_nav_label":                  translator.T(c, translator.NavbarLabelsUserNavLabel),
		"user_nav_management_label":       translator.T(c, translator.NavbarLabelsUserNavManagementLabel),
		"orchard_nav_label":               translator.T(c, translator.NavbarLabelsOrchardNavLabel),
		"orchard_nav_management_label":    translator.T(c, translator.NavbarLabelsOrchardNavManagementLabel),
		"crop_type_nav_label":             translator.T(c, translator.NavbarLabelsCropTypeNavLabel),
		"crop_type_nav_management_label":  translator.T(c, translator.NavbarLabelsCropTypeNavManagementLabel),
		"sensor_nav_label":                translator.T(c, translator.NavbarLabelsSensorNavLabel),
		"sensor_nav_management_label":     translator.T(c, translator.NavbarLabelsSensorNavManagementLabel),
		"users_nav_label":                 translator.T(c, translator.NavbarLabelsUsersNavLabel),
		"users_sensor_nav_label":          translator.T(c, translator.NavbarLabelsUserSensorNavLabel),
		"users_sensor_nav_list_label":     translator.T(c, translator.NavbarLabelsUserSensorNavListLabel),
		"users_orchard_nav_label":         translator.T(c, translator.NavbarLabelsUserOrchardNavLabel),
		"users_orchard_nav_list_label":    translator.T(c, translator.NavbarLabelsUserOrchardNavListLabel),
		"users_activities_nav_label":      translator.T(c, translator.NavbarLabelsUserActivitiesNavLabel),
		"users_activities_nav_list_label": translator.T(c, translator.NavbarLabelsUserActivitiesNavListLabel),
		"collapsed_view_label":            translator.T(c, translator.NavbarLabelsCollapsedViewLabel),
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
