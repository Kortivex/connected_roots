package connected_roots

import (
	"github.com/Kortivex/connected_roots/pkg/pagination"
	"time"
)

type Sensors struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Location        string    `json:"location"`
	ModelNumber     string    `json:"model_number"`
	Manufacturer    string    `json:"manufacturer"`
	CalibrationDate time.Time `json:"calibration_date"`
	BatteryLife     float64   `json:"battery_life"`
	SSID            string    `json:"ssid"`
	Channel         string    `json:"channel"`
	DNS             string    `json:"dns"`
	IP              string    `json:"ip"`
	Subnet          string    `json:"subnet"`
	MAC             string    `json:"mac"`
	Status          int       `json:"status"`
	FirmwareVersion float64   `json:"firmware_version"`
	OrchardID       string    `json:"orchard_id"`
	Orchard         *Orchards `json:"orchard,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type SensorFilters struct {
	Name            []string `query:"name[]"`
	FirmwareVersion []string `query:"firmware_version[]"`
	Manufacturer    []string `query:"manufacturer[]"`
	BatteryLife     []string `query:"battery_life[]"`
	Status          []string `query:"status[]"`
}

type SensorPaginationFilters struct {
	pagination.PaginatorParams
	SensorFilters
}

type SensorsData struct {
	ID             string    `json:"id"`
	Voltage        float64   `json:"voltage"`
	Battery        float64   `json:"battery"`
	Soil           int       `json:"soil"`
	Salt           int       `json:"salt"`
	Light          float64   `json:"light"`
	TemperatureIn  float64   `json:"temperature_in"`
	TemperatureOut float64   `json:"temperature_out"`
	HumidityIn     float64   `json:"humidity_in"`
	HumidityOut    float64   `json:"humidity_out"`
	Pressure       float64   `json:"pressure"`
	Altitude       float64   `json:"altitude"`
	SensorID       string    `json:"sensor_id"`
	Sensor         *Sensors  `json:"sensor,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
}

type SensorDataFilters struct {
	ID             []string `query:"id[]"`
	Voltage        []string `query:"voltage[]"`
	Battery        []string `query:"battery[]"`
	Soil           []string `query:"soil[]"`
	Salt           []string `query:"salt[]"`
	Light          []string `query:"light[]"`
	TemperatureIn  []string `query:"temperature_in[]"`
	TemperatureOut []string `query:"temperature_out[]"`
	HumidityIn     []string `query:"humidity_in[]"`
	HumidityOut    []string `query:"humidity_out[]"`
	Pressure       []string `query:"pressure[]"`
	Altitude       []string `query:"altitude[]"`
	SensorID       []string `query:"sensor_id[]"`
}

type SensorDataPaginationFilters struct {
	pagination.PaginatorParams
	SensorDataFilters
}

type TotalSensors struct {
	Total int64 `json:"total"`
}
