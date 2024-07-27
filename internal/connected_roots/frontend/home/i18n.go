package home

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/i18n/translator"
	"github.com/labstack/echo/v4"
)

func CommonHomePageI18N(c echo.Context) map[string]interface{} {
	return map[string]interface{}{
		"site_title":                       translator.T(c, translator.PagesCommonSiteTitle),
		"title":                            translator.T(c, translator.PagesAdminHomeTitle),
		"active":                           "home",
		"home_total_active_sessions_label": translator.T(c, translator.PagesAdminHomeTotalActiveSessionsLabel),
		"home_total_users_label":           translator.T(c, translator.PagesAdminHomeTotalUsersLabel),
		"home_total_orchards_label":        translator.T(c, translator.PagesAdminHomeTotalOrchardsLabel),
		"home_total_sensors_label":         translator.T(c, translator.PagesAdminHomeTotalSensorsLabel),
		"home_total_crops_label":           translator.T(c, translator.PagesAdminHomeTotalCropsLabel),
		"home_total_activities_label":      translator.T(c, translator.PagesAdminHomeTotalActivitiesLabel),
	}
}
