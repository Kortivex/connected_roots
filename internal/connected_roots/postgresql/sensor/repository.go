package sensor

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
	tracingSensor           = "repository-db.sensor"
	tracingSensorCreate     = "repository-db.sensor.create"
	tracingSensorUpdate     = "repository-db.sensor.update"
	tracingSensorGetByID    = "repository-db.sensor.get-by-id"
	tracingSensorListAllBy  = "repository-db.sensor.list-all-by"
	tracingSensorDeleteByID = "repository-db.sensor.delete-by-id"
)

type Repository struct {
	conf   *config.Config
	db     *gorm.DB
	logger *logger.Logger
}

func NewRepository(conf *config.Config, db *gorm.DB, logr *logger.Logger) *Repository {
	loggerEmpty := logr.NewEmpty()
	log := loggerEmpty.WithTag(tracingSensor)

	return &Repository{
		conf:   conf,
		db:     db,
		logger: log,
	}
}

func (r *Repository) Create(ctx context.Context, sensor *connected_roots.Sensors) (*connected_roots.Sensors, error) {
	ctx, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingSensorCreate)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingSensorCreate)

	log.Debug(fmt.Sprintf("sensor: %+v", sensor))

	id, err := ulid.Generate()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingSensorCreate, err)
	}

	sensorDB := toDB(sensor, id)
	result := r.db.WithContext(ctx).Model(&Sensors{}).
		Create(&sensorDB)

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingSensorCreate, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return nil, fmt.Errorf("%s: %w", tracingSensorCreate, gorm.ErrRecordNotFound)
	}

	log.Debug(fmt.Sprintf("sensor: %+v", sensorDB))

	return toDomain(sensorDB), nil
}

func (r *Repository) Update(ctx context.Context, sensor *connected_roots.Sensors) (*connected_roots.Sensors, error) {
	ctx, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingSensorUpdate)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingSensorUpdate)

	log.Debug(fmt.Sprintf("sensor: %+v", sensor))

	sensorDB := toDB(sensor, sensor.ID)
	result := r.db.WithContext(ctx).Model(&Sensors{}).
		Omit("id", "created_at", "deleted_at").
		Where("id = ?", sensor.ID).
		Updates(&sensorDB)

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingSensorUpdate, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return nil, fmt.Errorf("%s: %w", tracingSensorUpdate, gorm.ErrRecordNotFound)
	}

	log.Debug(fmt.Sprintf("sensor: %+v", sensorDB))

	return toDomain(sensorDB), nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (*connected_roots.Sensors, error) {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingSensorGetByID)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingSensorGetByID)

	log.Debug(fmt.Sprintf("sensor id: %+v", id))

	sensorDB := &Sensors{}
	result := r.db.WithContext(ctx).Model(&Sensors{}).
		Preload("Orchard").
		Preload("Orchard.User").
		Preload("Orchard.CropType").
		Where("id = ?", id).
		First(&sensorDB)

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingSensorGetByID, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return nil, fmt.Errorf("%s: %w", tracingSensorGetByID, gorm.ErrRecordNotFound)
	}

	log.Debug(fmt.Sprintf("sensor: %+v", sensorDB))

	return toDomain(sensorDB), nil
}

func (r *Repository) ListAllBy(ctx context.Context, sensorFilters *connected_roots.SensorPaginationFilters, preloads ...string) (*pagination.Pagination, error) {
	ctx, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingSensorListAllBy)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingSensorListAllBy)

	log.Debug(fmt.Sprintf("filters: %+v", sensorFilters))

	rulesBuilder := pagination.SortRulesBuilder{
		Sorts:               sensorFilters.Sort,
		ValidateFields:      TableSensorsFields,
		DBStructAssociation: TableSensorsSortMap,
		TableName:           TableSensorsName,
	}

	rules, err := rulesBuilder.ObtainRules()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingSensorListAllBy, err)
	}

	if len(rules) == 0 {
		rules = append(rules, DefaultSensorRule)
	}

	pg, err := pagination.CreatePaginator(&sensorFilters.PaginatorParams, rules)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingSensorListAllBy, err)
	}

	var sensorsDB []*Sensors
	query := r.db.WithContext(ctx).Model(&Sensors{})
	for _, p := range preloads {
		query = query.Preload(p)
	}
	AddOrchardFilters(query, &sensorFilters.SensorFilters)
	query.Find(&sensorsDB)

	result, cursor, err := pg.Paginate(query, &sensorsDB)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingSensorListAllBy, err)
	}

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingSensorListAllBy, result.Error)
	}

	previousCursor, nextCursor, err := pagination.EncodeURLValues(cursor)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingSensorListAllBy, err)
	}

	sensorsPaginated := &pagination.Pagination{
		Data: toDomainSlice(sensorsDB),
		Paging: pagination.Paging{
			NextCursor:     nextCursor,
			PreviousCursor: previousCursor,
		},
	}

	log.Debug(fmt.Sprintf("sensors: %+v", sensorsDB))

	return sensorsPaginated, nil
}

func (r *Repository) DeleteByID(ctx context.Context, id string) error {
	ctx, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingSensorDeleteByID)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingSensorDeleteByID)

	log.Debug(fmt.Sprintf("sensor id: %+v", id))

	sensorDB := &Sensors{ID: id}
	result := r.db.WithContext(ctx).Model(&Sensors{}).
		Unscoped().
		Where("id = ?", id).
		Delete(&sensorDB)

	if result.Error != nil {
		return fmt.Errorf("%s: %w", tracingSensorDeleteByID, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return fmt.Errorf("%s: %w", tracingSensorDeleteByID, gorm.ErrRecordNotFound)
	}

	return nil
}
