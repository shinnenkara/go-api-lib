package di

import (
	"go-api-lib/api"
	"go-api-lib/logger"
)

type Module interface {
	Init(logger *logger.Logger)
	GetController() api.Controller
}
