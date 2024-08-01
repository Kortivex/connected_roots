package orchard

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/i18n/translator"
	"github.com/labstack/echo/v4"
)

func CommonOrchardCreatePageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":                       translator.T(c, translator.PagesCommonSiteTitle),
		"title":                            translator.T(c, translator.PagesAdminOrchardsCreateTitle),
		"active":                           "orchards-management",
		"breadcrumb_orchards":              translator.T(c, translator.BreadcrumbOrchards),
		"breadcrumb_create_orchards":       translator.T(c, translator.BreadcrumbCreateOrchards),
		"create_name_label":                translator.T(c, translator.PagesAdminOrchardsCreateNameLabel),
		"create_location_label":            translator.T(c, translator.PagesAdminOrchardsCreateLocationLabel),
		"create_size_label":                translator.T(c, translator.PagesAdminOrchardsCreateSizeLabel),
		"create_soil_label":                translator.T(c, translator.PagesAdminOrchardsCreateSoilLabel),
		"create_fertilizer_label":          translator.T(c, translator.PagesAdminOrchardsCreateFertilizerLabel),
		"create_composting_label":          translator.T(c, translator.PagesAdminOrchardsCreateCompostingLabel),
		"create_user_label":                translator.T(c, translator.PagesAdminOrchardsCreateUserLabel),
		"create_user_selection_label":      translator.T(c, translator.PagesAdminOrchardsCreateUserSelectionLabel),
		"create_crop_type_label":           translator.T(c, translator.PagesAdminOrchardsCreateCropTypeLabel),
		"create_crop_type_selection_label": translator.T(c, translator.PagesAdminOrchardsCreateCropTypeSelectionLabel),
		"create_button_submit_orchard":     translator.T(c, translator.PagesAdminOrchardsCreateButtonSubmitOrchard),
	}
}

func CommonOrchardUpdatePageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":                       translator.T(c, translator.PagesCommonSiteTitle),
		"title":                            translator.T(c, translator.PagesAdminOrchardsUpdateTitle),
		"active":                           "orchards-management",
		"breadcrumb_orchards":              translator.T(c, translator.BreadcrumbOrchards),
		"breadcrumb_update_orchards":       translator.T(c, translator.BreadcrumbUpdateOrchards),
		"update_name_label":                translator.T(c, translator.PagesAdminOrchardsUpdateNameLabel),
		"update_location_label":            translator.T(c, translator.PagesAdminOrchardsUpdateLocationLabel),
		"update_size_label":                translator.T(c, translator.PagesAdminOrchardsUpdateSizeLabel),
		"update_soil_label":                translator.T(c, translator.PagesAdminOrchardsUpdateSoilLabel),
		"update_fertilizer_label":          translator.T(c, translator.PagesAdminOrchardsUpdateFertilizerLabel),
		"update_composting_label":          translator.T(c, translator.PagesAdminOrchardsUpdateCompostingLabel),
		"update_user_label":                translator.T(c, translator.PagesAdminOrchardsUpdateUserLabel),
		"update_user_selection_label":      translator.T(c, translator.PagesAdminOrchardsUpdateUserSelectionLabel),
		"update_crop_type_label":           translator.T(c, translator.PagesAdminOrchardsUpdateCropTypeLabel),
		"update_crop_type_selection_label": translator.T(c, translator.PagesAdminOrchardsUpdateCropTypeSelectionLabel),
		"update_button_submit_orchard":     translator.T(c, translator.PagesAdminOrchardsUpdateButtonSubmitOrchard),
	}
}

func CommonOrchardViewPageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":               translator.T(c, translator.PagesCommonSiteTitle),
		"title":                    translator.T(c, translator.PagesAdminOrchardsViewTitle),
		"active":                   "orchards-management",
		"breadcrumb_orchards":      translator.T(c, translator.BreadcrumbOrchards),
		"breadcrumb_view_orchards": translator.T(c, translator.BreadcrumbViewOrchards),
		"view_name_label":          translator.T(c, translator.PagesAdminOrchardsViewNameLabel),
		"view_location_label":      translator.T(c, translator.PagesAdminOrchardsViewLocationLabel),
		"view_size_label":          translator.T(c, translator.PagesAdminOrchardsViewSizeLabel),
		"view_soil_label":          translator.T(c, translator.PagesAdminOrchardsViewSoilLabel),
		"view_fertilizer_label":    translator.T(c, translator.PagesAdminOrchardsViewFertilizerLabel),
		"view_composting_label":    translator.T(c, translator.PagesAdminOrchardsViewCompostingLabel),
		"view_user_label":          translator.T(c, translator.PagesAdminOrchardsViewUserLabel),
		"view_crop_type_label":     translator.T(c, translator.PagesAdminOrchardsViewCropTypeLabel),
	}
}

func CommonOrchardViewReportPageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title": translator.T(c, translator.PagesCommonSiteTitle),
		"title":      translator.T(c, translator.PagesOrchardsViewReportTitle),
		"active":     "orchards-management",
		"days_of_week": []string{
			translator.T(c, translator.PagesOrchardsMonday),
			translator.T(c, translator.PagesOrchardsTuesday),
			translator.T(c, translator.PagesOrchardsWednesday),
			translator.T(c, translator.PagesOrchardsThursday),
			translator.T(c, translator.PagesOrchardsFriday),
			translator.T(c, translator.PagesOrchardsSaturday),
			translator.T(c, translator.PagesOrchardsSunday),
		},
		"chart_titles": []string{
			translator.T(c, translator.PagesOrchardTemperature),
			translator.T(c, translator.PagesOrchardHumidity),
			translator.T(c, translator.PagesOrchardSoil),
			translator.T(c, translator.PagesOrchardSalt),
			translator.T(c, translator.PagesOrchardLight),
			translator.T(c, translator.PagesOrchardAltitude),
			translator.T(c, translator.PagesOrchardPressure),
		},
		"chart_legend": []string{
			translator.T(c, translator.PagesOrchardInternal),
			translator.T(c, translator.PagesOrchardExternal),
		},
	}
}

func CommonOrchardListPageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":                     translator.T(c, translator.PagesCommonSiteTitle),
		"title":                          translator.T(c, translator.PagesAdminOrchardsListTitle),
		"active":                         "orchards-management",
		"breadcrumb_orchards":            translator.T(c, translator.BreadcrumbOrchards),
		"breadcrumb_list_orchards":       translator.T(c, translator.BreadcrumbListOrchards),
		"list_button_create_orchard":     translator.T(c, translator.PagesAdminOrchardsListButtonCreateOrchard),
		"list_name_label":                translator.T(c, translator.PagesAdminOrchardsListNameLabel),
		"list_location_label":            translator.T(c, translator.PagesAdminOrchardsListLocationLabel),
		"list_user_label":                translator.T(c, translator.PagesAdminOrchardsListUserLabel),
		"list_crop_type_label":           translator.T(c, translator.PagesAdminOrchardsListCropTypeLabel),
		"list_actions_label":             translator.T(c, translator.PagesAdminOrchardsListActionsLabel),
		"list_actions_view_label":        translator.T(c, translator.PagesAdminOrchardsListActionsViewLabel),
		"list_actions_view_report_label": translator.T(c, translator.PagesAdminOrchardsListActionsViewReportLabel),
		"list_actions_edit_label":        translator.T(c, translator.PagesAdminOrchardsListActionsEditLabel),
		"list_actions_remove_label":      translator.T(c, translator.PagesAdminOrchardsListActionsRemoveLabel),
		"list_next_label":                translator.T(c, translator.PagesAdminOrchardsListNextLabel),
		"list_previous_label":            translator.T(c, translator.PagesAdminOrchardsListPreviousLabel),
	}
}

func CommonOrchardDeletePageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":                   translator.T(c, translator.PagesCommonSiteTitle),
		"title":                        translator.T(c, translator.PagesAdminOrchardsDeleteTitle),
		"active":                       "orchards-management",
		"breadcrumb_roles":             translator.T(c, translator.BreadcrumbOrchards),
		"breadcrumb_delete_orchards":   translator.T(c, translator.BreadcrumbDeleteOrchards),
		"delete_warning_title":         translator.T(c, translator.PagesAdminOrchardsDeleteWarningTitle),
		"delete_warning_message":       translator.T(c, translator.PagesAdminOrchardsDeleteWarningMessage),
		"delete_button_submit_orchard": translator.T(c, translator.PagesAdminOrchardsDeleteButtonSubmitOrchard),
	}
}

func CommonOrchardUserViewPageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":               translator.T(c, translator.PagesCommonSiteTitle),
		"title":                    translator.T(c, translator.PagesAdminOrchardsViewTitle),
		"active":                   "user-orchards-list",
		"breadcrumb_orchards":      translator.T(c, translator.BreadcrumbOrchards),
		"breadcrumb_view_orchards": translator.T(c, translator.BreadcrumbViewOrchards),
		"view_name_label":          translator.T(c, translator.PagesAdminOrchardsViewNameLabel),
		"view_location_label":      translator.T(c, translator.PagesAdminOrchardsViewLocationLabel),
		"view_size_label":          translator.T(c, translator.PagesAdminOrchardsViewSizeLabel),
		"view_soil_label":          translator.T(c, translator.PagesAdminOrchardsViewSoilLabel),
		"view_fertilizer_label":    translator.T(c, translator.PagesAdminOrchardsViewFertilizerLabel),
		"view_composting_label":    translator.T(c, translator.PagesAdminOrchardsViewCompostingLabel),
		"view_user_label":          translator.T(c, translator.PagesAdminOrchardsViewUserLabel),
		"view_crop_type_label":     translator.T(c, translator.PagesAdminOrchardsViewCropTypeLabel),
	}
}

func CommonOrchardUserListPageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":                 translator.T(c, translator.PagesCommonSiteTitle),
		"title":                      translator.T(c, translator.PagesAdminOrchardsListTitle),
		"active":                     "user-orchards-list",
		"breadcrumb_orchards":        translator.T(c, translator.BreadcrumbOrchards),
		"breadcrumb_list_orchards":   translator.T(c, translator.BreadcrumbListOrchards),
		"list_button_create_orchard": translator.T(c, translator.PagesAdminOrchardsListButtonCreateOrchard),
		"list_name_label":            translator.T(c, translator.PagesAdminOrchardsListNameLabel),
		"list_location_label":        translator.T(c, translator.PagesAdminOrchardsListLocationLabel),
		"list_user_label":            translator.T(c, translator.PagesAdminOrchardsListUserLabel),
		"list_crop_type_label":       translator.T(c, translator.PagesAdminOrchardsListCropTypeLabel),
		"list_actions_label":         translator.T(c, translator.PagesAdminOrchardsListActionsLabel),
		"list_actions_view_label":    translator.T(c, translator.PagesAdminOrchardsListActionsViewLabel),
		"list_actions_edit_label":    translator.T(c, translator.PagesAdminOrchardsListActionsEditLabel),
		"list_actions_remove_label":  translator.T(c, translator.PagesAdminOrchardsListActionsRemoveLabel),
		"list_next_label":            translator.T(c, translator.PagesAdminOrchardsListNextLabel),
		"list_previous_label":        translator.T(c, translator.PagesAdminOrchardsListPreviousLabel),
	}
}
