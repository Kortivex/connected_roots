package crop_types

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/i18n/translator"
	"github.com/labstack/echo/v4"
)

func CommonCropTypeViewPageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":                 translator.T(c, translator.PagesCommonSiteTitle),
		"title":                      translator.T(c, translator.PagesAdminCropTypesViewTitle),
		"active":                     "crop-types-management",
		"breadcrumb_crop_types":      translator.T(c, translator.BreadcrumbCropTypes),
		"breadcrumb_view_crop_types": translator.T(c, translator.BreadcrumbViewCropTypes),
		"view_name_label":            translator.T(c, translator.PagesAdminCropTypesViewNameLabel),
		"view_scientific_name_label": translator.T(c, translator.PagesAdminCropTypesViewScientificNameLabel),
		"view_life_cycle_label":      translator.T(c, translator.PagesAdminCropTypesViewLifeCycleLabel),
		"view_planting_season_label": translator.T(c, translator.PagesAdminCropTypesViewPlantingSeasonLabel),
		"view_harvest_season_label":  translator.T(c, translator.PagesAdminCropTypesViewHarvestSeasonLabel),
		"view_irrigation_label":      translator.T(c, translator.PagesAdminCropTypesViewIrrigationLabel),
	}
}

func CommonCropTypeListPageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":                   translator.T(c, translator.PagesCommonSiteTitle),
		"title":                        translator.T(c, translator.PagesAdminCropTypesListTitle),
		"active":                       "crop-types-management",
		"breadcrumb_crop_types":        translator.T(c, translator.BreadcrumbCropTypes),
		"breadcrumb_list_crop_types":   translator.T(c, translator.BreadcrumbListCropTypes),
		"list_button_create_crop_type": translator.T(c, translator.PagesAdminCropTypesListButtonCreateCropType),
		"list_name_label":              translator.T(c, translator.PagesAdminCropTypesListNameLabel),
		"list_scientific_name_label":   translator.T(c, translator.PagesAdminCropTypesListScientificNameLabel),
		"list_actions_label":           translator.T(c, translator.PagesAdminCropTypesListActionsLabel),
		"list_actions_view_label":      translator.T(c, translator.PagesAdminCropTypesListActionsViewLabel),
		"list_actions_edit_label":      translator.T(c, translator.PagesAdminCropTypesListActionsEditLabel),
		"list_actions_remove_label":    translator.T(c, translator.PagesAdminCropTypesListActionsRemoveLabel),
		"list_next_label":              translator.T(c, translator.PagesAdminCropTypesListNextLabel),
		"list_previous_label":          translator.T(c, translator.PagesAdminCropTypesListPreviousLabel),
	}
}
