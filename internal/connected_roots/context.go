package connected_roots

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"gorm.io/gorm"
)

type Context struct {
	Gorm   *gorm.DB
	Logger *logger.Logger
	Conf   *config.Config
}
