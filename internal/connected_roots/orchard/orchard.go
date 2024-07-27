package orchard

import (
	"context"
	"fmt"
	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql/orchard"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/Kortivex/connected_roots/pkg/pagination"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

const (
	tracingOrchard               = "service.orchard"
	tracingOrchardSave           = "service.orchard.save"
	tracingOrchardUpdate         = "service.orchard.update"
	tracingOrchardObtain         = "service.orchard.obtain"
	tracingOrchardObtainAll      = "service.orchard.obtain-all"
	tracingOrchardRemove         = "service.orchard.remove"
	tracingOrchardCountAll       = "service.orchard.count-all"
	tracingOrchardCountAllByUser = "service.orchard.count-all-by-user"
)

type Service struct {
	conf   *config.Config
	logger *logger.Logger
	// Repositories
	orchardRep *orchard.Repository
}

func New(conf *config.Config, db *gorm.DB, logr *logger.Logger) *Service {
	loggerEmpty := logr.NewEmpty()
	logr = loggerEmpty.WithTag(tracingOrchard)

	return &Service{
		conf:       conf,
		logger:     logr,
		orchardRep: orchard.NewRepository(conf, db, logr),
	}
}

func (s *Service) Save(ctx context.Context, orchard *connected_roots.Orchards) (*connected_roots.Orchards, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingOrchardSave)
	defer span.End()

	orchardsRes, err := s.orchardRep.Create(ctx, orchard)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingOrchardSave, err)
	}

	orchardsRes, err = s.orchardRep.GetByID(ctx, orchardsRes.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingOrchardSave, err)
	}

	return orchardsRes, nil
}

func (s *Service) Update(ctx context.Context, orchard *connected_roots.Orchards) (*connected_roots.Orchards, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingOrchardUpdate)
	defer span.End()

	orchardsRes, err := s.orchardRep.Update(ctx, orchard)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingOrchardUpdate, err)
	}

	orchardsRes, err = s.orchardRep.GetByID(ctx, orchardsRes.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingOrchardUpdate, err)
	}

	return orchardsRes, nil
}

func (s *Service) Obtain(ctx context.Context, id string) (*connected_roots.Orchards, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingOrchardObtain)
	defer span.End()

	orchardRes, err := s.orchardRep.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingOrchardObtain, err)
	}

	return orchardRes, nil
}

func (s *Service) ObtainAll(ctx context.Context, filters *connected_roots.OrchardPaginationFilters) (*pagination.Pagination, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingOrchardObtainAll)
	defer span.End()

	orchardRes, err := s.orchardRep.ListAllBy(ctx, filters, []string{"User", "CropType"}...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingOrchardObtainAll, err)
	}

	return orchardRes, nil
}

func (s *Service) Remove(ctx context.Context, id string) error {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingOrchardRemove)
	defer span.End()

	if err := s.orchardRep.DeleteByID(ctx, id); err != nil {
		return fmt.Errorf("%s: %w", tracingOrchardRemove, err)
	}

	return nil
}

func (s *Service) CountAll(ctx context.Context) (*connected_roots.TotalOrchards, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingOrchardCountAll)
	defer span.End()

	loggerNew := s.logger.New()
	log := loggerNew.WithTag(tracingOrchardCountAll)

	total, err := s.orchardRep.Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingOrchardCountAll, err)
	}

	log.Debug(fmt.Sprintf("total: %+v", total))

	return &connected_roots.TotalOrchards{Total: total}, nil
}

func (s *Service) CountAllByUser(ctx context.Context, userID string) (*connected_roots.TotalOrchards, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingOrchardCountAllByUser)
	defer span.End()

	loggerNew := s.logger.New()
	log := loggerNew.WithTag(tracingOrchardCountAllByUser)

	total, err := s.orchardRep.CountAllByUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingOrchardCountAllByUser, err)
	}

	log.Debug(fmt.Sprintf("total: %+v", total))

	return &connected_roots.TotalOrchards{Total: total}, nil
}
