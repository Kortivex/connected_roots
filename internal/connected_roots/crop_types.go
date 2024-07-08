package connected_roots

import (
	"time"
)

type CropTypes struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	ScientificName string    `json:"scientific_name"`
	LifeCycle      string    `json:"life_cycle"`
	PlantingSeason string    `json:"planting_season"`
	HarvestSeason  string    `json:"harvest_season"`
	Irrigation     string    `json:"irrigation"`
	Description    string    `json:"description"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
