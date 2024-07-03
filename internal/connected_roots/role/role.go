package role

import (
	"context"
	"fmt"
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql/role"
	"github.com/Kortivex/connected_roots/pkg/pagination"

	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

const (
	tracingRole          = "service.role"
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

func (s *Service) ObtainAll(ctx context.Context, filters *connected_roots.RolePaginationFilters) (*pagination.Pagination, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingRoleObtainAll)
	defer span.End()

	rolesRes, err := s.roleRep.ListAllBy(ctx, filters, []string{}...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingRoleObtainAll, err)
	}

	return rolesRes, nil
}
