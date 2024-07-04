package role

import "github.com/Kortivex/connected_roots/internal/connected_roots/postgresql"

type Roles struct {
	ID          string `gorm:"column:id;type:varchar(26);primaryKey;not null"`
	Name        string `gorm:"column:name;type:varchar(50)"`
	Description string `gorm:"column:description;type:varchar(255)"`
	Protected   bool   `gorm:"column:protected;type:boolean"`
	postgresql.BaseModel
}
