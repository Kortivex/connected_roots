package session

import "github.com/Kortivex/connected_roots/internal/connected_roots/postgresql"

type Sessions struct {
	ID   string `gorm:"column:id;type:text;primaryKey;not null"`
	Data string `gorm:"column:dat;type:text"`
	postgresql.BaseModel
}

// TableName returns the name of the table associated with the struct.
func (s *Sessions) TableName() string {
	return "sessions"
}
