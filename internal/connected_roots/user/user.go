package user

import (
	"context"
	"fmt"

	"github.com/Kortivex/connected_roots/pkg/hashing"

	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql/user"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

const (
	tracingUser                = "service.user"
	tracingUserObtainFromEmail = "service.user.obtain-from-email"
	tracingUserObtainFromID    = "service.user.obtain-from-id"
)

type Service struct {
	conf   *config.Config
	logger *logger.Logger
	// Repositories
	userRep *user.Repository
}

func New(conf *config.Config, db *gorm.DB, logr *logger.Logger) *Service {
	loggerEmpty := logr.NewEmpty()
	logr = loggerEmpty.WithTag(tracingUser)

	return &Service{
		conf:    conf,
		logger:  logr,
		userRep: user.NewRepository(conf, db, logr),
	}
}

func (s *Service) ObtainFromID(ctx context.Context, email string) (*connected_roots.Users, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingUserObtainFromID)
	defer span.End()

	loggerNew := s.logger.New()
	log := loggerNew.WithTag(tracingUserObtainFromID)

	usr, err := s.userRep.GetBy(ctx, "id", email)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingUserObtainFromID, err)
	}

	log.Debug(fmt.Sprintf("user: %+v", usr))

	return usr, nil
}

func (s *Service) ObtainFromEmail(ctx context.Context, email string) (*connected_roots.Users, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingUserObtainFromEmail)
	defer span.End()

	loggerNew := s.logger.New()
	log := loggerNew.WithTag(tracingUserObtainFromEmail)

	usr, err := s.userRep.GetBy(ctx, "email", email)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingUserObtainFromEmail, err)
	}

	log.Debug(fmt.Sprintf("user: %+v", usr))

	return usr, nil
}

func (s *Service) IsValidPassword(ctx context.Context, email, password string) (bool, error) {
	_, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingUserObtainFromEmail)
	defer span.End()

	loggerNew := s.logger.New()
	log := loggerNew.WithTag(tracingUserObtainFromEmail)

	usr, err := s.userRep.GetBy(ctx, "email", email)
	if err != nil {
		return false, fmt.Errorf("%s: %w", tracingUserObtainFromEmail, err)
	}

	log.Debug(fmt.Sprintf("user: %+v", usr))
	log.Debug(fmt.Sprintf("password: %+v", password))

	return hashing.PasswordHashingValidation(password, usr.Password), nil
}
