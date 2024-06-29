package connected_roots

import "time"

type Users struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	Telephone string    `json:"telephone"`
	Language  string    `json:"language"`
	Role      *Roles    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UsersAuthentication struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Valid    bool   `json:"valid,omitempty"`
}
