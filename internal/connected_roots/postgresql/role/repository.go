package role

import (
	"context"
	"fmt"
	"github.com/Kortivex/connected_roots/pkg/pagination"
	"github.com/Kortivex/connected_roots/pkg/ulid"

	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

const (
	tracingRole           = "repository-db.role"
	tracingRoleCreate     = "repository-db.role.create"
	tracingRoleUpdate     = "repository-db.role.update"
	tracingRoleGetByID    = "repository-db.role.get-by-id"
	tracingRoleListAllBy  = "repository-db.role.list-all-by"
	tracingRoleDeleteByID = "repository-db.role.delete-by-id"
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

func (r *Repository) Create(ctx context.Context, role *connected_roots.Roles) (*connected_roots.Roles, error) {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingRoleCreate)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingRoleCreate)

	log.Debug(fmt.Sprintf("role: %+v", role))

	id, err := ulid.Generate()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingRoleCreate, err)
	}

	roleDB := toDB(role, id)
	result := r.db.WithContext(ctx).Model(&Roles{}).
		Create(&roleDB)

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingRoleCreate, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return nil, fmt.Errorf("%s: %w", tracingRoleCreate, gorm.ErrRecordNotFound)
	}

	log.Debug(fmt.Sprintf("role: %+v", roleDB))

	return toDomain(roleDB), nil
}

func (r *Repository) Update(ctx context.Context, role *connected_roots.Roles) (*connected_roots.Roles, error) {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingRoleUpdate)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingRoleUpdate)

	log.Debug(fmt.Sprintf("role: %+v", role))

	roleDB := toDB(role, role.ID)
	result := r.db.WithContext(ctx).Model(&Roles{}).
		Omit("id", "created_at", "deleted_at").
		Where("id = ?", role.ID).
		Updates(&roleDB)

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingRoleUpdate, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return nil, fmt.Errorf("%s: %w", tracingRoleUpdate, gorm.ErrRecordNotFound)
	}

	log.Debug(fmt.Sprintf("role: %+v", roleDB))

	return toDomain(roleDB), nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (*connected_roots.Roles, error) {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingRoleGetByID)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingRoleGetByID)

	log.Debug(fmt.Sprintf("role id: %+v", id))

	roleDB := &Roles{}
	result := r.db.WithContext(ctx).Model(&Roles{}).
		Where("id = ?", id).
		First(&roleDB)

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingRoleGetByID, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return nil, fmt.Errorf("%s: %w", tracingRoleGetByID, gorm.ErrRecordNotFound)
	}

	log.Debug(fmt.Sprintf("role: %+v", roleDB))

	return toDomain(roleDB), nil
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

func (r *Repository) DeleteByID(ctx context.Context, id string) error {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingRoleDeleteByID)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingRoleDeleteByID)

	log.Debug(fmt.Sprintf("role id: %+v", id))

	roleDB := &Roles{ID: id}
	result := r.db.WithContext(ctx).Model(&Roles{}).
		Unscoped().
		Where("id = ?", id).
		Delete(&roleDB)

	if result.Error != nil {
		return fmt.Errorf("%s: %w", tracingRoleDeleteByID, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return fmt.Errorf("%s: %w", tracingRoleDeleteByID, gorm.ErrRecordNotFound)
	}

	return nil
}
