package sdk_models

import "time"

type RolesBody struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Protected   bool      `json:"protected"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RolesResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Protected   bool      `json:"protected"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (rl *RolesResponse) ToRoleBody() *RolesBody {
	return &RolesBody{
		ID:          rl.ID,
		Name:        rl.Name,
		Description: rl.Description,
		Protected:   rl.Protected,
		CreatedAt:   rl.CreatedAt,
		UpdatedAt:   rl.UpdatedAt,
	}
}
