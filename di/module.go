package di

import (
	"github.com/shinnenkara/go-api-lib/api"
	"github.com/shinnenkara/go-api-lib/logger"
)

type Module interface {
	Init(logger *logger.Logger)
	GetController() api.Controller
}
