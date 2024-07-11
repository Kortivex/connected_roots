package crop_type

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
	tracingCropTypes           = "repository-db.crop-types"
	tracingCropTypesCreate     = "repository-db.crop-types.create"
	tracingCropTypesUpdate     = "repository-db.crop-types.update"
	tracingCropTypesGetByID    = "repository-db.crop-types.get-by-id"
	tracingCropTypesListAllBy  = "repository-db.crop-types.list-all-by"
	tracingCropTypesDeleteByID = "repository-db.crop-types.delete-by-id"
)

type Repository struct {
	conf   *config.Config
	db     *gorm.DB
	logger *logger.Logger
}

func NewRepository(conf *config.Config, db *gorm.DB, logr *logger.Logger) *Repository {
	loggerEmpty := logr.NewEmpty()
	log := loggerEmpty.WithTag(tracingCropTypes)

	return &Repository{
		conf:   conf,
		db:     db,
		logger: log,
	}
}

func (r *Repository) Create(ctx context.Context, cropType *connected_roots.CropTypes) (*connected_roots.CropTypes, error) {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingCropTypesCreate)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingCropTypesCreate)

	log.Debug(fmt.Sprintf("crop-type: %+v", cropType))

	id, err := ulid.Generate()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingCropTypesCreate, err)
	}

	cropTypeDB := toDB(cropType, id)
	result := r.db.WithContext(ctx).Model(&CropTypes{}).
		Create(&cropTypeDB)

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingCropTypesCreate, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return nil, fmt.Errorf("%s: %w", tracingCropTypesCreate, gorm.ErrRecordNotFound)
	}

	log.Debug(fmt.Sprintf("crop-type: %+v", cropTypeDB))

	return toDomain(cropTypeDB), nil
}

func (r *Repository) Update(ctx context.Context, cropType *connected_roots.CropTypes) (*connected_roots.CropTypes, error) {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingCropTypesUpdate)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingCropTypesUpdate)

	log.Debug(fmt.Sprintf("crop-type: %+v", cropType))

	cropTypeDB := toDB(cropType, cropType.ID)
	result := r.db.WithContext(ctx).Model(&CropTypes{}).
		Omit("id", "created_at", "deleted_at").
		Where("id = ?", cropType.ID).
		Updates(&cropTypeDB)

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingCropTypesUpdate, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return nil, fmt.Errorf("%s: %w", tracingCropTypesUpdate, gorm.ErrRecordNotFound)
	}

	log.Debug(fmt.Sprintf("crop-type: %+v", cropTypeDB))

	return toDomain(cropTypeDB), nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (*connected_roots.CropTypes, error) {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingCropTypesGetByID)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingCropTypesGetByID)

	log.Debug(fmt.Sprintf("crop-type id: %+v", id))

	cropTypeDB := &CropTypes{}
	result := r.db.WithContext(ctx).Model(&CropTypes{}).
		Where("id = ?", id).
		First(&cropTypeDB)

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingCropTypesGetByID, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return nil, fmt.Errorf("%s: %w", tracingCropTypesGetByID, gorm.ErrRecordNotFound)
	}

	log.Debug(fmt.Sprintf("crop-type: %+v", cropTypeDB))

	return toDomain(cropTypeDB), nil
}

func (r *Repository) ListAllBy(ctx context.Context, cropTypeFilters *connected_roots.CropTypePaginationFilters, preloads ...string) (*pagination.Pagination, error) {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingCropTypesListAllBy)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingCropTypesListAllBy)

	log.Debug(fmt.Sprintf("filters: %+v", cropTypeFilters))

	rulesBuilder := pagination.SortRulesBuilder{
		Sorts:               cropTypeFilters.Sort,
		ValidateFields:      TableCropTypesFields,
		DBStructAssociation: TableCropTypesSortMap,
		TableName:           TableCropTypesName,
	}

	rules, err := rulesBuilder.ObtainRules()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingCropTypesListAllBy, err)
	}

	if len(rules) == 0 {
		rules = append(rules, DefaultCropTypeRule)
	}

	pg, err := pagination.CreatePaginator(&cropTypeFilters.PaginatorParams, rules)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingCropTypesListAllBy, err)
	}

	var cropTypesDB []*CropTypes
	query := r.db.WithContext(ctx).Model(&CropTypes{})
	for _, p := range preloads {
		query = query.Preload(p)
	}
	AddCropTypeFilters(query, &cropTypeFilters.CropTypeFilters)
	query.Find(&cropTypesDB)

	result, cursor, err := pg.Paginate(query, &cropTypesDB)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingCropTypesListAllBy, err)
	}

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingCropTypesListAllBy, result.Error)
	}

	previousCursor, nextCursor, err := pagination.EncodeURLValues(cursor)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingCropTypesListAllBy, err)
	}

	cropTypesPaginated := &pagination.Pagination{
		Data: toDomainSlice(cropTypesDB),
		Paging: pagination.Paging{
			NextCursor:     nextCursor,
			PreviousCursor: previousCursor,
		},
	}

	log.Debug(fmt.Sprintf("cropTypes: %+v", cropTypesDB))

	return cropTypesPaginated, nil
}

func (r *Repository) DeleteByID(ctx context.Context, id string) error {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingCropTypesDeleteByID)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingCropTypesDeleteByID)

	log.Debug(fmt.Sprintf("crop-type id: %+v", id))

	cropTypeDB := &CropTypes{ID: id}
	result := r.db.WithContext(ctx).Model(&CropTypes{}).
		Unscoped().
		Where("id = ?", id).
		Delete(&cropTypeDB)

	if result.Error != nil {
		return fmt.Errorf("%s: %w", tracingCropTypesDeleteByID, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return fmt.Errorf("%s: %w", tracingCropTypesDeleteByID, gorm.ErrRecordNotFound)
	}

	return nil
}
