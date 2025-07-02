package di

import (
	"github.com/shinnenkara/go-api-lib/api"
	"github.com/shinnenkara/go-api-lib/logger"
	"gorm.io/gorm"
)

type DbModule interface {
	Init(db *gorm.DB, logger *logger.Logger)
	GetController() api.Controller
}
