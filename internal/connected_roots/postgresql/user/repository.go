package user

import (
	"context"
	"fmt"

	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

const (
	tracingUser          = "repository-db.user"
	tracingUserGet       = "repository-db.user.get-by"
	tracingUserUpdateAll = "repository-db.user.update-all"
)

type Repository struct {
	conf   *config.Config
	db     *gorm.DB
	logger *logger.Logger
}

func NewRepository(conf *config.Config, db *gorm.DB, logr *logger.Logger) *Repository {
	loggerEmpty := logr.NewEmpty()
	log := loggerEmpty.WithTag(tracingUser)

	return &Repository{
		conf:   conf,
		db:     db,
		logger: log,
	}
}

func (r *Repository) GetBy(ctx context.Context, args ...string) (*connected_roots.Users, error) {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingUserGet)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingUserGet)

	userDB := &Users{}
	result := r.db.WithContext(ctx).Model(&Users{}).
		Preload("Role").
		Where(fmt.Sprintf("%s = ?", args[0]), args[1]).
		First(&userDB)

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingUserGet, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return nil, fmt.Errorf("%s: %w", tracingUserGet, gorm.ErrRecordNotFound)
	}

	log.Debug(fmt.Sprintf("user: %+v", userDB))

	return toDomain(userDB), nil
}

func (r *Repository) UpdateAll(ctx context.Context, user *connected_roots.Users) (*connected_roots.Users, error) {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingUserUpdateAll)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingUserUpdateAll)

	userDB := toDB(user)
	result := r.db.WithContext(ctx).Model(&Users{}).
		Select("*").
		Where(fmt.Sprintf("%s = ?", "email"), user.Email).
		Updates(&userDB)

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingUserUpdateAll, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return nil, fmt.Errorf("%s: %w", tracingUserUpdateAll, gorm.ErrRecordNotFound)
	}

	log.Debug(fmt.Sprintf("user: %+v", userDB))

	return toDomain(userDB), nil
}
