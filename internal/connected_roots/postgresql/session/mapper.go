package session

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/gorilla/sessions"
)

func toDomain(session *sessions.Session) *connected_roots.Session {
	email := ""
	if emailVal, ok := session.Values["email"]; ok {
		email = emailVal.(string)
	}

	userID := ""
	if userIDVal, ok := session.Values["user_id"]; ok {
		userID = userIDVal.(string)
	}

	name := ""
	if nameVal, ok := session.Values["name"]; ok {
		name = nameVal.(string)
	}

	surname := ""
	if surnameVal, ok := session.Values["surname"]; ok {
		surname = surnameVal.(string)
	}

	language := ""
	if languageVal, ok := session.Values["language"]; ok {
		language = languageVal.(string)
	}

	roleID := ""
	if roleVal, ok := session.Values["role_id"]; ok {
		roleID = roleVal.(string)
	}

	role := ""
	if roleVal, ok := session.Values["role"]; ok {
		role = roleVal.(string)
	}

	return &connected_roots.Session{
		ID:       session.ID,
		Email:    email,
		UserID:   userID,
		Name:     name,
		Surname:  surname,
		Language: language,
		RoleID:   roleID,
		Role:     role,
		Cookie: &connected_roots.Cookie{
			Path:     session.Options.Path,
			Domain:   session.Options.Domain,
			MaxAge:   session.Options.MaxAge,
			Secure:   session.Options.Secure,
			HTTPOnly: session.Options.HttpOnly,
			SameSite: session.Options.SameSite,
		},
	}
}
