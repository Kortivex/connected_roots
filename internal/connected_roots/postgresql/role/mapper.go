package role

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots"
)

func toDomain(role *Roles) *connected_roots.Roles {
	return &connected_roots.Roles{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
		Protected:   role.Protected,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}
}

func toDomainSlice(roles []*Roles) []*connected_roots.Roles {
	rolesDomain := []*connected_roots.Roles{}
	for _, role := range roles {
		roleDomain := toDomain(role)
		rolesDomain = append(rolesDomain, roleDomain)
	}
	return rolesDomain
}

func toDB(role *connected_roots.Roles, id string) *Roles {
	return &Roles{
		ID:          id,
		Name:        role.Name,
		Description: role.Description,
		Protected:   role.Protected,
	}
}
