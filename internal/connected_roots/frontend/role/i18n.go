package role

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/i18n/translator"
	"github.com/labstack/echo/v4"
)

func CommonRoleCreatePageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":                translator.T(c, translator.PagesCommonSiteTitle),
		"title":                     translator.T(c, translator.PagesAdminRolesCreateTitle),
		"active":                    "roles-management",
		"breadcrumb_roles":          translator.T(c, translator.BreadcrumbRoles),
		"breadcrumb_create_roles":   translator.T(c, translator.BreadcrumbCreateRoles),
		"create_name_label":         translator.T(c, translator.PagesAdminRolesCreateNameLabel),
		"create_description_label":  translator.T(c, translator.PagesAdminRolesCreateDescriptionLabel),
		"create_protected_label":    translator.T(c, translator.PagesAdminRolesCreateProtectedLabel),
		"create_button_submit_role": translator.T(c, translator.PagesAdminRolesCreateButtonSubmitRole),
	}
}

func CommonRoleUpdatePageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":                translator.T(c, translator.PagesCommonSiteTitle),
		"title":                     translator.T(c, translator.PagesAdminRolesUpdateTitle),
		"active":                    "roles-management",
		"breadcrumb_roles":          translator.T(c, translator.BreadcrumbRoles),
		"breadcrumb_update_roles":   translator.T(c, translator.BreadcrumbUpdateRoles),
		"update_name_label":         translator.T(c, translator.PagesAdminRolesUpdateNameLabel),
		"update_description_label":  translator.T(c, translator.PagesAdminRolesUpdateDescriptionLabel),
		"update_protected_label":    translator.T(c, translator.PagesAdminRolesUpdateProtectedLabel),
		"update_button_submit_role": translator.T(c, translator.PagesAdminRolesUpdateButtonSubmitRole),
	}
}

func CommonRoleViewPageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":             translator.T(c, translator.PagesCommonSiteTitle),
		"title":                  translator.T(c, translator.PagesAdminRolesViewTitle),
		"active":                 "roles-management",
		"breadcrumb_roles":       translator.T(c, translator.BreadcrumbRoles),
		"breadcrumb_view_roles":  translator.T(c, translator.BreadcrumbViewRoles),
		"view_name_label":        translator.T(c, translator.PagesAdminRolesViewNameLabel),
		"view_description_label": translator.T(c, translator.PagesAdminRolesViewDescriptionLabel),
		"view_protected_label":   translator.T(c, translator.PagesAdminRolesViewProtectedLabel),
	}
}

func CommonRoleListPageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":                translator.T(c, translator.PagesCommonSiteTitle),
		"title":                     translator.T(c, translator.PagesAdminRolesListTitle),
		"active":                    "roles-management",
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

func CommonRoleDeletePageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":                translator.T(c, translator.PagesCommonSiteTitle),
		"title":                     translator.T(c, translator.PagesAdminRolesDeleteTitle),
		"active":                    "roles-management",
		"breadcrumb_roles":          translator.T(c, translator.BreadcrumbRoles),
		"breadcrumb_delete_roles":   translator.T(c, translator.BreadcrumbDeleteRoles),
		"delete_warning_title":      translator.T(c, translator.PagesAdminRolesDeleteWarningTitle),
		"delete_warning_message":    translator.T(c, translator.PagesAdminRolesDeleteWarningMessage),
		"delete_button_submit_role": translator.T(c, translator.PagesAdminRolesDeleteButtonSubmitRole),
	}
}
