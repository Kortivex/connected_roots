package role

import (
	"context"
	"fmt"
	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql/role"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/Kortivex/connected_roots/pkg/pagination"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

const (
	tracingRole          = "service.role"
	tracingRoleSave      = "service.role.save"
	tracingRoleUpdate    = "service.role.update"
	tracingRoleObtain    = "service.role.obtain"
	tracingRoleObtainAll = "service.role.obtain-all"
)

type Service struct {
	conf   *config.Config
	logger *logger.Logger
	// Repositories
	roleRep *role.Repository
}

func New(conf *config.Config, db *gorm.DB, logr *logger.Logger) *Service {
	loggerEmpty := logr.NewEmpty()
	logr = loggerEmpty.WithTag(tracingRole)

	return &Service{
		conf:    conf,
		logger:  logr,
		roleRep: role.NewRepository(conf, db, logr),
	}
}

func (s *Service) Save(ctx context.Context, role *connected_roots.Roles) (*connected_roots.Roles, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingRoleSave)
	defer span.End()

	rolesRes, err := s.roleRep.Create(ctx, role)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingRoleSave, err)
	}

	rolesRes, err = s.roleRep.GetByID(ctx, rolesRes.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingRoleSave, err)
	}

	return rolesRes, nil
}

func (s *Service) Update(ctx context.Context, role *connected_roots.Roles) (*connected_roots.Roles, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingRoleUpdate)
	defer span.End()

	rolesRes, err := s.roleRep.Update(ctx, role)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingRoleUpdate, err)
	}

	rolesRes, err = s.roleRep.GetByID(ctx, rolesRes.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingRoleUpdate, err)
	}

	return rolesRes, nil
}

func (s *Service) Obtain(ctx context.Context, id string) (*connected_roots.Roles, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingRoleObtain)
	defer span.End()

	roleRes, err := s.roleRep.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingRoleObtain, err)
	}

	return roleRes, nil
}

func (s *Service) ObtainAll(ctx context.Context, filters *connected_roots.RolePaginationFilters) (*pagination.Pagination, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingRoleObtainAll)
	defer span.End()

	rolesRes, err := s.roleRep.ListAllBy(ctx, filters, []string{}...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingRoleObtainAll, err)
	}

	return rolesRes, nil
}
