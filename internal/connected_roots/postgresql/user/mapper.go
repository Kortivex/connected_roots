package user

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql"
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
		RoleID:    user.RoleID,
		Role:      toRoleDomain(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func toRoleDomain(rl *role.Roles) *connected_roots.Roles {
	if rl == nil {
		return nil
	}
	return &connected_roots.Roles{
		ID:          rl.ID,
		Name:        rl.Name,
		Description: rl.Description,
		CreatedAt:   rl.CreatedAt,
		UpdatedAt:   rl.UpdatedAt,
	}
}

func toDomainSlice(users []*Users) []*connected_roots.Users {
	usersDomain := []*connected_roots.Users{}
	for _, user := range users {
		roleDomain := toDomain(user)
		usersDomain = append(usersDomain, roleDomain)
	}
	return usersDomain
}

func toDB(user *connected_roots.Users, id string) *Users {
	return &Users{
		ID:        id,
		Name:      user.Name,
		Surname:   user.Surname,
		Email:     user.Email,
		Password:  user.Password,
		Telephone: user.Telephone,
		Language:  user.Language,
		RoleID:    user.RoleID,
		BaseModel: postgresql.BaseModel{
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}
}
