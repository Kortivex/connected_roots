package postgresql

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/thejerf/suture/v4"

	"github.com/Kortivex/connected_roots/pkg/logger/commons"

	"github.com/Kortivex/connected_roots/pkg/logger"
	gormlogger "github.com/Kortivex/connected_roots/pkg/logger/extend/gorm"

	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/pkg/service"
	pgmigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/opentelemetry/tracing"
)

//go:embed migrations/*
var content embed.FS

const (
	debug = "debug"

	defaultIOSource        = "iofs"
	defaultDBType          = "postgres"
	defaultDBMigrationPath = "migrations"

	searchPathKey = "search_path"

	ErrMsgCannotOpenDBConnection   = "cannot open db connection"
	ErrMsgCannotObtainBDConnection = "cannot get postgresql db connection"
	ErrMsgDBIsNotReachable         = "postgresql db is not reachable"

	ErrMsgParseQuery        = "parse query wrong format"
	ErrMsgSearchPathMissing = "search_path missing"
	ErrMsgMalformedDSN      = "malformed DSN"

	ErrMsgCannotExtractSchemaFromDSN            = "cannot extract schema from dsn"
	ErrMsgCannotOpeningMigrationFolder          = "cannot opening migration folder"
	ErrMsgCannotCreateMigrationPostgresConfig   = "cannot create postgresql migration config instance"
	ErrMsgCannotCreateMigrationPostgresInstance = "cannot create postgresql migration instance"
	ErrMsgCannotApplyMigrationDB                = "cannot apply database migration"

	InfMsgNoChangesInDBSchema = "no changes in database schema"
)

type Service struct {
	service.Service
	Gorm   *gorm.DB
	DB     *sql.DB
	logger *logger.Logger
	conf   *config.Config
}

// NewService This function creates and returns a new Service struct with the given parameters.
func NewService(name string, conf *config.Config, logr *logger.Logger) *Service {
	loggerEmpty := logr.NewEmpty()
	log := loggerEmpty.WithTag(commons.TagPlatformGorm)

	srv := &Service{
		Service: service.Service{
			Name:     name,
			Started:  make(chan bool),
			Status:   make(chan int),
			Release:  make(chan bool),
			Stop:     make(chan bool),
			Existing: 0,
			M:        sync.Mutex{},
			Running:  false,
		},
		logger: log,
		conf:   conf,
	}

	return srv
}

// provide open the gorm.DB and set the sql.DB to use in the Service.
func (s *Service) provide() {
	// Generic DB ORM Config.
	gormConfig := &gorm.Config{
		PrepareStmt: true,
	}
	// Setup DB Logger.
	setLogger(s.conf.DB, s.logger, gormConfig)
	// Setup DB Naming Strategy.
	dbSchema, err := getDBSchemaName(s.conf.DB.Postgres.DSN)
	if err != nil {
		s.logger.Fatal(fmt.Errorf(ErrMsgCannotExtractSchemaFromDSN+": %w", err).Error())
	}
	setNameStrategy(dbSchema, false, gormConfig)

	// PostgresDB Driver Config.
	postgresConfig := postgres.Config{
		DSN: s.conf.DB.Postgres.DSN,
	}

	s.logger.Debug("Establishing Connection")
	// Setup postgresql connection.
	db, err := gorm.Open(postgres.New(postgresConfig), gormConfig)
	if err != nil {
		s.logger.Fatal(fmt.Errorf(ErrMsgCannotOpenDBConnection+": %w", err).Error())
	}

	// Set telemetry for gorm db.
	if err = db.Use(tracing.NewPlugin()); err != nil {
		s.logger.Fatal(err.Error())
	}

	// Set Connections Properties.
	pgDB, err := db.DB()
	if err != nil {
		s.logger.Fatal(fmt.Errorf(ErrMsgCannotObtainBDConnection+": %w", err).Error())
	}
	setConnectionProps(s.conf.DB, pgDB)
	s.DB = pgDB

	// Setup debug mode for GORM.
	if strings.EqualFold(s.conf.App.LogLevel, debug) {
		db = db.Debug()
	}

	// Ping Connection
	err = ping(pgDB)
	if err != nil {
		s.logger.Fatal(err.Error())
	}

	s.logger.Debug("Connection Established")

	s.Gorm = db
}

// doMigrations applies the corresponding migrations to the database schema.
func (s *Service) doMigrations() {
	source, err := iofs.New(content, defaultDBMigrationPath)
	if err != nil {
		s.logger.Fatal(fmt.Errorf(ErrMsgCannotOpeningMigrationFolder+": %w", err).Error())
	}

	driver, err := pgmigrate.WithInstance(s.DB, &pgmigrate.Config{})
	if err != nil {
		s.logger.Fatal(fmt.Errorf(ErrMsgCannotCreateMigrationPostgresConfig+": %w", err).Error())
	}

	m, err := migrate.NewWithInstance(defaultIOSource, source, defaultDBType, driver)
	if err != nil {
		s.logger.Fatal(fmt.Errorf(ErrMsgCannotCreateMigrationPostgresInstance+": %w", err).Error())
	}

	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		s.logger.Info(InfMsgNoChangesInDBSchema)
	} else if err != nil {
		s.logger.Fatal(fmt.Errorf(ErrMsgCannotApplyMigrationDB+": %w", err).Error())
	}
}

// getDBSchemaName get the schema db name from the DSN string.
func getDBSchemaName(dsn string) (string, error) {
	connQs := strings.Split(dsn, "?")
	if len(connQs) == 2 {
		qs, err := url.ParseQuery(connQs[1])
		if err != nil {
			return "", fmt.Errorf(ErrMsgParseQuery+": %w", err)
		}
		searchPath := qs.Get(searchPathKey)
		if searchPath == "" {
			return "", fmt.Errorf(ErrMsgSearchPathMissing+": %w", errors.New(dsn))
		}

		return searchPath, nil
	}

	return "", fmt.Errorf(ErrMsgMalformedDSN+": %w", errors.New(dsn))
}

// setNameStrategy apply tables and columns naming strategy.
func setNameStrategy(tableNamePrefix string, singularTable bool, gormConfig *gorm.Config) {
	tableNamePrefix += "."
	nameStrategy := schema.NamingStrategy{
		TablePrefix:   tableNamePrefix,
		SingularTable: singularTable,
	}
	gormConfig.NamingStrategy = nameStrategy
}

// setLogger apply logger config from logger.Logger to gormLogger.
func setLogger(dbConf config.DB, logr *logger.Logger, gormConfig *gorm.Config) {
	slowThreshold := time.Duration(dbConf.Postgres.Logger.SlowThreshold) * time.Second
	ignoreRecordNotFound := dbConf.Postgres.Logger.IgnoreRecordNotFoundError

	gormLogger := gormlogger.NewLogger(logr, slowThreshold, ignoreRecordNotFound)

	gormConfig.Logger = &gormLogger
}

// setConnectionProps set connections properties to postgresql driver.
func setConnectionProps(dbConf config.DB, pgDB *sql.DB) {
	pgDB.SetMaxIdleConns(dbConf.Postgres.Connection.MaxIdleConns)
	pgDB.SetMaxOpenConns(dbConf.Postgres.Connection.MaxOpenConns)
	pgDB.SetConnMaxIdleTime(time.Duration(dbConf.Postgres.Connection.ConnMaxIdleTime) * time.Minute)
	pgDB.SetConnMaxLifetime(time.Duration(dbConf.Postgres.Connection.ConnMaxLifetime) * time.Minute)
}

// ping to check if the database works.
func ping(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		return fmt.Errorf(ErrMsgDBIsNotReachable+": %w", err)
	}

	return nil
}

// Serve method from interface suture.Service to handle Service cycle life.
func (s *Service) Serve(ctx context.Context) error {
	s.M.Lock()
	if s.Existing != 0 {
		(&sync.Mutex{}).Unlock()
	}
	s.Existing++
	s.Running = true
	s.M.Unlock()

	defer func() {
		s.M.Lock()
		s.Running = false
		s.M.Unlock()
	}()

	releaseExistence := func() {
		s.M.Lock()
		s.Existing--
		s.M.Unlock()
	}

	s.Started <- true

	useStopChan := false

	for {
		select {
		case val := <-s.Status:
			switch val {
			case service.Run:
				// Start Service
				s.provide()
				// Apply DB Migrations
				s.doMigrations()
			case service.Heartbeat:
				go func() {
					ticker := time.NewTicker(time.Duration(s.conf.DB.Postgres.Health.Frequency) * time.Second)
					for range ticker.C {
						if err := ping(s.DB); err != nil {
							s.logger.Debug(service.PingKO)
							s.logger.Error(err.Error())
							releaseExistence()
							os.Exit(1)
						}
						s.logger.Debug(service.PingOK)
					}
				}()
			case service.Fail:
				releaseExistence()
				if useStopChan {
					s.Stop <- true
				}
				return nil
			case service.Panic:
				releaseExistence()
				panic(service.ErrPanicService)
			case service.Hang:
				<-s.Release
			case service.UseStopChan:
				useStopChan = true
			case service.TerminateTree:
				return suture.ErrTerminateSupervisorTree
			case service.DoNotRestart:
				return suture.ErrDoNotRestart
			}
		case <-ctx.Done():
			releaseExistence()
			if useStopChan {
				s.Stop <- true
			}
			return fmt.Errorf(service.ErrFailureServiceEnding+": %w", ctx.Err())
		}
	}
}
