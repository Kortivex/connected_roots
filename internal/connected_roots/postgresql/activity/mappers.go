package activity

import (
	"time"

	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql/crop_type"
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql/orchard"
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql/role"
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql/user"
)

func toDomain(activity *Activities) *connected_roots.Activities {
	return &connected_roots.Activities{
		ID:          activity.ID,
		Name:        activity.Name,
		Description: activity.Description,
		OrchardID:   activity.OrchardID,
		Orchard:     toOrchardDomain(activity.Orchard),
		CreatedAt:   activity.CreatedAt,
		UpdatedAt:   activity.UpdatedAt,
	}
}

func toOrchardDomain(orchard *orchard.Orchards) *connected_roots.Orchards {
	if orchard == nil {
		return nil
	}
	return &connected_roots.Orchards{
		ID:         orchard.ID,
		Name:       orchard.Name,
		Location:   orchard.Location,
		Size:       orchard.Size,
		Soil:       orchard.Soil,
		Fertilizer: orchard.Fertilizer,
		Composting: orchard.Composting,
		ImageURL:   orchard.ImageURL,
		UserID:     orchard.UserID,
		User:       toUserDomain(orchard.User),
		CropTypeID: orchard.CropTypeID,
		CropType:   toCropTypeDomain(orchard.CropType),
		CreatedAt:  orchard.CreatedAt,
		UpdatedAt:  orchard.UpdatedAt,
	}
}

func toUserDomain(usr *user.Users) *connected_roots.Users {
	if usr == nil {
		return nil
	}
	return &connected_roots.Users{
		ID:        usr.ID,
		Name:      usr.Name,
		Surname:   usr.Surname,
		Email:     usr.Email,
		Password:  usr.Password,
		Telephone: usr.Telephone,
		Language:  usr.Language,
		RoleID:    usr.RoleID,
		Role:      toRoleDomain(usr.Role),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
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

func toCropTypeDomain(ctp *crop_type.CropTypes) *connected_roots.CropTypes {
	if ctp == nil {
		return nil
	}
	return &connected_roots.CropTypes{
		ID:             ctp.ID,
		Name:           ctp.Name,
		ScientificName: ctp.ScientificName,
		LifeCycle:      ctp.LifeCycle,
		PlantingSeason: ctp.PlantingSeason,
		HarvestSeason:  ctp.HarvestSeason,
		Irrigation:     ctp.Irrigation,
		Description:    ctp.Description,
		CreatedAt:      ctp.CreatedAt,
		UpdatedAt:      ctp.UpdatedAt,
	}
}

func toDomainSlice(activities []*Activities) []*connected_roots.Activities {
	activitiesDomain := []*connected_roots.Activities{}
	for _, activity := range activities {
		activityDomain := toDomain(activity)
		activitiesDomain = append(activitiesDomain, activityDomain)
	}
	return activitiesDomain
}

func toDB(activity *connected_roots.Activities, id string) *Activities {
	return &Activities{
		ID:          id,
		Name:        activity.Name,
		Description: activity.Description,
		OrchardID:   activity.OrchardID,
	}
}
