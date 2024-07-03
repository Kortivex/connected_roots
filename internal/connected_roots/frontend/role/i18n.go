package role

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/i18n/translator"
	"github.com/labstack/echo/v4"
)

func CommonRoleListPageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":                translator.T(c, translator.PagesCommonSiteTitle),
		"title":                     translator.T(c, translator.PagesAdminRolesListTitle),
		"breadcrumb_roles":          translator.T(c, translator.BreadcrumbRoles),
		"breadcrumb_list_roles":     translator.T(c, translator.BreadcrumbListRoles),
		"list_button_create_role":   translator.T(c, translator.PagesAdminRolesListButtonCreateRole),
		"list_name_label":           translator.T(c, translator.PagesAdminRolesListNameLabel),
		"list_description_label":    translator.T(c, translator.PagesAdminRolesListDescriptionLabel),
		"list_actions_label":        translator.T(c, translator.PagesAdminRolesListActionsLabel),
		"list_actions_view_label":   translator.T(c, translator.PagesAdminRolesListActionsViewLabel),
		"list_actions_edit_label":   translator.T(c, translator.PagesAdminRolesListActionsEditLabel),
		"list_actions_remove_label": translator.T(c, translator.PagesAdminRolesListActionsRemoveLabel),
		"list_next_label":           translator.T(c, translator.PagesAdminRolesListNextLabel),
		"list_previous_label":       translator.T(c, translator.PagesAdminRolesListPreviousLabel),
	}
}