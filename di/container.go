package di

import (
	"github.com/shinnenkara/go-api-lib/api"
)

type Container struct {
	Controller api.Controller
	Providers  []interface{}
}
