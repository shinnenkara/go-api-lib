package api

import (
	"fmt"
	"reflect"
)

func EntityNotFoundError[T any](entity T, id string) Error[any] {
	name := reflect.TypeOf(entity).String()
	message := fmt.Sprintf("Not found %s with Id: %s", name, id)

	notFoundError := GetError[any](404)
	notFoundError.Message = message

	return notFoundError
}
