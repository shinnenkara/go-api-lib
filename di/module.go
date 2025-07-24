package di

import (
	"github.com/shinnenkara/go-api-lib/api"
)

type Module interface {
	Init(dependencies Dependencies)
	GetController() api.Controller
}
