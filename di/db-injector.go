package di

import (
	"github.com/shinnenkara/go-api-lib/logger"
	"gorm.io/gorm"
)

type DbInjector struct {
	DB     *gorm.DB
	Logger *logger.Logger
}

func (injector *DbInjector) Inject(modules []DbModule) []DbModule {
	for _, module := range modules {
		module.Init(injector.DB, injector.Logger)
	}

	return modules
}
