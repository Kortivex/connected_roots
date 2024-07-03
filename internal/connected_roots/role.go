package connected_roots

import (
	"github.com/Kortivex/connected_roots/pkg/pagination"
	"time"
)

type Roles struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RoleFilters struct {
	Name []string `query:"name[]"`
}

type RolePaginationFilters struct {
	pagination.PaginatorParams
	RoleFilters
}
