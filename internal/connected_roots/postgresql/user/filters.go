package user

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
	"gorm.io/gorm"
)

const TableUsersName = "users."

var (
	TableUsersFields = []string{"id", "name", "surname", "email", "password", "telephone", "language", "role_id",
		"created_at", "updated_at", "deleted_at"}
	TableUsersSortMap = map[string]string{
		"id":         "ID",
		"name":       "Name",
		"surname":    "Surname",
		"email":      "Email",
		"password":   "Password",
		"telephone":  "Telephone",
		"language":   "Language",
		"created_at": "CreatedAt",
		"updated_at": "UpdatedAt",
		"deleted_at": "DeletedAt",
	}
	DefaultUserRule = paginator.Rule{
		Key:     "ID",
		SQLRepr: TableUsersName + "id",
	}
)

func AddRoleFilters(db *gorm.DB, filters *connected_roots.UserFilters) {
	if len(filters.Name) > 0 {
		db.Where(TableUsersName+"name IN ?", filters.Name)
	}
	if len(filters.Surname) > 0 {
		db.Where(TableUsersName+"surname IN ?", filters.Surname)
	}
	if len(filters.Email) > 0 {
		db.Where(TableUsersName+"email IN ?", filters.Email)
	}
}
