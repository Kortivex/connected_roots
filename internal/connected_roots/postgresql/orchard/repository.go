package orchard

import (
	"context"
	"fmt"
	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/Kortivex/connected_roots/pkg/pagination"
	"github.com/Kortivex/connected_roots/pkg/ulid"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

const (
	tracingOrchard           = "repository-db.orchard"
	tracingOrchardCreate     = "repository-db.orchard.create"
	tracingOrchardUpdate     = "repository-db.orchard.update"
	tracingOrchardGetByID    = "repository-db.orchard.get-by-id"
	tracingOrchardListAllBy  = "repository-db.orchard.list-all-by"
	tracingOrchardDeleteByID = "repository-db.orchard.delete-by-id"
)

type Repository struct {
	conf   *config.Config
	db     *gorm.DB
	logger *logger.Logger
}

func NewRepository(conf *config.Config, db *gorm.DB, logr *logger.Logger) *Repository {
	loggerEmpty := logr.NewEmpty()
	log := loggerEmpty.WithTag(tracingOrchard)

	return &Repository{
		conf:   conf,
		db:     db,
		logger: log,
	}
}

func (r *Repository) Create(ctx context.Context, orchard *connected_roots.Orchards) (*connected_roots.Orchards, error) {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingOrchardCreate)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingOrchardCreate)

	log.Debug(fmt.Sprintf("orchard: %+v", orchard))

	id, err := ulid.Generate()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingOrchardCreate, err)
	}

	orchardDB := toDB(orchard, id)
	result := r.db.WithContext(ctx).Model(&Orchards{}).
		Create(&orchardDB)

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingOrchardCreate, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return nil, fmt.Errorf("%s: %w", tracingOrchardCreate, gorm.ErrRecordNotFound)
	}

	log.Debug(fmt.Sprintf("orchard: %+v", orchardDB))

	return toDomain(orchardDB), nil
}

func (r *Repository) Update(ctx context.Context, orchard *connected_roots.Orchards) (*connected_roots.Orchards, error) {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingOrchardUpdate)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingOrchardUpdate)

	log.Debug(fmt.Sprintf("orchard: %+v", orchard))

	orchardDB := toDB(orchard, orchard.ID)
	result := r.db.WithContext(ctx).Model(&Orchards{}).
		Omit("id", "created_at", "deleted_at").
		Where("id = ?", orchard.ID).
		Updates(&orchardDB)

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingOrchardUpdate, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return nil, fmt.Errorf("%s: %w", tracingOrchardUpdate, gorm.ErrRecordNotFound)
	}

	log.Debug(fmt.Sprintf("orchard: %+v", orchardDB))

	return toDomain(orchardDB), nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (*connected_roots.Orchards, error) {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingOrchardGetByID)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingOrchardGetByID)

	log.Debug(fmt.Sprintf("orchard id: %+v", id))

	orchardDB := &Orchards{}
	result := r.db.WithContext(ctx).Model(&Orchards{}).
		Where("id = ?", id).
		First(&orchardDB)

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingOrchardGetByID, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return nil, fmt.Errorf("%s: %w", tracingOrchardGetByID, gorm.ErrRecordNotFound)
	}

	log.Debug(fmt.Sprintf("role: %+v", orchardDB))

	return toDomain(orchardDB), nil
}

func (r *Repository) ListAllBy(ctx context.Context, orchardFilters *connected_roots.OrchardPaginationFilters, preloads ...string) (*pagination.Pagination, error) {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingOrchardListAllBy)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingOrchardListAllBy)

	log.Debug(fmt.Sprintf("filters: %+v", orchardFilters))

	rulesBuilder := pagination.SortRulesBuilder{
		Sorts:               orchardFilters.Sort,
		ValidateFields:      TableOrchardsFields,
		DBStructAssociation: TableOrchardsSortMap,
		TableName:           TableOrchardsName,
	}

	rules, err := rulesBuilder.ObtainRules()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingOrchardListAllBy, err)
	}

	if len(rules) == 0 {
		rules = append(rules, DefaultOrchardRule)
	}

	pg, err := pagination.CreatePaginator(&orchardFilters.PaginatorParams, rules)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingOrchardListAllBy, err)
	}

	var orchardsDB []*Orchards
	query := r.db.WithContext(ctx).Model(&Orchards{})
	for _, p := range preloads {
		query = query.Preload(p)
	}
	AddOrchardFilters(query, &orchardFilters.OrchardFilters)
	query.Find(&orchardsDB)

	result, cursor, err := pg.Paginate(query, &orchardsDB)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingOrchardListAllBy, err)
	}

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingOrchardListAllBy, result.Error)
	}

	previousCursor, nextCursor, err := pagination.EncodeURLValues(cursor)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingOrchardListAllBy, err)
	}

	orchardsPaginated := &pagination.Pagination{
		Data: toDomainSlice(orchardsDB),
		Paging: pagination.Paging{
			NextCursor:     nextCursor,
			PreviousCursor: previousCursor,
		},
	}

	log.Debug(fmt.Sprintf("orchards: %+v", orchardsDB))

	return orchardsPaginated, nil
}

func (r *Repository) DeleteByID(ctx context.Context, id string) error {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingOrchardDeleteByID)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingOrchardDeleteByID)

	log.Debug(fmt.Sprintf("orchard id: %+v", id))

	orchardDB := &Orchards{ID: id}
	result := r.db.WithContext(ctx).Model(&Orchards{}).
		Unscoped().
		Where("id = ?", id).
		Delete(&orchardDB)

	if result.Error != nil {
		return fmt.Errorf("%s: %w", tracingOrchardDeleteByID, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return fmt.Errorf("%s: %w", tracingOrchardDeleteByID, gorm.ErrRecordNotFound)
	}

	return nil
}
