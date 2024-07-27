package connected_roots

import (
	"github.com/Kortivex/connected_roots/pkg/pagination"
	"time"
)

type Users struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	Telephone string    `json:"telephone"`
	Language  string    `json:"language"`
	RoleID    string    `json:"role_id,omitempty"`
	Role      *Roles    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UsersAuthentication struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Valid    bool   `json:"valid,omitempty"`
}

type UserFilters struct {
	Name    []string `query:"name[]"`
	Surname []string `query:"surname[]"`
	Email   []string `query:"email[]"`
}

type UserPaginationFilters struct {
	pagination.PaginatorParams
	UserFilters
}

type TotalUsers struct {
	Total int64 `json:"total"`
}
