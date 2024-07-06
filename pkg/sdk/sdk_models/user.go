package sdk_models

import (
	"time"
)

type UsersBody struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Surname   string     `json:"surname"`
	Email     string     `json:"email"`
	Password  string     `json:"password,omitempty"`
	Telephone string     `json:"telephone"`
	Language  string     `json:"language"`
	RoleID    string     `json:"role_id,omitempty"`
	Role      *RolesBody `json:"role,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type UsersResponse struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Surname   string         `json:"surname"`
	Email     string         `json:"email"`
	Password  string         `json:"password,omitempty"`
	Telephone string         `json:"telephone"`
	Language  string         `json:"language"`
	RoleID    string         `json:"role_id,omitempty"`
	Role      *RolesResponse `json:"role,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func (ur *UsersResponse) ToUsersBody() *UsersBody {
	return &UsersBody{
		ID:        ur.ID,
		Name:      ur.Name,
		Surname:   ur.Surname,
		Email:     ur.Email,
		Password:  ur.Password,
		Telephone: ur.Telephone,
		Language:  ur.Language,
		RoleID:    ur.RoleID,
		CreatedAt: ur.CreatedAt,
		UpdatedAt: ur.UpdatedAt,
	}
}

type UsersAuthenticationBody struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Valid    bool   `json:"valid,omitempty"`
}

type UsersAuthenticationResponse struct {
	Valid bool `json:"valid"`
}
