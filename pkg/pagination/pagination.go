package pagination

import (
	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
)

type Pagination struct {
	Data   any    `json:"data"`
	Paging Paging `json:"paging"`
}

type Paging struct {
	NextCursor     string `json:"next_cursor"`
	PreviousCursor string `json:"previous_cursor"`
}

type PaginatorParams struct {
	Limit int             `query:"limit"`
	Order paginator.Order `query:"order"`

	NextCursor     string   `query:"next_cursor"`
	PreviousCursor string   `query:"previous_cursor"`
	Sort           []string `query:"sort"`
}

func CreatePaginator(queryParams *PaginatorParams, rules []paginator.Rule) (*paginator.Paginator, error) {
	if err := getPaginatorParams(queryParams); err != nil {
		return nil, err
	}

	p := paginator.New(
		&paginator.Config{
			After:  queryParams.NextCursor,
			Before: queryParams.PreviousCursor,
			Rules:  rules,
			Limit:  queryParams.Limit,
			Order:  queryParams.Order,
		},
	)
	return p, nil
}
