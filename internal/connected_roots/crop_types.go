package connected_roots

import (
	"github.com/Kortivex/connected_roots/pkg/pagination"
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
	ImageURL       string    `json:"image_url"`
	Description    string    `json:"description"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CropTypeFilters struct {
	Name           []string `query:"name[]"`
	ScientificName []string `query:"scientific_name[]"`
	PlantingSeason []string `query:"planting_season[]"`
	HarvestSeason  []string `query:"harvest_season[]"`
}

type CropTypePaginationFilters struct {
	pagination.PaginatorParams
	CropTypeFilters
}
