package sensor

import (
	"context"
	"fmt"
	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql/sensor"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/Kortivex/connected_roots/pkg/pagination"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

const (
	tracingSensor                  = "service.sensor"
	tracingSensorSave              = "service.sensor.save"
	tracingSensorUpdate            = "service.sensor.update"
	tracingSensorObtain            = "service.sensor.obtain"
	tracingSensorObtainAll         = "service.sensor.obtain-all"
	tracingSensorRemove            = "service.sensor.remove"
	tracingSensorSaveData          = "service.sensor.save-data"
	tracingSensorObtainAllData     = "service.sensor.obtain-all-data"
	tracingSensorObtainAllByUserID = "service.sensor.obtain-all-by-user-id"
	tracingSensorCountAll          = "service.sensor.count-all"
	tracingSensorCountAllByUser    = "service.sensor.count-all-by-user"
)

type Service struct {
	conf   *config.Config
	logger *logger.Logger
	// Repositories
	sensorRep *sensor.Repository
}

func New(conf *config.Config, db *gorm.DB, logr *logger.Logger) *Service {
	loggerEmpty := logr.NewEmpty()
	logr = loggerEmpty.WithTag(tracingSensor)

	return &Service{
		conf:      conf,
		logger:    logr,
		sensorRep: sensor.NewRepository(conf, db, logr),
	}
}

func (s *Service) Save(ctx context.Context, sensor *connected_roots.Sensors) (*connected_roots.Sensors, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingSensorSave)
	defer span.End()

	sensorRes, err := s.sensorRep.Create(ctx, sensor)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingSensorSave, err)
	}

	sensorRes, err = s.sensorRep.GetByID(ctx, sensorRes.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingSensorSave, err)
	}

	return sensorRes, nil
}

func (s *Service) Update(ctx context.Context, sensor *connected_roots.Sensors) (*connected_roots.Sensors, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingSensorUpdate)
	defer span.End()

	sensorRes, err := s.sensorRep.Update(ctx, sensor)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingSensorUpdate, err)
	}

	sensorRes, err = s.sensorRep.GetByID(ctx, sensorRes.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingSensorUpdate, err)
	}

	return sensorRes, nil
}

func (s *Service) Obtain(ctx context.Context, id string) (*connected_roots.Sensors, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingSensorObtain)
	defer span.End()

	sensorRes, err := s.sensorRep.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingSensorObtain, err)
	}

	return sensorRes, nil
}

func (s *Service) ObtainAll(ctx context.Context, filters *connected_roots.SensorPaginationFilters) (*pagination.Pagination, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingSensorObtainAll)
	defer span.End()

	sensorsRes, err := s.sensorRep.ListAllBy(ctx, filters, []string{"Orchard", "Orchard.User", "Orchard.CropType"}...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingSensorObtainAll, err)
	}

	return sensorsRes, nil
}

func (s *Service) Remove(ctx context.Context, id string) error {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingSensorRemove)
	defer span.End()

	if err := s.sensorRep.DeleteByID(ctx, id); err != nil {
		return fmt.Errorf("%s: %w", tracingSensorRemove, err)
	}

	return nil
}

func (s *Service) SaveData(ctx context.Context, sensorData *connected_roots.SensorsData) (*connected_roots.SensorsData, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingSensorSaveData)
	defer span.End()

	sensorDataRes, err := s.sensorRep.CreateData(ctx, sensorData)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingSensorSaveData, err)
	}

	sensorDataRes, err = s.sensorRep.GetDataByID(ctx, sensorDataRes.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingSensorSaveData, err)
	}

	return sensorDataRes, nil
}

func (s *Service) ObtainAllData(ctx context.Context, filters *connected_roots.SensorDataPaginationFilters) (*pagination.Pagination, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingSensorObtainAllData)
	defer span.End()

	sensorsDataRes, err := s.sensorRep.ListAllDataBy(ctx, filters, []string{"Sensor"}...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingSensorObtainAllData, err)
	}

	return sensorsDataRes, nil
}

func (s *Service) ObtainAllByUserID(ctx context.Context, userID string, filters *connected_roots.SensorPaginationFilters) (*pagination.Pagination, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingSensorObtainAllByUserID)
	defer span.End()

	sensorsRes, err := s.sensorRep.ListAllByUserID(ctx, userID, filters, []string{"Orchard", "Orchard.User", "Orchard.CropType"}...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingSensorObtainAllByUserID, err)
	}

	return sensorsRes, nil
}

func (s *Service) CountAll(ctx context.Context) (*connected_roots.TotalSensors, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingSensorCountAll)
	defer span.End()

	loggerNew := s.logger.New()
	log := loggerNew.WithTag(tracingSensorCountAll)

	total, err := s.sensorRep.Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingSensorCountAll, err)
	}

	log.Debug(fmt.Sprintf("total: %+v", total))

	return &connected_roots.TotalSensors{Total: total}, nil
}

func (s *Service) CountAllByUserID(ctx context.Context, userID string) (*connected_roots.TotalSensors, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingSensorCountAllByUser)
	defer span.End()

	loggerNew := s.logger.New()
	log := loggerNew.WithTag(tracingSensorCountAllByUser)

	total, err := s.sensorRep.CountAllByUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingSensorCountAllByUser, err)
	}

	log.Debug(fmt.Sprintf("total: %+v", total))

	return &connected_roots.TotalSensors{Total: total}, nil
}
