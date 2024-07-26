package activity

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/i18n/translator"
	"github.com/labstack/echo/v4"
)

func CommonActivityCreatePageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":                     translator.T(c, translator.PagesCommonSiteTitle),
		"title":                          translator.T(c, translator.PagesUserActivitiesCreateTitle),
		"active":                         "user-activities-list",
		"breadcrumb_activities":          translator.T(c, translator.BreadcrumbActivities),
		"breadcrumb_create_activities":   translator.T(c, translator.BreadcrumbCreateActivities),
		"create_name_label":              translator.T(c, translator.PagesUserActivitiesCreateNameLabel),
		"create_description_label":       translator.T(c, translator.PagesUserActivitiesCreateDescriptionLabel),
		"create_orchard_label":           translator.T(c, translator.PagesUserActivitiesCreateOrchardLabel),
		"create_orchard_selection_label": translator.T(c, translator.PagesUserActivitiesCreateOrchardSelectionLabel),
		"create_button_submit_activity":  translator.T(c, translator.PagesUserActivitiesCreateButtonSubmitActivity),
	}
}

func CommonActivityUpdatePageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":                     translator.T(c, translator.PagesCommonSiteTitle),
		"title":                          translator.T(c, translator.PagesUserActivitiesUpdateTitle),
		"active":                         "user-activities-list",
		"breadcrumb_activities":          translator.T(c, translator.BreadcrumbActivities),
		"breadcrumb_update_activities":   translator.T(c, translator.BreadcrumbUpdateActivities),
		"update_name_label":              translator.T(c, translator.PagesUserActivitiesUpdateNameLabel),
		"update_description_label":       translator.T(c, translator.PagesUserActivitiesUpdateDescriptionLabel),
		"update_orchard_label":           translator.T(c, translator.PagesUserActivitiesUpdateOrchardLabel),
		"update_orchard_selection_label": translator.T(c, translator.PagesUserActivitiesUpdateOrchardSelectionLabel),
		"update_button_submit_activity":  translator.T(c, translator.PagesUserActivitiesUpdateButtonSubmitActivity),
	}
}

func CommonActivityViewPageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":                 translator.T(c, translator.PagesCommonSiteTitle),
		"title":                      translator.T(c, translator.PagesUserActivitiesViewTitle),
		"active":                     "user-activities-list",
		"breadcrumb_activities":      translator.T(c, translator.BreadcrumbActivities),
		"breadcrumb_view_activities": translator.T(c, translator.BreadcrumbViewActivities),
		"view_name_label":            translator.T(c, translator.PagesUserActivitiesViewNameLabel),
		"view_description_label":     translator.T(c, translator.PagesUserActivitiesViewDescriptionNameLabel),
		"view_orchard_label":         translator.T(c, translator.PagesUserActivitiesViewOrchardLabel),
	}
}

func CommonActivityListPageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":                  translator.T(c, translator.PagesCommonSiteTitle),
		"title":                       translator.T(c, translator.PagesUserActivitiesListTitle),
		"active":                      "user-activities-list",
		"breadcrumb_activities":       translator.T(c, translator.BreadcrumbActivities),
		"breadcrumb_list_activities":  translator.T(c, translator.BreadcrumbListActivities),
		"list_button_create_activity": translator.T(c, translator.PagesUserActivitiesListButtonCreateActivity),
		"list_name_label":             translator.T(c, translator.PagesUserActivitiesListNameLabel),
		"list_description_label":      translator.T(c, translator.PagesUserActivitiesListDescriptionLabel),
		"list_orchard_label":          translator.T(c, translator.PagesUserActivitiesListOrchardNameLabel),
		"list_actions_label":          translator.T(c, translator.PagesUserActivitiesListActionsLabel),
		"list_actions_view_label":     translator.T(c, translator.PagesUserActivitiesListActionsViewLabel),
		"list_actions_edit_label":     translator.T(c, translator.PagesUserActivitiesListActionsEditLabel),
		"list_actions_remove_label":   translator.T(c, translator.PagesUserActivitiesListActionsRemoveLabel),
		"list_next_label":             translator.T(c, translator.PagesUserActivitiesListNextLabel),
		"list_previous_label":         translator.T(c, translator.PagesUserActivitiesListPreviousLabel),
	}
}

func CommonActivityDeletePageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":                    translator.T(c, translator.PagesCommonSiteTitle),
		"title":                         translator.T(c, translator.PagesUserActivitiesDeleteTitle),
		"active":                        "user-activities-list",
		"breadcrumb_activities":         translator.T(c, translator.BreadcrumbActivities),
		"breadcrumb_delete_activities":  translator.T(c, translator.BreadcrumbDeleteActivities),
		"delete_warning_title":          translator.T(c, translator.PagesUserActivitiesDeleteWarningTitle),
		"delete_warning_message":        translator.T(c, translator.PagesUserActivitiesDeleteWarningMessage),
		"delete_button_submit_activity": translator.T(c, translator.PagesUserActivitiesDeleteButtonSubmitActivity),
	}
}
