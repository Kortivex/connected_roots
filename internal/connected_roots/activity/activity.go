package activity

import (
	"context"
	"fmt"
	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql/activity"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/Kortivex/connected_roots/pkg/pagination"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

const (
	tracingActivity               = "service.activity"
	tracingActivitySave           = "service.activity.save"
	tracingActivityUpdate         = "service.activity.update"
	tracingActivityObtain         = "service.activity.obtain"
	tracingActivityObtainAll      = "service.activity.obtain-all"
	tracingActivityRemove         = "service.activity.remove"
	tracingActivityCountAll       = "service.activity.count-all"
	tracingActivityCountAllByUser = "service.activity.count-all-by-user"
)

type Service struct {
	conf   *config.Config
	logger *logger.Logger
	// Repositories
	activityRep *activity.Repository
}

func New(conf *config.Config, db *gorm.DB, logr *logger.Logger) *Service {
	loggerEmpty := logr.NewEmpty()
	logr = loggerEmpty.WithTag(tracingActivity)

	return &Service{
		conf:        conf,
		logger:      logr,
		activityRep: activity.NewRepository(conf, db, logr),
	}
}

func (s *Service) Save(ctx context.Context, activity *connected_roots.Activities) (*connected_roots.Activities, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingActivitySave)
	defer span.End()

	activitiesRes, err := s.activityRep.Create(ctx, activity)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingActivitySave, err)
	}

	activitiesRes, err = s.activityRep.GetByID(ctx, activitiesRes.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingActivitySave, err)
	}

	return activitiesRes, nil
}

func (s *Service) Update(ctx context.Context, activity *connected_roots.Activities) (*connected_roots.Activities, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingActivityUpdate)
	defer span.End()

	activitiesRes, err := s.activityRep.Update(ctx, activity)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingActivityUpdate, err)
	}

	activitiesRes, err = s.activityRep.GetByID(ctx, activitiesRes.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingActivityUpdate, err)
	}

	return activitiesRes, nil
}

func (s *Service) Obtain(ctx context.Context, id string) (*connected_roots.Activities, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingActivityObtain)
	defer span.End()

	activitiesRes, err := s.activityRep.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingActivityObtain, err)
	}

	return activitiesRes, nil
}

func (s *Service) ObtainAll(ctx context.Context, filters *connected_roots.ActivityPaginationFilters) (*pagination.Pagination, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingActivityObtainAll)
	defer span.End()

	activitiesRes, err := s.activityRep.ListAllBy(ctx, filters, []string{"Orchard", "Orchard.User", "Orchard.CropType"}...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingActivityObtainAll, err)
	}

	return activitiesRes, nil
}

func (s *Service) Remove(ctx context.Context, id string) error {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingActivityRemove)
	defer span.End()

	if err := s.activityRep.DeleteByID(ctx, id); err != nil {
		return fmt.Errorf("%s: %w", tracingActivityRemove, err)
	}

	return nil
}

func (s *Service) CountAll(ctx context.Context) (*connected_roots.TotalActivities, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingActivityCountAll)
	defer span.End()

	loggerNew := s.logger.New()
	log := loggerNew.WithTag(tracingActivityCountAll)

	total, err := s.activityRep.Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingActivityCountAll, err)
	}

	log.Debug(fmt.Sprintf("total: %+v", total))

	return &connected_roots.TotalActivities{Total: total}, nil
}

func (s *Service) CountAllByUser(ctx context.Context, userID string) (*connected_roots.TotalActivities, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingActivityCountAllByUser)
	defer span.End()

	loggerNew := s.logger.New()
	log := loggerNew.WithTag(tracingActivityCountAllByUser)

	total, err := s.activityRep.CountAllByUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingActivityCountAllByUser, err)
	}

	log.Debug(fmt.Sprintf("total: %+v", total))

	return &connected_roots.TotalActivities{Total: total}, nil
}
