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

func AddSensorFilters(db *gorm.DB, filters *connected_roots.SensorFilters) {
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

const TableSensorsDataName = "sensor_data."

var (
	TableSensorsDataFields  = []string{"id", "voltage", "battery", "soil", "salt", "light", "temperature_in", "temperature_out", "humidity_in", "humidity_out", "pressure", "altitude", "created_at", "updated_at", "deleted_at"}
	TableSensorsDataSortMap = map[string]string{
		"id":              "ID",
		"voltage":         "Voltage",
		"battery":         "Battery",
		"soil":            "Soil",
		"salt":            "Salt",
		"light":           "Light",
		"temperature_in":  "TemperatureIn",
		"temperature_out": "TemperatureOut",
		"humidity_in":     "HumidityIn",
		"humidity_out":    "HumidityOut",
		"pressure":        "Pressure",
		"altitude":        "Altitude",
		"created_at":      "CreatedAt",
		"updated_at":      "UpdatedAt",
		"deleted_at":      "DeletedAt",
	}
	DefaultSensorDataRule = paginator.Rule{
		Key:     "ID",
		SQLRepr: TableSensorsDataName + "id",
	}
)

func AddSensorDataFilters(db *gorm.DB, filters *connected_roots.SensorDataFilters) {
	if len(filters.ID) > 0 {
		db.Where(TableSensorsName+"id IN ?", filters.ID)
	}
	if len(filters.Voltage) > 0 {
		db.Where(TableSensorsName+"voltage IN ?", filters.Voltage)
	}
	if len(filters.Battery) > 0 {
		db.Where(TableSensorsName+"battery IN ?", filters.Battery)
	}
	if len(filters.Soil) > 0 {
		db.Where(TableSensorsName+"soil IN ?", filters.Soil)
	}
	if len(filters.Salt) > 0 {
		db.Where(TableSensorsName+"salt IN ?", filters.Salt)
	}
	if len(filters.Light) > 0 {
		db.Where(TableSensorsName+"light IN ?", filters.Light)
	}
	if len(filters.TemperatureIn) > 0 {
		db.Where(TableSensorsName+"temperature_in IN ?", filters.TemperatureIn)
	}
	if len(filters.TemperatureOut) > 0 {
		db.Where(TableSensorsName+"temperature_out IN ?", filters.TemperatureOut)
	}
	if len(filters.HumidityIn) > 0 {
		db.Where(TableSensorsName+"humidity_in IN ?", filters.HumidityIn)
	}
	if len(filters.HumidityOut) > 0 {
		db.Where(TableSensorsName+"humidity_out IN ?", filters.HumidityOut)
	}
	if len(filters.Pressure) > 0 {
		db.Where(TableSensorsName+"pressure IN ?", filters.Pressure)
	}
	if len(filters.Altitude) > 0 {
		db.Where(TableSensorsName+"altitude IN ?", filters.Altitude)
	}
}
