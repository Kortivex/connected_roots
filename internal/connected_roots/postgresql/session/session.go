package session

import (
	"time"
)

type Sessions struct {
	ID        string    `gorm:"column:id;type:text;primaryKey;not null"`
	Data      string    `gorm:"column:dat;type:text"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;autoUpdateTime"`
	ExpiredAt time.Time `gorm:"column:expired_at;type:timestamp;"`
}

// TableName returns the name of the table associated with the struct.
func (s *Sessions) TableName() string {
	return "sessions"
}
