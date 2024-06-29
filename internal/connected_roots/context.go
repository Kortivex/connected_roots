package connected_roots

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/Kortivex/connected_roots/pkg/sdk"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gorm.io/gorm"
)

type Context struct {
	Gorm   *gorm.DB
	Logger *logger.Logger
	Conf   *config.Config
	I18n   *i18n.Bundle
	SDK    *sdk.ExternalAPI
}
