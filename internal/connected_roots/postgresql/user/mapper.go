package user

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql/role"
)

func toDomain(user *Users) *connected_roots.Users {
	return &connected_roots.Users{
		ID:        user.ID,
		Name:      user.Name,
		Surname:   user.Surname,
		Email:     user.Email,
		Password:  user.Password,
		Telephone: user.Telephone,
		Language:  user.Language,
		Role:      toRoleDomain(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func toRoleDomain(rl *role.Roles) *connected_roots.Roles {
	return &connected_roots.Roles{
		ID:          rl.ID,
		Name:        rl.Name,
		Description: rl.Description,
		CreatedAt:   rl.CreatedAt,
		UpdatedAt:   rl.UpdatedAt,
	}
}
