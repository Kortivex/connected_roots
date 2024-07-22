package sensor

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/i18n/translator"
	"github.com/labstack/echo/v4"
)

func CommonSensorViewPageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":                  translator.T(c, translator.PagesCommonSiteTitle),
		"title":                       translator.T(c, translator.PagesAdminSensorsViewTitle),
		"active":                      "sensors-management",
		"breadcrumb_sensors":          translator.T(c, translator.BreadcrumbSensors),
		"breadcrumb_view_sensors":     translator.T(c, translator.BreadcrumbViewSensors),
		"view_name_label":             translator.T(c, translator.PagesAdminSensorsViewNameLabel),
		"view_location_label":         translator.T(c, translator.PagesAdminSensorsViewLocationLabel),
		"view_model_number_label":     translator.T(c, translator.PagesAdminSensorsViewModelNumberLabel),
		"view_manufacturer_label":     translator.T(c, translator.PagesAdminSensorsViewManufacturerLabel),
		"view_calibration_date_label": translator.T(c, translator.PagesAdminSensorsViewCalibrationDateLabel),
		"view_battery_life_label":     translator.T(c, translator.PagesAdminSensorsViewBatteryLifeLabel),
		"view_ssid_label":             translator.T(c, translator.PagesAdminSensorsViewSSIDLabel),
		"view_channel_label":          translator.T(c, translator.PagesAdminSensorsViewChannelLabel),
		"view_dns_label":              translator.T(c, translator.PagesAdminSensorsViewDNSLabel),
		"view_ip_label":               translator.T(c, translator.PagesAdminSensorsViewIPLabel),
		"view_subnet_label":           translator.T(c, translator.PagesAdminSensorsViewSubnetLabel),
		"view_mac_label":              translator.T(c, translator.PagesAdminSensorsViewMACLabel),
		"view_status_label":           translator.T(c, translator.PagesAdminSensorsViewStatusLabel),
		"view_online_label":           translator.T(c, translator.PagesAdminSensorsViewOnlineLabel),
		"view_offline_label":          translator.T(c, translator.PagesAdminSensorsViewOfflineLabel),
		"view_firmware_version_label": translator.T(c, translator.PagesAdminSensorsViewFirmwareVersionLabel),
		"view_orchard_name_label":     translator.T(c, translator.PagesAdminSensorsViewOrchardNameLabel),
	}
}

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

func CommonSensorDeletePageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":                  translator.T(c, translator.PagesCommonSiteTitle),
		"title":                       translator.T(c, translator.PagesAdminSensorsDeleteTitle),
		"active":                      "sensors-management",
		"breadcrumb_sensors":          translator.T(c, translator.BreadcrumbSensors),
		"breadcrumb_delete_sensors":   translator.T(c, translator.BreadcrumbDeleteSensors),
		"delete_warning_title":        translator.T(c, translator.PagesAdminSensorsDeleteWarningTitle),
		"delete_warning_message":      translator.T(c, translator.PagesAdminSensorsDeleteWarningMessage),
		"delete_button_submit_sensor": translator.T(c, translator.PagesAdminSensorsDeleteButtonSubmitSensor),
	}
}
