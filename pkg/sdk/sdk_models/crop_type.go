package sdk_models

import "time"

type CropTypesBody struct {
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

type CropTypesResponse struct {
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

func (ct *CropTypesResponse) ToCropTypeBody() *CropTypesBody {
	return &CropTypesBody{
		ID:             ct.ID,
		Name:           ct.Name,
		ScientificName: ct.ScientificName,
		LifeCycle:      ct.LifeCycle,
		PlantingSeason: ct.PlantingSeason,
		HarvestSeason:  ct.HarvestSeason,
		Irrigation:     ct.Irrigation,
		Description:    ct.Description,
		CreatedAt:      ct.CreatedAt,
		UpdatedAt:      ct.UpdatedAt,
	}
}
