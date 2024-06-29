package sdk_models

import (
	"time"
)

type UsersResponse struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Surname   string         `json:"surname"`
	Email     string         `json:"email"`
	Password  string         `json:"password,omitempty"`
	Telephone string         `json:"telephone"`
	Language  string         `json:"language"`
	Role      *RolesResponse `json:"role,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type UsersAuthenticationBody struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Valid    bool   `json:"valid,omitempty"`
}

type UsersAuthenticationResponse struct {
	Valid bool `json:"valid"`
}
