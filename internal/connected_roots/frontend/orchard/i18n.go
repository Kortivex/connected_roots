package orchard

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/i18n/translator"
	"github.com/labstack/echo/v4"
)

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

func CommonOrchardListPageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":                 translator.T(c, translator.PagesCommonSiteTitle),
		"title":                      translator.T(c, translator.PagesAdminOrchardsListTitle),
		"active":                     "orchards-management",
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
