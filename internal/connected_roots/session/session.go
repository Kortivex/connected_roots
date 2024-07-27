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
	tracingSession                     = "service.session"
	tracingSessionSave                 = "service.session.save"
	tracingSessionObtain               = "service.session.obtain"
	tracingSessionClose                = "service.session.close"
	tracingSessionIsValid              = "service.session.is-valid"
	tracingSessionSaveMessage          = "service.session.save-message"
	tracingSessionObtainMessage        = "service.session.obtain-message"
	tracingSessionIsAdmin              = "service.session.is-admin"
	tracingSessionIsTechnical          = "service.session.is-technical"
	tracingSessionIsUser               = "service.session.is-user"
	tracingSessionIsAdminTechnical     = "service.session.is-admin-technical"
	tracingSessionIsAdminTechnicalUser = "service.session.is-admin-technical-user"
	tracingSessionCountAll             = "service.session.count-all"
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

func (s *Service) Close(ctx context.Context, c echo.Context) error {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingSessionClose)
	defer span.End()

	if err := s.sessionRep.Delete(ctx, c); err != nil {
		return fmt.Errorf("%s: %w", tracingSessionClose, err)
	}

	return nil
}

func (s *Service) SaveMessage(ctx context.Context, c echo.Context, name, value string) error {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingSessionSaveMessage)
	defer span.End()

	loggerNew := s.logger.New()
	log := loggerNew.WithTag(tracingSessionSaveMessage)

	log.Debug(fmt.Sprintf("flash name: %+v", name))
	log.Debug(fmt.Sprintf("flash value: %+v", value))

	if err := s.sessionRep.SetMessage(ctx, c, name, value); err != nil {
		return fmt.Errorf("%s: %w", tracingSessionSaveMessage, err)
	}

	return nil
}

func (s *Service) ObtainMessage(ctx context.Context, c echo.Context, name string) ([]string, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingSessionObtainMessage)
	defer span.End()

	loggerNew := s.logger.New()
	log := loggerNew.WithTag(tracingSessionObtainMessage)

	fm, err := s.sessionRep.GetMessage(ctx, c, name)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingSessionObtainMessage, err)
	}

	log.Debug(fmt.Sprintf("flashes: %+v", fm))

	return fm, nil
}

func (s *Service) IsAdmin(ctx context.Context, c echo.Context) (bool, error) {
	sn, err := s.Obtain(ctx, c)
	if err != nil {
		return false, fmt.Errorf("%s: %w", tracingSessionIsAdmin, err)
	}

	return sn.RoleID == s.conf.Roles.Protected[0], nil
}

func (s *Service) IsTechnical(ctx context.Context, c echo.Context) (bool, error) {
	sn, err := s.Obtain(ctx, c)
	if err != nil {
		return false, fmt.Errorf("%s: %w", tracingSessionIsTechnical, err)
	}

	return sn.RoleID == s.conf.Roles.Protected[1], nil
}

func (s *Service) IsUser(ctx context.Context, c echo.Context) (bool, error) {
	sn, err := s.Obtain(ctx, c)
	if err != nil {
		return false, fmt.Errorf("%s: %w", tracingSessionIsUser, err)
	}

	return sn.RoleID == s.conf.Roles.Protected[2], nil
}

func (s *Service) IsAdminTechnical(ctx context.Context, c echo.Context) (bool, error) {
	sn, err := s.Obtain(ctx, c)
	if err != nil {
		return false, fmt.Errorf("%s: %w", tracingSessionIsAdminTechnical, err)
	}

	return sn.RoleID == s.conf.Roles.Protected[0] || sn.RoleID == s.conf.Roles.Protected[1], nil
}

func (s *Service) IsAdminTechnicalUser(ctx context.Context, c echo.Context) (bool, error) {
	sn, err := s.Obtain(ctx, c)
	if err != nil {
		return false, fmt.Errorf("%s: %w", tracingSessionIsAdminTechnicalUser, err)
	}

	return sn.RoleID == s.conf.Roles.Protected[0] || sn.RoleID == s.conf.Roles.Protected[1] || sn.RoleID == s.conf.Roles.Protected[2], nil
}

func (s *Service) CountAll(ctx context.Context) (*connected_roots.TotalSessions, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingSessionCountAll)
	defer span.End()

	loggerNew := s.logger.New()
	log := loggerNew.WithTag(tracingSessionCountAll)

	total, err := s.sessionRep.Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingSessionCountAll, err)
	}

	log.Debug(fmt.Sprintf("total: %+v", total))

	return &connected_roots.TotalSessions{Total: total}, nil
}
