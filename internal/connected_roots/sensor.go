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
