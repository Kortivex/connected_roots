package crop_type

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
	"gorm.io/gorm"
)

const TableCropTypesName = "crop_types."

var (
	TableCropTypesFields  = []string{"id", "name", "scientific_name", "life_cycle", "planting_season", "harvest_season", "irrigation", "description", "created_at", "updated_at", "deleted_at"}
	TableCropTypesSortMap = map[string]string{
		"id":              "ID",
		"name":            "Name",
		"scientific_name": "ScientificName",
		"life_cycle":      "LifeCycle",
		"planting_season": "PlantingSeason",
		"harvest_season":  "HarvestSeason",
		"irrigation":      "Irrigation",
		"description":     "Description",
		"created_at":      "CreatedAt",
		"updated_at":      "UpdatedAt",
		"deleted_at":      "DeletedAt",
	}
	DefaultCropTypeRule = paginator.Rule{
		Key:     "ID",
		SQLRepr: TableCropTypesName + "id",
	}
)

func AddCropTypeFilters(db *gorm.DB, filters *connected_roots.CropTypeFilters) {
	if len(filters.Name) > 0 {
		db.Where(TableCropTypesName+"name IN ?", filters.Name)
	}

	if len(filters.ScientificName) > 0 {
		db.Where(TableCropTypesName+"scientific_name IN ?", filters.ScientificName)
	}

	if len(filters.PlantingSeason) > 0 {
		db.Where(TableCropTypesName+"planting_season IN ?", filters.PlantingSeason)
	}

	if len(filters.HarvestSeason) > 0 {
		db.Where(TableCropTypesName+"harvest_season IN ?", filters.HarvestSeason)
	}
}
