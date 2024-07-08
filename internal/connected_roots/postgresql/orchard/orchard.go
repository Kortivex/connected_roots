package orchard

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql"
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql/crop_type"
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql/user"
)

type Orchards struct {
	ID         string               `gorm:"column:id;type:varchar(26);primaryKey;not null"`
	Name       string               `gorm:"column:name;type:varchar(100)"`
	Location   string               `gorm:"column:location;type:varchar(255)"`
	Size       float64              `gorm:"column:size;type:numeric"`
	Soil       string               `gorm:"column:soil;type:varchar(255)"`
	Fertilizer string               `gorm:"column:fertilizer;type:varchar(255)"`
	Composting string               `gorm:"column:composting;type:varchar(255)"`
	UserID     string               `gorm:"column:user_id;type:varchar(100)"`
	User       *user.Users          `gorm:"foreignKey:UserID;references:ID"`
	CropTypeID string               `gorm:"column:crop_type_id;type:varchar(26)"`
	CropType   *crop_type.CropTypes `gorm:"foreignKey:CropTypeID;references:ID"`
	postgresql.BaseModel
}

func (o *Orchards) TableName() string {
	return "orchards"
}
