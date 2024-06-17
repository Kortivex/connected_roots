package postgresql

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;"`
}
