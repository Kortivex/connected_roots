package crop_type

import (
	"context"
	"fmt"
	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/Kortivex/connected_roots/pkg/pagination"
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
