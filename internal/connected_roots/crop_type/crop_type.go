package crop_type

import (
	"context"
	"fmt"

	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql/crop_type"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/Kortivex/connected_roots/pkg/pagination"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

const (
	tracingCropType             = "service.crop-types"
	tracingCropTypeSave         = "service.crop-types.save"
	tracingCropTypeUpdate       = "service.crop-types.update"
	tracingCropTypeObtainFromID = "service.crop-types.obtain-from-id"
	tracingCropTypeObtainAll    = "service.crop-types.obtain-all"
	tracingCropTypeRemoveByID   = "service.crop-types.remove-by-id"
	tracingCropTypeCountAll     = "service.crop-types.count-all"
)

type Service struct {
	conf   *config.Config
	logger *logger.Logger
	// Repositories
	cropTypeRep *crop_type.Repository
}

func New(conf *config.Config, db *gorm.DB, logr *logger.Logger) *Service {
	loggerEmpty := logr.NewEmpty()
	logr = loggerEmpty.WithTag(tracingCropType)

	return &Service{
		conf:        conf,
		logger:      logr,
		cropTypeRep: crop_type.NewRepository(conf, db, logr),
	}
}

func (s *Service) Save(ctx context.Context, cropType *connected_roots.CropTypes) (*connected_roots.CropTypes, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingCropTypeSave)
	defer span.End()

	cropTypeRes, err := s.cropTypeRep.Create(ctx, cropType)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingCropTypeSave, err)
	}

	cropTypeRes, err = s.cropTypeRep.GetByID(ctx, cropTypeRes.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingCropTypeSave, err)
	}

	return cropTypeRes, nil
}

func (s *Service) Update(ctx context.Context, cropType *connected_roots.CropTypes) (*connected_roots.CropTypes, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingCropTypeUpdate)
	defer span.End()

	cropTypeRes, err := s.cropTypeRep.Update(ctx, cropType)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingCropTypeUpdate, err)
	}

	cropTypeRes, err = s.cropTypeRep.GetByID(ctx, cropType.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingCropTypeUpdate, err)
	}

	return cropTypeRes, nil
}

func (s *Service) Obtain(ctx context.Context, id string) (*connected_roots.CropTypes, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingCropTypeObtainFromID)
	defer span.End()

	cropTypeRes, err := s.cropTypeRep.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingCropTypeObtainFromID, err)
	}

	return cropTypeRes, nil
}

func (s *Service) ObtainAll(ctx context.Context, filters *connected_roots.CropTypePaginationFilters) (*pagination.Pagination, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingCropTypeObtainAll)
	defer span.End()

	cropTypesRes, err := s.cropTypeRep.ListAllBy(ctx, filters, []string{}...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingCropTypeObtainAll, err)
	}

	return cropTypesRes, nil
}

func (s *Service) Remove(ctx context.Context, id string) error {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingCropTypeRemoveByID)
	defer span.End()

	if err := s.cropTypeRep.DeleteByID(ctx, id); err != nil {
		return fmt.Errorf("%s: %w", tracingCropTypeRemoveByID, err)
	}

	return nil
}

func (s *Service) CountAll(ctx context.Context) (*connected_roots.TotalCropTypes, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingCropTypeCountAll)
	defer span.End()

	loggerNew := s.logger.New()
	log := loggerNew.WithTag(tracingCropTypeCountAll)

	total, err := s.cropTypeRep.Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingCropTypeCountAll, err)
	}

	log.Debug(fmt.Sprintf("total: %+v", total))

	return &connected_roots.TotalCropTypes{Total: total}, nil
}
