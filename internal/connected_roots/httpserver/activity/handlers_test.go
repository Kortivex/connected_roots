package activity

import (
	"testing"

	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/Kortivex/connected_roots/pkg/sdk"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gorm.io/gorm"
)

// Initializes ActivitiesHandlers with valid Context.
func TestNewActivitiesHandlersWithValidContext(t *testing.T) {
	conf := &config.Config{}
	db := &gorm.DB{}
	logr := &logger.Logger{}
	i18n := &i18n.Bundle{}
	sdk := &sdk.ExternalAPI{}

	appCtx := &connected_roots.Context{
		Gorm:   db,
		Logger: logr,
		Conf:   conf,
		I18n:   i18n,
		SDK:    sdk,
	}

	handlers := NewActivitiesHandlers(appCtx)

	if handlers.gorm != db {
		t.Errorf("Expected gorm to be %v, got %v", db, handlers.gorm)
	}

	if handlers.logger == nil {
		t.Error("Expected logger to be initialized")
	}

	if handlers.conf != conf {
		t.Errorf("Expected conf to be %v, got %v", conf, handlers.conf)
	}

	if handlers.activitySvc == nil {
		t.Error("Expected activitySvc to be initialized")
	}

	if handlers.userSvc == nil {
		t.Error("Expected userSvc to be initialized")
	}
}

// Context with nil Gorm should be handled gracefully.
func TestNewActivitiesHandlersWithNilGorm(t *testing.T) {
	conf := &config.Config{}
	logr := &logger.Logger{}
	i18n := &i18n.Bundle{}
	sdk := &sdk.ExternalAPI{}

	appCtx := &connected_roots.Context{
		Gorm:   nil,
		Logger: logr,
		Conf:   conf,
		I18n:   i18n,
		SDK:    sdk,
	}

	handlers := NewActivitiesHandlers(appCtx)

	if handlers.gorm != nil {
		t.Errorf("Expected gorm to be nil, got %v", handlers.gorm)
	}

	if handlers.logger == nil {
		t.Error("Expected logger to be initialized")
	}

	if handlers.conf != conf {
		t.Errorf("Expected conf to be %v, got %v", conf, handlers.conf)
	}

	if handlers.activitySvc == nil {
		t.Error("Expected activitySvc to be initialized")
	}

	if handlers.userSvc == nil {
		t.Error("Expected userSvc to be initialized")
	}
}
