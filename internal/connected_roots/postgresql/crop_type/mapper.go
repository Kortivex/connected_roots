package crop_type

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots"
)

func toDomain(cropType *CropTypes) *connected_roots.CropTypes {
	return &connected_roots.CropTypes{
		ID:             cropType.ID,
		Name:           cropType.Name,
		ScientificName: cropType.ScientificName,
		LifeCycle:      cropType.LifeCycle,
		PlantingSeason: cropType.PlantingSeason,
		HarvestSeason:  cropType.HarvestSeason,
		Irrigation:     cropType.Irrigation,
		Description:    cropType.Description,
		CreatedAt:      cropType.CreatedAt,
		UpdatedAt:      cropType.UpdatedAt,
	}
}

func toDomainSlice(cropTypes []*CropTypes) []*connected_roots.CropTypes {
	cropTypesDomain := []*connected_roots.CropTypes{}
	for _, cropType := range cropTypes {
		cropTypeDomain := toDomain(cropType)
		cropTypesDomain = append(cropTypesDomain, cropTypeDomain)
	}
	return cropTypesDomain
}

func toDB(cropType *connected_roots.CropTypes, id string) *CropTypes {
	return &CropTypes{
		ID:             id,
		Name:           cropType.Name,
		ScientificName: cropType.ScientificName,
		LifeCycle:      cropType.LifeCycle,
		PlantingSeason: cropType.PlantingSeason,
		HarvestSeason:  cropType.HarvestSeason,
		Irrigation:     cropType.Irrigation,
		Description:    cropType.Description,
	}
}
