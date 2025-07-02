package di

import "go-api-lib/api"

type Container struct {
	Controller api.Controller
	Providers  []interface{}
}
