package di

import (
	"go-api-lib/api"
	"go-api-lib/logger"
	"gorm.io/gorm"
)

type DbModule interface {
	Init(db *gorm.DB, logger *logger.Logger)
	GetController() api.Controller
}
