package connected_roots

import (
	"net/http"
)

type Session struct {
	ID       string  `json:"id"`
	Email    string  `json:"email"`
	UserID   string  `json:"user_id"`
	Language string  `json:"language"`
	Role     string  `json:"role"`
	Cookie   *Cookie `json:"cookie,omitempty"`
}

type Cookie struct {
	Path     string        `json:"path"`
	Domain   string        `json:"domain"`
	MaxAge   int           `json:"max_age"`
	Secure   bool          `json:"secure"`
	HTTPOnly bool          `json:"http_only"`
	SameSite http.SameSite `json:"same_site"`
}
