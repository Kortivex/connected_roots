package connected_roots

import (
	"github.com/Kortivex/connected_roots/pkg/pagination"
	"time"
)

type Activities struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OrchardID   string    `json:"orchard_id"`
	Orchard     *Orchards `json:"orchard,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ActivityFilters struct {
	Name      []string `query:"name[]"`
	OrchardID []string `query:"orchard_id[]"`
}

type ActivityPaginationFilters struct {
	pagination.PaginatorParams
	ActivityFilters
}
