package activity

import (
	"testing"

	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"gorm.io/gorm"
)

// Initializes Repository with provided config, db, and logger.
func TestInitializeRepositoryWithProvidedConfigDbLogger(t *testing.T) {
	// Given
	conf := &config.Config{}
	db := &gorm.DB{}
	logr := &logger.Logger{}

	// When
	repo := NewRepository(conf, db, logr)

	// Then
	if repo.conf != conf {
		t.Errorf("Expected conf to be %v, got %v", conf, repo.conf)
	}
	if repo.db != db {
		t.Errorf("Expected db to be %v, got %v", db, repo.db)
	}
	if repo.logger == nil {
		t.Error("Expected logger to be initialized, got nil")
	}
}

// Creates a new empty logger instance.
func TestCreatesNewEmptyLoggerInstance(t *testing.T) {
	// Given
	conf := &config.Config{}
	db := &gorm.DB{}
	logr := &logger.Logger{}

	// When
	repo := NewRepository(conf, db, logr)

	// Then
	if repo.logger == logr {
		t.Error("Expected a new logger instance, got the same instance")
	}
}

// Handles nil config input gracefully.
func TestHandlesNilConfigInputGracefully(t *testing.T) {
	// Given
	var conf *config.Config = nil
	db := &gorm.DB{}
	logr := &logger.Logger{}

	// When
	repo := NewRepository(conf, db, logr)

	// Then
	if repo.conf != nil {
		t.Errorf("Expected conf to be nil, got %v", repo.conf)
	}
}

// Handles nil db input gracefully.
func TestHandlesNilDbInputGracefully(t *testing.T) {
	// Given
	conf := &config.Config{}
	var db *gorm.DB = nil
	logr := &logger.Logger{}

	// When
	repo := NewRepository(conf, db, logr)

	// Then
	if repo.db != nil {
		t.Errorf("Expected db to be nil, got %v", repo.db)
	}
}
