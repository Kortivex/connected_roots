package crop_type

import "github.com/Kortivex/connected_roots/internal/connected_roots/postgresql"

type CropTypes struct {
	ID             string `gorm:"column:id;type:varchar(26);primaryKey;not null"`
	Name           string `gorm:"column:name;type:varchar(100)"`
	ScientificName string `gorm:"column:scientific_name;type:varchar(100)"`
	LifeCycle      string `gorm:"column:life_cycle;type:varchar(100)"`
	PlantingSeason string `gorm:"column:planting_season;type:varchar(100)"`
	HarvestSeason  string `gorm:"column:harvest_season;type:varchar(100)"`
	Irrigation     string `gorm:"column:irrigation;type:varchar(100)"`
	Description    string `gorm:"column:description;type:text"`
	postgresql.BaseModel
}

func (s *CropTypes) TableName() string {
	return "crop_types"
}
