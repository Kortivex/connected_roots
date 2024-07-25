package activity

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
	"gorm.io/gorm"
)

const TableActivitiesName = "agricultural_activities."

var (
	TableActivitiesFields  = []string{"id", "name", "description", "created_at", "updated_at", "deleted_at"}
	TableActivitiesSortMap = map[string]string{
		"id":          "ID",
		"name":        "Name",
		"description": "Description",
		"composting":  "Composting",
		"created_at":  "CreatedAt",
		"updated_at":  "UpdatedAt",
		"deleted_at":  "DeletedAt",
	}
	DefaultActivityRule = paginator.Rule{
		Key:     "ID",
		SQLRepr: TableActivitiesName + "id",
	}
)

func AddActivityFilters(db *gorm.DB, filters *connected_roots.ActivityFilters) {
	if len(filters.Name) > 0 {
		db.Where(TableActivitiesName+"name IN ?", filters.Name)
	}
	if len(filters.OrchardID) > 0 {
		db.Where(TableActivitiesName+"orchard_id IN ?", filters.OrchardID)
	}
}
