package sdk_models

import "time"

type ActivitiesBody struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	OrchardID   string        `json:"orchard_id"`
	Orchard     *OrchardsBody `json:"orchard,omitempty"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

type ActivitiesResponse struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	OrchardID   string            `json:"orchard_id"`
	Orchard     *OrchardsResponse `json:"orchard,omitempty"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

func (ct *ActivitiesResponse) ToActivityBody() *ActivitiesBody {
	return &ActivitiesBody{
		ID:          ct.ID,
		Name:        ct.Name,
		Description: ct.Description,
		OrchardID:   ct.OrchardID,
		Orchard:     ct.Orchard.ToOrchardBody(),
		CreatedAt:   ct.CreatedAt,
		UpdatedAt:   ct.UpdatedAt,
	}
}

type TotalActivitiesResponse struct {
	Total int64 `json:"total"`
}
