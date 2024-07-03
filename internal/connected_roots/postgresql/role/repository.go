package role

import (
	"context"
	"fmt"
	"github.com/Kortivex/connected_roots/pkg/pagination"

	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

const (
	tracingRole          = "repository-db.role"
	tracingRoleListAllBy = "repository-db.role.list-all-by"
)

type Repository struct {
	conf   *config.Config
	db     *gorm.DB
	logger *logger.Logger
}

func NewRepository(conf *config.Config, db *gorm.DB, logr *logger.Logger) *Repository {
	loggerEmpty := logr.NewEmpty()
	log := loggerEmpty.WithTag(tracingRole)

	return &Repository{
		conf:   conf,
		db:     db,
		logger: log,
	}
}

func (r *Repository) ListAllBy(ctx context.Context, rolesFilters *connected_roots.RolePaginationFilters, preloads ...string) (*pagination.Pagination, error) {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingRoleListAllBy)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingRoleListAllBy)

	log.Debug(fmt.Sprintf("filters: %+v", rolesFilters))

	rulesBuilder := pagination.SortRulesBuilder{
		Sorts:               rolesFilters.Sort,
		ValidateFields:      TableRolesFields,
		DBStructAssociation: TableRolesSortMap,
		TableName:           TableRolesName,
	}

	rules, err := rulesBuilder.ObtainRules()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingRoleListAllBy, err)
	}

	if len(rules) == 0 {
		rules = append(rules, DefaultRoleRule)
	}

	pg, err := pagination.CreatePaginator(&rolesFilters.PaginatorParams, rules)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingRoleListAllBy, err)
	}

	var rolesDB []*Roles
	query := r.db.WithContext(ctx).Model(&Roles{})
	for _, p := range preloads {
		query = query.Preload(p)
	}
	AddRoleFilters(query, &rolesFilters.RoleFilters)
	query.Find(&rolesDB)

	result, cursor, err := pg.Paginate(query, &rolesDB)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingRoleListAllBy, err)
	}

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingRoleListAllBy, result.Error)
	}

	previousCursor, nextCursor, err := pagination.EncodeURLValues(cursor)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingRoleListAllBy, err)
	}

	rolesPaginated := &pagination.Pagination{
		Data: toDomainSlice(rolesDB),
		Paging: pagination.Paging{
			NextCursor:     nextCursor,
			PreviousCursor: previousCursor,
		},
	}

	log.Debug(fmt.Sprintf("roles: %+v", rolesDB))

	return rolesPaginated, nil
}
