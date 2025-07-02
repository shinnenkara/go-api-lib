package di

import (
	"github.com/shinnenkara/go-api-lib/logger"
)

type Injector struct {
	Logger *logger.Logger
}

func (injector *Injector) Inject(modules []Module) []Module {
	for _, module := range modules {
		module.Init(injector.Logger)
	}

	return modules
}
