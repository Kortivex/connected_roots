package sensor

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
	"gorm.io/gorm"
)

const TableSensorsName = "sensors."

var (
	TableSensorsFields  = []string{"id", "name", "location", "model_number", "manufacturer", "calibration_date", "battery_life", "ssid", "channel", "dns", "ip", "subnet", "mac", "status", "firmware_version", "created_at", "updated_at", "deleted_at"}
	TableSensorsSortMap = map[string]string{
		"id":               "ID",
		"name":             "Name",
		"location":         "Location",
		"model_number":     "ModelNumber",
		"manufacturer":     "Manufacturer",
		"calibration_date": "CalibrationDate",
		"battery_life":     "BatteryLife",
		"ssid":             "SSID",
		"channel":          "Channel",
		"dns":              "DNS",
		"ip":               "IP",
		"subnet":           "Subnet",
		"mac":              "MAC",
		"status":           "Status",
		"firmware_version": "FirmwareVersion",
		"created_at":       "CreatedAt",
		"updated_at":       "UpdatedAt",
		"deleted_at":       "DeletedAt",
	}
	DefaultSensorRule = paginator.Rule{
		Key:     "ID",
		SQLRepr: TableSensorsName + "id",
	}
)

func AddOrchardFilters(db *gorm.DB, filters *connected_roots.SensorFilters) {
	if len(filters.Name) > 0 {
		db.Where(TableSensorsName+"name IN ?", filters.Name)
	}
	if len(filters.FirmwareVersion) > 0 {
		db.Where(TableSensorsName+"firmware_version IN ?", filters.FirmwareVersion)
	}
	if len(filters.Manufacturer) > 0 {
		db.Where(TableSensorsName+"manufacturer IN ?", filters.Manufacturer)
	}
	if len(filters.BatteryLife) > 0 {
		db.Where(TableSensorsName+"battery_life IN ?", filters.BatteryLife)
	}
	if len(filters.Status) > 0 {
		db.Where(TableSensorsName+"status IN ?", filters.Status)
	}
}
