package sensor

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql"
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql/orchard"
	"time"
)

type Sensors struct {
	ID              string            `gorm:"column:id;type:varchar(26);primaryKey;not null"`
	Name            string            `gorm:"column:name;type:varchar(255)"`
	Location        string            `gorm:"column:location;type:varchar(255)"`
	ModelNumber     string            `gorm:"column:model_number;type:varchar(255)"`
	Manufacturer    string            `gorm:"column:manufacturer;type:varchar(255)"`
	CalibrationDate time.Time         `gorm:"column:calibration_date;type:timestamp"`
	BatteryLife     float64           `gorm:"column:battery_life;type:decimal"`
	SSID            string            `gorm:"column:ssid;type:varchar(255)"`
	Channel         string            `gorm:"column:channel;type:varchar(255)"`
	DNS             string            `gorm:"column:dns;type:varchar(255)"`
	IP              string            `gorm:"column:ip;type:varchar(255)"`
	Subnet          string            `gorm:"column:subnet;type:varchar(255)"`
	MAC             string            `gorm:"column:mac;type:varchar(255)"`
	Status          int               `gorm:"column:status;type:int"`
	FirmwareVersion float64           `gorm:"column:firmware_version;type:decimal"`
	OrchardID       string            `gorm:"column:orchard_id;type:varchar(26)"`
	Orchard         *orchard.Orchards `gorm:"foreignKey:OrchardID;references:ID"`
	postgresql.BaseModel
}

func (s *Sensors) TableName() string {
	return "sensors"
}

type SensorsData struct {
	ID             string   `gorm:"column:id;type:varchar(26);primaryKey;not null"`
	Voltage        float64  `gorm:"column:voltage;type:double precision"`
	Battery        float64  `gorm:"column:battery;type:double precision"`
	Soil           int      `gorm:"column:soil;type:int"`
	Salt           int      `gorm:"column:salt;type:int"`
	Light          float64  `gorm:"column:light;type:double precision"`
	TemperatureIn  float64  `gorm:"column:temperature_in;type:double precision"`
	TemperatureOut float64  `gorm:"column:temperature_out;type:double precision"`
	HumidityIn     float64  `gorm:"column:humidity_in;type:double precision"`
	HumidityOut    float64  `gorm:"column:humidity_out;type:double precision"`
	Pressure       float64  `gorm:"column:pressure;type:double precision"`
	Altitude       float64  `gorm:"column:altitude;type:double precision"`
	SensorID       string   `gorm:"column:sensor_id;type:varchar(26)"`
	Sensor         *Sensors `gorm:"foreignKey:SensorID;references:ID"`
	postgresql.BaseModel
}

func (sd *SensorsData) TableName() string {
	return "sensor_data"
}

type SensorsDataWeekdayAverage struct {
	Weekday           int     `gorm:"column:weekday"`
	AvgVoltage        float64 `gorm:"column:avg_voltage"`
	AvgBattery        float64 `gorm:"column:avg_battery"`
	AvgSoil           float64 `gorm:"column:avg_soil"`
	AvgSalt           float64 `gorm:"column:avg_salt"`
	AvgLight          float64 `gorm:"column:avg_light"`
	AvgTemperatureIn  float64 `gorm:"column:avg_temperature_in"`
	AvgTemperatureOut float64 `gorm:"column:avg_temperature_out"`
	AvgHumidityIn     float64 `gorm:"column:avg_humidity_in"`
	AvgHumidityOut    float64 `gorm:"column:avg_humidity_out"`
	AvgPressure       float64 `gorm:"column:avg_pressure"`
	AvgAltitude       float64 `gorm:"column:avg_altitude"`
}

func (sd *SensorsDataWeekdayAverage) TableName() string {
	return "sensor_data"
}
