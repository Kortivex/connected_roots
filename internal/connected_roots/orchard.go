package connected_roots

import (
	"github.com/Kortivex/connected_roots/pkg/pagination"
	"time"
)

type Orchards struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Location   string     `json:"location"`
	Size       float64    `json:"size"`
	Soil       string     `json:"soil"`
	Fertilizer string     `json:"fertilizer"`
	Composting string     `json:"composting"`
	UserID     string     `json:"user_id"`
	User       *Users     `json:"user"`
	CropTypeID string     `json:"crop_type_id"`
	CropType   *CropTypes `json:"crop_type"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type OrchardFilters struct {
	Name     []string `query:"name[]"`
	Location []string `query:"location[]"`
	UserID   []string `query:"user_id[]"`
}

type OrchardPaginationFilters struct {
	pagination.PaginatorParams
	OrchardFilters
}
