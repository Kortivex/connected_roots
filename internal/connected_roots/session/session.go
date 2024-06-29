package session

import (
	"context"
	"fmt"

	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql/session"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

const (
	tracingSession        = "service.session"
	tracingSessionSave    = "service.session.save"
	tracingSessionObtain  = "service.session.obtain"
	tracingSessionIsValid = "service.session.is-valid"
)

type Service struct {
	conf   *config.Config
	logger *logger.Logger
	// Repositories
	sessionRep *session.Repository
}

func New(conf *config.Config, db *gorm.DB, logr *logger.Logger) *Service {
	loggerEmpty := logr.NewEmpty()
	logr = loggerEmpty.WithTag(tracingSession)

	return &Service{
		conf:       conf,
		logger:     logr,
		sessionRep: session.NewRepository(conf, db, logr),
	}
}

func (s *Service) Save(ctx context.Context, c echo.Context, sess *connected_roots.Session) (*connected_roots.Session, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingSessionSave)
	defer span.End()

	loggerNew := s.logger.New()
	log := loggerNew.WithTag(tracingSessionSave)

	sn, err := s.sessionRep.Create(ctx, c, sess)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingSessionSave, err)
	}

	log.Debug(fmt.Sprintf("session: %+v", sn))

	return sn, nil
}

func (s *Service) Obtain(ctx context.Context, c echo.Context) (*connected_roots.Session, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingSessionObtain)
	defer span.End()

	loggerNew := s.logger.New()
	log := loggerNew.WithTag(tracingSessionObtain)

	sn, err := s.sessionRep.Get(ctx, c)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingSessionObtain, err)
	}

	log.Debug(fmt.Sprintf("session: %+v", sn))

	return sn, nil
}

func (s *Service) IsValid(ctx context.Context, c echo.Context) (*connected_roots.Session, bool, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingSessionIsValid)
	defer span.End()

	loggerNew := s.logger.New()
	log := loggerNew.WithTag(tracingSessionIsValid)

	sn, err := s.sessionRep.Get(ctx, c)
	if err != nil {
		return nil, false, fmt.Errorf("%s: %w", tracingSessionIsValid, err)
	}

	log.Debug(fmt.Sprintf("session: %+v", sn))

	if sn.ID == "" || sn.Email == "" || sn.UserID == "" {
		return nil, false, nil
	}

	return sn, true, nil
}
