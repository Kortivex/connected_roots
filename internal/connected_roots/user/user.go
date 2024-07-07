package user

import (
	"context"
	"fmt"
	"github.com/Kortivex/connected_roots/pkg/pagination"

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
	tracingUserSave            = "service.user.save"
	tracingUserUpdate          = "service.user.update"
	tracingUserObtainFromEmail = "service.user.obtain-from-email"
	tracingUserObtainFromID    = "service.user.obtain-from-id"
	tracingUserObtainAll       = "service.user.obtain-all"
	tracingUserRemoveByID      = "service.user.remove-by-id"
	tracingUserRemoveByEmail   = "service.user.remove-by-email"
	tracingUserIsValidPassword = "service.user.is-valid-password"
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

func (s *Service) Save(ctx context.Context, user *connected_roots.Users) (*connected_roots.Users, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingUserSave)
	defer span.End()

	passwordHashing, err := hashing.PasswordHashing(user.Password)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingUserSave, err)
	}
	user.Password = string(passwordHashing)

	userRes, err := s.userRep.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingUserSave, err)
	}

	userRes, err = s.userRep.GetByID(ctx, userRes.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingUserSave, err)
	}

	return userRes, nil
}

func (s *Service) Update(ctx context.Context, user *connected_roots.Users) (*connected_roots.Users, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingUserUpdate)
	defer span.End()

	loggerNew := s.logger.New()
	log := loggerNew.WithTag(tracingUserUpdate)

	if user.Password != "" {
		passwordHashing, err := hashing.PasswordHashing(user.Password)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", tracingUserUpdate, err)
		}
		user.Password = string(passwordHashing)
	}

	usr, err := s.userRep.UpdateAll(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingUserUpdate, err)
	}

	log.Debug(fmt.Sprintf("user: %+v", usr))

	return usr, nil
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

func (s *Service) ObtainAll(ctx context.Context, filters *connected_roots.UserPaginationFilters) (*pagination.Pagination, error) {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingUserObtainAll)
	defer span.End()

	rolesRes, err := s.userRep.ListAllBy(ctx, filters, []string{"Role"}...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingUserObtainAll, err)
	}

	return rolesRes, nil
}

func (s *Service) RemoveByID(ctx context.Context, id string) error {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingUserRemoveByID)
	defer span.End()

	if err := s.userRep.DeleteByID(ctx, id); err != nil {
		return fmt.Errorf("%s: %w", tracingUserRemoveByID, err)
	}

	return nil
}

func (s *Service) RemoveByEmail(ctx context.Context, email string) error {
	ctx, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingUserRemoveByEmail)
	defer span.End()

	if err := s.userRep.DeleteByEmail(ctx, email); err != nil {
		return fmt.Errorf("%s: %w", tracingUserRemoveByEmail, err)
	}

	return nil
}

func (s *Service) IsValidPassword(ctx context.Context, email, password string) (bool, error) {
	_, span := otel.Tracer(s.conf.App.Name).Start(ctx, tracingUserIsValidPassword)
	defer span.End()

	loggerNew := s.logger.New()
	log := loggerNew.WithTag(tracingUserIsValidPassword)

	usr, err := s.userRep.GetBy(ctx, "email", email)
	if err != nil {
		return false, fmt.Errorf("%s: %w", tracingUserIsValidPassword, err)
	}

	log.Debug(fmt.Sprintf("user: %+v", usr))
	log.Debug(fmt.Sprintf("password: %+v", password))

	return hashing.PasswordHashingValidation(password, usr.Password), nil
}
