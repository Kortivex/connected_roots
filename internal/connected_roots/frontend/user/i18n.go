package user

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/i18n/translator"
	"github.com/labstack/echo/v4"
)

func CommonUserCreatePageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":                      translator.T(c, translator.PagesCommonSiteTitle),
		"title":                           translator.T(c, translator.PagesAdminUsersCreateTitle),
		"active":                          "users-management",
		"breadcrumb_users":                translator.T(c, translator.BreadcrumbUsers),
		"breadcrumb_create_users":         translator.T(c, translator.BreadcrumbCreateUsers),
		"create_name_label":               translator.T(c, translator.PagesAdminUsersCreateNameLabel),
		"create_surname_label":            translator.T(c, translator.PagesAdminUsersCreateSurnameLabel),
		"create_email_label":              translator.T(c, translator.PagesAdminUsersCreateEmailLabel),
		"create_password_label":           translator.T(c, translator.PagesAdminUsersCreatePasswordLabel),
		"create_phone_label":              translator.T(c, translator.PagesAdminUsersCreatePhoneLabel),
		"create_language_label":           translator.T(c, translator.PagesAdminUsersCreateLanguageLabel),
		"create_language_selection_label": translator.T(c, translator.PagesAdminUsersCreateLanguageSelectionLabel),
		"create_role_label":               translator.T(c, translator.PagesAdminUsersCreateRoleLabel),
		"create_role_selection_label":     translator.T(c, translator.PagesAdminUsersCreateRoleSelectionLabel),
		"create_button_submit_user":       translator.T(c, translator.PagesAdminUsersCreateButtonSubmitUser),
	}
}

func CommonUserViewPageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":            translator.T(c, translator.PagesCommonSiteTitle),
		"title":                 translator.T(c, translator.PagesAdminUsersViewTitle),
		"active":                "users-management",
		"breadcrumb_users":      translator.T(c, translator.BreadcrumbUsers),
		"breadcrumb_view_users": translator.T(c, translator.BreadcrumbViewUsers),
		"view_name_label":       translator.T(c, translator.PagesAdminUsersViewNameLabel),
		"view_surname_label":    translator.T(c, translator.PagesAdminUsersViewSurnameLabel),
		"view_email_label":      translator.T(c, translator.PagesAdminUsersViewEmailLabel),
		"view_phone_label":      translator.T(c, translator.PagesAdminUsersViewPhoneLabel),
		"view_language_label":   translator.T(c, translator.PagesAdminUsersViewLanguageLabel),
		"view_role_label":       translator.T(c, translator.PagesAdminUsersViewRoleLabel),
	}
}

func CommonUserListPageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":                translator.T(c, translator.PagesCommonSiteTitle),
		"title":                     translator.T(c, translator.PagesAdminUsersListTitle),
		"active":                    "users-management",
		"breadcrumb_users":          translator.T(c, translator.BreadcrumbUsers),
		"breadcrumb_list_users":     translator.T(c, translator.BreadcrumbListUsers),
		"list_button_create_user":   translator.T(c, translator.PagesAdminUsersListButtonCreateUser),
		"list_name_label":           translator.T(c, translator.PagesAdminUsersListNameLabel),
		"list_surname_label":        translator.T(c, translator.PagesAdminUsersListSurnameLabel),
		"list_email_label":          translator.T(c, translator.PagesAdminUsersListEmailLabel),
		"list_phone_label":          translator.T(c, translator.PagesAdminUsersListPhoneLabel),
		"list_role_label":           translator.T(c, translator.PagesAdminUsersListRoleLabel),
		"list_actions_label":        translator.T(c, translator.PagesAdminUsersListActionsLabel),
		"list_actions_view_label":   translator.T(c, translator.PagesAdminUsersListActionsViewLabel),
		"list_actions_edit_label":   translator.T(c, translator.PagesAdminUsersListActionsEditLabel),
		"list_actions_remove_label": translator.T(c, translator.PagesAdminUsersListActionsRemoveLabel),
		"list_next_label":           translator.T(c, translator.PagesAdminUsersListNextLabel),
		"list_previous_label":       translator.T(c, translator.PagesAdminUsersListPreviousLabel),
	}
}

func CommonUserDeletePageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":                translator.T(c, translator.PagesCommonSiteTitle),
		"title":                     translator.T(c, translator.PagesAdminUsersDeleteTitle),
		"active":                    "users-management",
		"breadcrumb_users":          translator.T(c, translator.BreadcrumbUsers),
		"breadcrumb_delete_users":   translator.T(c, translator.BreadcrumbDeleteUsers),
		"delete_warning_title":      translator.T(c, translator.PagesAdminUsersDeleteWarningTitle),
		"delete_warning_message":    translator.T(c, translator.PagesAdminUsersDeleteWarningMessage),
		"delete_button_submit_user": translator.T(c, translator.PagesAdminUsersDeleteButtonSubmitUser),
	}
}

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
