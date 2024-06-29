package user

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql"
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql/role"
)

type Users struct {
	ID        string      `gorm:"column:id;type:varchar(26);primaryKey;not null"`
	Name      string      `gorm:"column:name;type:varchar(100)"`
	Surname   string      `gorm:"column:surname;type:varchar(100)"`
	Email     string      `gorm:"column:email;type:varchar(255);unique;not null"`
	Password  string      `gorm:"column:password;type:varchar(255);not null"`
	Telephone string      `gorm:"column:telephone;type:varchar(30)"`
	Language  string      `gorm:"column:language;type:varchar(3)"`
	RoleID    string      `gorm:"column:role_id;type:varchar(26);not null"`
	Role      *role.Roles `gorm:"foreignKey:RoleID;references:ID"`
	postgresql.BaseModel
}

// TableName returns the name of the table associated with the struct.
func (s *Users) TableName() string {
	return "users"
}
