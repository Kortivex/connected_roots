package activity

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql"
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql/orchard"
)

type Activities struct {
	ID          string            `gorm:"column:id;type:varchar(26);primaryKey;not null"`
	Name        string            `gorm:"column:name;type:varchar(100)"`
	Description string            `gorm:"column:description;type:text"`
	OrchardID   string            `gorm:"column:orchard_id;type:varchar(26)"`
	Orchard     *orchard.Orchards `gorm:"foreignKey:OrchardID;references:ID"`
	postgresql.BaseModel
}

func (a *Activities) TableName() string {
	return "agricultural_activities"
}
