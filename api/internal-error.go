package api

import "fmt"

func NewInternalServerError(err error) Error[any] {
	fmt.Println(err.Error())
	serverError := GetError[any](500)
	return serverError
}
