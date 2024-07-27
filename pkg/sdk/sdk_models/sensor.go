package sdk_models

import "time"

type SensorsBody struct {
	ID              string        `json:"id"`
	Name            string        `json:"name"`
	Location        string        `json:"location"`
	ModelNumber     string        `json:"model_number"`
	Manufacturer    string        `json:"manufacturer"`
	CalibrationDate time.Time     `json:"calibration_date"`
	BatteryLife     float64       `json:"battery_life"`
	SSID            string        `json:"ssid"`
	Channel         string        `json:"channel"`
	DNS             string        `json:"dns"`
	IP              string        `json:"ip"`
	Subnet          string        `json:"subnet"`
	MAC             string        `json:"mac"`
	Status          int           `json:"status"`
	FirmwareVersion float64       `json:"firmware_version"`
	OrchardID       string        `json:"orchard_id"`
	Orchard         *OrchardsBody `json:"orchard,omitempty"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}

type SensorsResponse struct {
	ID              string            `json:"id"`
	Name            string            `json:"name"`
	Location        string            `json:"location"`
	ModelNumber     string            `json:"model_number"`
	Manufacturer    string            `json:"manufacturer"`
	CalibrationDate time.Time         `json:"calibration_date"`
	BatteryLife     float64           `json:"battery_life"`
	SSID            string            `json:"ssid"`
	Channel         string            `json:"channel"`
	DNS             string            `json:"dns"`
	IP              string            `json:"ip"`
	Subnet          string            `json:"subnet"`
	MAC             string            `json:"mac"`
	Status          int               `json:"status"`
	FirmwareVersion float64           `json:"firmware_version"`
	OrchardID       string            `json:"orchard_id"`
	Orchard         *OrchardsResponse `json:"orchard,omitempty"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
}

func (or *SensorsResponse) ToSensorBody() *SensorsBody {
	return &SensorsBody{
		ID:              or.ID,
		Name:            or.Name,
		Location:        or.Location,
		ModelNumber:     or.ModelNumber,
		Manufacturer:    or.Manufacturer,
		CalibrationDate: or.CalibrationDate,
		BatteryLife:     or.BatteryLife,
		SSID:            or.SSID,
		Channel:         or.Channel,
		DNS:             or.DNS,
		IP:              or.IP,
		Subnet:          or.Subnet,
		MAC:             or.MAC,
		Status:          or.Status,
		FirmwareVersion: or.FirmwareVersion,
		OrchardID:       or.OrchardID,
		Orchard:         or.Orchard.ToOrchardBody(),
		CreatedAt:       or.CreatedAt,
		UpdatedAt:       or.UpdatedAt,
	}
}

type TotalSensorsResponse struct {
	Total int64 `json:"total"`
}
