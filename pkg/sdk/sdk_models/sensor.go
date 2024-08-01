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

type SensorsDataResponse struct {
	ID             string           `json:"id"`
	Voltage        float64          `json:"voltage"`
	Battery        float64          `json:"battery"`
	Soil           int              `json:"soil"`
	Salt           int              `json:"salt"`
	Light          float64          `json:"light"`
	TemperatureIn  float64          `json:"temperature_in"`
	TemperatureOut float64          `json:"temperature_out"`
	HumidityIn     float64          `json:"humidity_in"`
	HumidityOut    float64          `json:"humidity_out"`
	Pressure       float64          `json:"pressure"`
	Altitude       float64          `json:"altitude"`
	SensorID       string           `json:"sensor_id"`
	Sensor         *SensorsResponse `json:"sensor,omitempty"`
	CreatedAt      time.Time        `json:"created_at"`
}

type SensorsDataWeekdayAverageResponse struct {
	Weekday           int     `json:"weekday"`
	AvgVoltage        float64 `json:"avg_voltage"`
	AvgBattery        float64 `json:"avg_battery"`
	AvgSoil           float64 `json:"avg_soil"`
	AvgSalt           float64 `json:"avg_salt"`
	AvgLight          float64 `json:"avg_light"`
	AvgTemperatureIn  float64 `json:"avg_temperature_in"`
	AvgTemperatureOut float64 `json:"avg_temperature_out"`
	AvgHumidityIn     float64 `json:"avg_humidity_in"`
	AvgHumidityOut    float64 `json:"avg_humidity_out"`
	AvgPressure       float64 `json:"avg_pressure"`
	AvgAltitude       float64 `json:"avg_altitude"`
}
