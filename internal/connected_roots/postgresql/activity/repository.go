package activity

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
	tracingActivity               = "repository-db.activity"
	tracingActivityCreate         = "repository-db.activity.create"
	tracingActivityUpdate         = "repository-db.activity.update"
	tracingActivityGetByID        = "repository-db.activity.get-by-id"
	tracingActivityListAllBy      = "repository-db.activity.list-all-by"
	tracingActivityDeleteByID     = "repository-db.activity.delete-by-id"
	tracingActivityCount          = "repository-db.activity.count"
	tracingActivityCountAllByUser = "repository-db.activity.count-all-by-user"
)

type Repository struct {
	conf   *config.Config
	db     *gorm.DB
	logger *logger.Logger
}

func NewRepository(conf *config.Config, db *gorm.DB, logr *logger.Logger) *Repository {
	loggerEmpty := logr.NewEmpty()
	log := loggerEmpty.WithTag(tracingActivity)

	return &Repository{
		conf:   conf,
		db:     db,
		logger: log,
	}
}

func (r *Repository) Create(ctx context.Context, activity *connected_roots.Activities) (*connected_roots.Activities, error) {
	ctx, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingActivityCreate)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingActivityCreate)

	log.Debug(fmt.Sprintf("activity: %+v", activity))

	id, err := ulid.Generate()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingActivityCreate, err)
	}

	activityDB := toDB(activity, id)
	result := r.db.WithContext(ctx).Model(&Activities{}).
		Create(&activityDB)

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingActivityCreate, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return nil, fmt.Errorf("%s: %w", tracingActivityCreate, gorm.ErrRecordNotFound)
	}

	log.Debug(fmt.Sprintf("activity: %+v", activityDB))

	return toDomain(activityDB), nil
}

func (r *Repository) Update(ctx context.Context, activity *connected_roots.Activities) (*connected_roots.Activities, error) {
	ctx, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingActivityUpdate)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingActivityUpdate)

	log.Debug(fmt.Sprintf("activity: %+v", activity))

	activityDB := toDB(activity, activity.ID)
	result := r.db.WithContext(ctx).Model(&Activities{}).
		Omit("id", "created_at", "deleted_at").
		Where("id = ?", activity.ID).
		Updates(&activityDB)

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingActivityUpdate, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return nil, fmt.Errorf("%s: %w", tracingActivityUpdate, gorm.ErrRecordNotFound)
	}

	log.Debug(fmt.Sprintf("activity: %+v", activityDB))

	return toDomain(activityDB), nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (*connected_roots.Activities, error) {
	ctx, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingActivityGetByID)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingActivityGetByID)

	log.Debug(fmt.Sprintf("orchard id: %+v", id))

	activityDB := &Activities{}
	result := r.db.WithContext(ctx).Model(&Activities{}).
		Preload("Orchard").
		Preload("Orchard.User").
		Preload("Orchard.CropType").
		Where("id = ?", id).
		First(&activityDB)

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingActivityGetByID, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return nil, fmt.Errorf("%s: %w", tracingActivityGetByID, gorm.ErrRecordNotFound)
	}

	log.Debug(fmt.Sprintf("activity: %+v", activityDB))

	return toDomain(activityDB), nil
}

func (r *Repository) ListAllBy(ctx context.Context, activityFilters *connected_roots.ActivityPaginationFilters, preloads ...string) (*pagination.Pagination, error) {
	ctx, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingActivityListAllBy)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingActivityListAllBy)

	log.Debug(fmt.Sprintf("filters: %+v", activityFilters))

	rulesBuilder := pagination.SortRulesBuilder{
		Sorts:               activityFilters.Sort,
		ValidateFields:      TableActivitiesFields,
		DBStructAssociation: TableActivitiesSortMap,
		TableName:           TableActivitiesName,
	}

	rules, err := rulesBuilder.ObtainRules()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingActivityListAllBy, err)
	}

	if len(rules) == 0 {
		rules = append(rules, DefaultActivityRule)
	}

	pg, err := pagination.CreatePaginator(&activityFilters.PaginatorParams, rules)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingActivityListAllBy, err)
	}

	var activitiesDB []*Activities
	query := r.db.WithContext(ctx).Model(&Activities{})
	for _, p := range preloads {
		query = query.Preload(p)
	}
	AddActivityFilters(query, &activityFilters.ActivityFilters)
	query.Find(&activitiesDB)

	result, cursor, err := pg.Paginate(query, &activitiesDB)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingActivityListAllBy, err)
	}

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingActivityListAllBy, result.Error)
	}

	previousCursor, nextCursor, err := pagination.EncodeURLValues(cursor)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingActivityListAllBy, err)
	}

	activitiesPaginated := &pagination.Pagination{
		Data: toDomainSlice(activitiesDB),
		Paging: pagination.Paging{
			NextCursor:     nextCursor,
			PreviousCursor: previousCursor,
		},
	}

	log.Debug(fmt.Sprintf("activities: %+v", activitiesDB))

	return activitiesPaginated, nil
}

func (r *Repository) DeleteByID(ctx context.Context, id string) error {
	ctx, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingActivityDeleteByID)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingActivityDeleteByID)

	log.Debug(fmt.Sprintf("activity id: %+v", id))

	activityDB := &Activities{ID: id}
	result := r.db.WithContext(ctx).Model(&Activities{}).
		Unscoped().
		Where("id = ?", id).
		Delete(&activityDB)

	if result.Error != nil {
		return fmt.Errorf("%s: %w", tracingActivityDeleteByID, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return fmt.Errorf("%s: %w", tracingActivityDeleteByID, gorm.ErrRecordNotFound)
	}

	return nil
}

func (r *Repository) Count(ctx context.Context) (int64, error) {
	ctx, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingActivityCount)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingActivityCount)

	var total int64
	result := r.db.WithContext(ctx).Model(&Activities{}).
		Count(&total)

	if result.Error != nil {
		return 0, fmt.Errorf("%s: %w", tracingActivityCount, result.Error)
	}

	log.Debug(fmt.Sprintf("total: %+v", total))

	return total, nil
}

func (r *Repository) CountAllByUser(ctx context.Context, userID string) (int64, error) {
	ctx, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingActivityCountAllByUser)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingActivityCountAllByUser)

	var total int64
	result := r.db.WithContext(ctx).Model(&Activities{}).
		Joins("JOIN orchards ON orchards.id = agricultural_activities.orchard_id").
		Where("orchards.user_id = ?", userID).
		Count(&total)

	if result.Error != nil {
		return 0, fmt.Errorf("%s: %w", tracingActivityCountAllByUser, result.Error)
	}

	log.Debug(fmt.Sprintf("total: %+v", total))

	return total, nil
}
