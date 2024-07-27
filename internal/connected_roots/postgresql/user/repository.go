package user

import (
	"context"
	"fmt"
	"github.com/Kortivex/connected_roots/pkg/pagination"
	"github.com/Kortivex/connected_roots/pkg/ulid"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

const (
	tracingUser              = "repository-db.user"
	tracingUserCreate        = "repository-db.user.create"
	tracingUserUpdateAll     = "repository-db.user.update-all"
	tracingUserGetByID       = "repository-db.user.get-by-id"
	tracingUserGetBy         = "repository-db.user.get-by"
	tracingUserListAllBy     = "repository-db.user.list-all-by"
	tracingUserDeleteByID    = "repository-db.user.delete-by-id"
	tracingUserDeleteByEmail = "repository-db.user.delete-by-email"
	tracingUserCount         = "repository-db.user.count"
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

func (r *Repository) Create(ctx context.Context, user *connected_roots.Users) (*connected_roots.Users, error) {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingUserCreate)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingUserCreate)

	log.Debug(fmt.Sprintf("user: %+v", user))

	id, err := ulid.Generate()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingUserCreate, err)
	}

	userDB := toDB(user, id)
	result := r.db.WithContext(ctx).Model(&Users{}).
		Create(&userDB)

	if result.Error != nil {
		if result.Error.(*pgconn.PgError).Code == "23505" {
			return nil, fmt.Errorf("%s: %w", tracingUserCreate, gorm.ErrDuplicatedKey)
		}
		return nil, fmt.Errorf("%s: %w", tracingUserCreate, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return nil, fmt.Errorf("%s: %w", tracingUserCreate, gorm.ErrRecordNotFound)
	}

	log.Debug(fmt.Sprintf("user: %+v", userDB))

	return toDomain(userDB), nil
}

func (r *Repository) UpdateAll(ctx context.Context, user *connected_roots.Users) (*connected_roots.Users, error) {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingUserUpdateAll)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingUserUpdateAll)

	userDB := toDB(user, user.ID)
	query := r.db.WithContext(ctx).Model(&Users{}).
		Select("*")

	if user.Password == "" {
		query = query.Omit("password")
	}

	query = query.Where(fmt.Sprintf("%s = ?", "email"), user.Email)
	result := query.Updates(&userDB)

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingUserUpdateAll, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return nil, fmt.Errorf("%s: %w", tracingUserUpdateAll, gorm.ErrRecordNotFound)
	}

	log.Debug(fmt.Sprintf("user: %+v", userDB))

	return toDomain(userDB), nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (*connected_roots.Users, error) {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingUserGetByID)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingUserGetByID)

	log.Debug(fmt.Sprintf("user id: %+v", id))

	userDB := &Users{}
	result := r.db.WithContext(ctx).Model(&Users{}).
		Where("id = ?", id).
		First(&userDB)

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingUserGetByID, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return nil, fmt.Errorf("%s: %w", tracingUserGetByID, gorm.ErrRecordNotFound)
	}

	log.Debug(fmt.Sprintf("user: %+v", userDB))

	return toDomain(userDB), nil
}

func (r *Repository) GetBy(ctx context.Context, args ...string) (*connected_roots.Users, error) {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingUserGetBy)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingUserGetBy)

	userDB := &Users{}
	result := r.db.WithContext(ctx).Model(&Users{}).
		Preload("Role").
		Where(fmt.Sprintf("%s = ?", args[0]), args[1]).
		First(&userDB)

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingUserGetBy, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return nil, fmt.Errorf("%s: %w", tracingUserGetBy, gorm.ErrRecordNotFound)
	}

	log.Debug(fmt.Sprintf("user: %+v", userDB))

	return toDomain(userDB), nil
}

func (r *Repository) ListAllBy(ctx context.Context, rolesFilters *connected_roots.UserPaginationFilters, preloads ...string) (*pagination.Pagination, error) {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingUserListAllBy)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingUserListAllBy)

	log.Debug(fmt.Sprintf("filters: %+v", rolesFilters))

	rulesBuilder := pagination.SortRulesBuilder{
		Sorts:               rolesFilters.Sort,
		ValidateFields:      TableUsersFields,
		DBStructAssociation: TableUsersSortMap,
		TableName:           TableUsersName,
	}

	rules, err := rulesBuilder.ObtainRules()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingUserListAllBy, err)
	}

	if len(rules) == 0 {
		rules = append(rules, DefaultUserRule)
	}

	pg, err := pagination.CreatePaginator(&rolesFilters.PaginatorParams, rules)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingUserListAllBy, err)
	}

	var usersDB []*Users
	query := r.db.WithContext(ctx).Model(&Users{})
	for _, p := range preloads {
		query = query.Preload(p)
	}
	AddRoleFilters(query, &rolesFilters.UserFilters)
	query.Find(&usersDB)

	result, cursor, err := pg.Paginate(query, &usersDB)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingUserListAllBy, err)
	}

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", tracingUserListAllBy, result.Error)
	}

	previousCursor, nextCursor, err := pagination.EncodeURLValues(cursor)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingUserListAllBy, err)
	}

	usersPaginated := &pagination.Pagination{
		Data: toDomainSlice(usersDB),
		Paging: pagination.Paging{
			NextCursor:     nextCursor,
			PreviousCursor: previousCursor,
		},
	}

	log.Debug(fmt.Sprintf("users: %+v", usersDB))

	return usersPaginated, nil
}

func (r *Repository) DeleteByID(ctx context.Context, id string) error {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingUserDeleteByID)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingUserDeleteByID)

	log.Debug(fmt.Sprintf("user id: %+v", id))

	userDB := &Users{ID: id}
	result := r.db.WithContext(ctx).Model(&Users{}).
		Unscoped().
		Where("id = ?", id).
		Delete(&userDB)

	if result.Error != nil {
		return fmt.Errorf("%s: %w", tracingUserDeleteByID, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return fmt.Errorf("%s: %w", tracingUserDeleteByID, gorm.ErrRecordNotFound)
	}

	return nil
}

func (r *Repository) DeleteByEmail(ctx context.Context, email string) error {
	_, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingUserDeleteByEmail)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingUserDeleteByEmail)

	log.Debug(fmt.Sprintf("user email: %+v", email))

	userDB := &Users{}
	result := r.db.WithContext(ctx).Model(&Users{}).
		Unscoped().
		Where("email = ?", email).
		Delete(&userDB)

	if result.Error != nil {
		return fmt.Errorf("%s: %w", tracingUserDeleteByEmail, result.Error)
	}

	if result.Error == nil && result.RowsAffected == 0 {
		return fmt.Errorf("%s: %w", tracingUserDeleteByEmail, gorm.ErrRecordNotFound)
	}

	return nil
}

func (r *Repository) Count(ctx context.Context) (int64, error) {
	ctx, span := otel.Tracer(r.conf.App.Name).Start(ctx, tracingUserCount)
	defer span.End()

	loggerNew := r.logger.New()
	log := loggerNew.WithTag(tracingUserCount)

	var total int64
	result := r.db.WithContext(ctx).Model(&Users{}).
		Count(&total)

	if result.Error != nil {
		return 0, fmt.Errorf("%s: %w", tracingUserCount, result.Error)
	}

	log.Debug(fmt.Sprintf("total: %+v", total))

	return total, nil
}
