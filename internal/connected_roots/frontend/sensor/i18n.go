package sensor

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/i18n/translator"
	"github.com/labstack/echo/v4"
)

func CommonSensorListPageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":                    translator.T(c, translator.PagesCommonSiteTitle),
		"title":                         translator.T(c, translator.PagesAdminSensorsListTitle),
		"active":                        "sensors-management",
		"breadcrumb_sensors":            translator.T(c, translator.BreadcrumbSensors),
		"breadcrumb_list_sensors":       translator.T(c, translator.BreadcrumbListSensors),
		"list_button_create_sensor":     translator.T(c, translator.PagesAdminSensorsListButtonCreateSensor),
		"list_name_label":               translator.T(c, translator.PagesAdminSensorsListNameLabel),
		"list_model_number_label":       translator.T(c, translator.PagesAdminSensorsListModelNumberLabel),
		"list_battery_level_label":      translator.T(c, translator.PagesAdminSensorsListBatteryLevelLabel),
		"list_status_label":             translator.T(c, translator.PagesAdminSensorsListStatusLabel),
		"list_firmware_version_label":   translator.T(c, translator.PagesAdminSensorsListFirmwareVersionLabel),
		"list_orchard_id_version_label": translator.T(c, translator.PagesAdminSensorsListOrchardIDLabel),
		"list_actions_label":            translator.T(c, translator.PagesAdminSensorsListActionsLabel),
		"list_actions_view_label":       translator.T(c, translator.PagesAdminSensorsListActionsViewLabel),
		"list_actions_edit_label":       translator.T(c, translator.PagesAdminSensorsListActionsEditLabel),
		"list_actions_remove_label":     translator.T(c, translator.PagesAdminSensorsListActionsRemoveLabel),
		"list_next_label":               translator.T(c, translator.PagesAdminSensorsListNextLabel),
		"list_previous_label":           translator.T(c, translator.PagesAdminSensorsListPreviousLabel),
	}
}
