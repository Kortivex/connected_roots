package orchard

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
	"gorm.io/gorm"
)

const TableOrchardsName = "orchards."

var (
	TableOrchardsFields  = []string{"id", "name", "location", "size", "size", "soil", "fertilizer", "composting", "created_at", "updated_at", "deleted_at"}
	TableOrchardsSortMap = map[string]string{
		"id":         "ID",
		"name":       "Name",
		"location":   "Location",
		"size":       "Size",
		"soil":       "Soil",
		"fertilizer": "Fertilizer",
		"composting": "Composting",
		"created_at": "CreatedAt",
		"updated_at": "UpdatedAt",
		"deleted_at": "DeletedAt",
	}
	DefaultOrchardRule = paginator.Rule{
		Key:     "ID",
		SQLRepr: TableOrchardsName + "id",
	}
)

func AddOrchardFilters(db *gorm.DB, filters *connected_roots.OrchardFilters) {
	if len(filters.Name) > 0 {
		db.Where(TableOrchardsName+"name IN ?", filters.Name)
	}
	if len(filters.Location) > 0 {
		db.Where(TableOrchardsName+"location IN ?", filters.Location)
	}
	if len(filters.UserID) > 0 {
		db.Where(TableOrchardsName+"user_id IN ?", filters.UserID)
	}
}
