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
