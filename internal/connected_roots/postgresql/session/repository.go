package session

import (
	"context"
	"fmt"

	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

const (
	tracingSession       = "repository-db.session"
	tracingSessionCreate = "repository-db.session.create"
	tracingSessionGet    = "repository-db.session.get"
	tracingSessionDelete = "repository-db.session.delete"
)

type Repository struct {
	conf   *config.Config
	db     *gorm.DB
	logger *logger.Logger
}

func NewRepository(conf *config.Config, db *gorm.DB, logr *logger.Logger) *Repository {
	loggerEmpty := logr.NewEmpty()
	log := loggerEmpty.WithTag(tracingSession)

	return &Repository{
		conf:   conf,
		db:     db,
		logger: log,
	}
}

func (r *Repository) Create(ctx context.Context, c echo.Context, sess *connected_roots.Session) (*connected_roots.Session, error) {
	ctx, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingSessionCreate)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingSessionCreate)

	result, err := session.Get(r.conf.Cookie.Name, c)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingSessionCreate, err)
	}
	result.Values["email"] = sess.Email
	result.Values["user_id"] = sess.UserID
	result.Values["language"] = sess.Language
	result.Values["role_id"] = sess.Role
	if err = result.Save(c.Request(), c.Response()); err != nil {
		return nil, fmt.Errorf("%s: %w", tracingSessionCreate, err)
	}

	sessionDB, err := r.Get(ctx, c)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingSessionCreate, err)
	}

	log.Debug(fmt.Sprintf("session: %+v", result))

	return sessionDB, nil
}

func (r *Repository) Get(ctx context.Context, c echo.Context) (*connected_roots.Session, error) {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingSessionGet)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingSessionGet)

	result, err := session.Get(r.conf.Cookie.Name, c)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingSessionGet, err)
	}

	log.Debug(fmt.Sprintf("session: %+v", result))

	return toDomain(result), nil
}

func (r *Repository) Delete(ctx context.Context, c echo.Context) error {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingSessionDelete)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingSessionDelete)

	result, err := session.Get(r.conf.Cookie.Name, c)
	if err != nil {
		return fmt.Errorf("%s: %w", tracingSessionDelete, err)
	}

	log.Debug(fmt.Sprintf("session: %+v", result))

	cookie, err := c.Cookie(r.conf.Cookie.Name)
	if err != nil {
		return fmt.Errorf("%s: %w", tracingSessionDelete, err)
	}
	cookie.MaxAge = -1
	cookie.Value = ""
	c.SetCookie(cookie)

	log.Debug(fmt.Sprintf("cookie: %+v", cookie))

	deleteRes := r.db.WithContext(ctx).Model(&Sessions{}).
		Unscoped().
		Delete(&result)

	if deleteRes.Error != nil {
		return fmt.Errorf("%s: %w", tracingSessionDelete, deleteRes.Error)
	}

	if deleteRes.Error == nil && deleteRes.RowsAffected == 0 {
		return fmt.Errorf("%s: %w", tracingSessionDelete, gorm.ErrRecordNotFound)
	}

	return nil
}
