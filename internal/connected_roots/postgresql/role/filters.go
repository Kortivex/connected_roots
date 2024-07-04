package role

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
	"gorm.io/gorm"
)

const TableRolesName = "roles."

var (
	TableRolesFields  = []string{"id", "name", "description", "protected", "created_at", "updated_at", "deleted_at"}
	TableRolesSortMap = map[string]string{
		"id":          "ID",
		"name":        "Name",
		"description": "Description",
		"protected":   "Protected",
		"created_at":  "CreatedAt",
		"updated_at":  "UpdatedAt",
		"deleted_at":  "DeletedAt",
	}
	DefaultRoleRule = paginator.Rule{
		Key:     "ID",
		SQLRepr: TableRolesName + "id",
	}
)

func AddRoleFilters(db *gorm.DB, filters *connected_roots.RoleFilters) {
	if len(filters.Name) > 0 {
		db.Where(TableRolesName+"name IN ?", filters.Name)
	}
}
