package api

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type Error[T any] struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
	Errors      []T    `json:"errors"`
}

func GetError[T any](code int) Error[T] {
	commonErrors := commonErrors[T]()

	var apiError Error[T]
	for _, commonError := range commonErrors {
		if commonError.Code == code {
			apiError = commonError
			break
		}
	}
	if apiError.Code == 0 {
		apiError = internalError[T]()
	}

	return apiError
}

func internalError[T any]() Error[T] {
	return Error[T]{
		Code:        500,
		Message:     "Internal Server Error",
		Description: "The server encountered an internal server error",
		Errors:      nil,
	}
}

func commonErrors[T any]() []Error[T] {
	return []Error[T]{
		{
			Code:        400,
			Message:     "Validation failed",
			Description: "One or more fields contain validation errors",
			Errors:      nil,
		},
		{
			Code:        401,
			Message:     "Authentication Error",
			Description: "Access denied to the resource or operation",
			Errors:      nil,
		},
		{
			Code:        403,
			Message:     "Authorization Error",
			Description: "Resource or operation is forbidden",
			Errors:      nil,
		},
		{
			Code:        404,
			Message:     "Not Found",
			Description: "The requested resource does not exist",
			Errors:      nil,
		},
		{
			Code:        405,
			Message:     "Method Not Allowed",
			Description: "The requested method is not allowed on the resource",
			Errors:      nil,
		},
		internalError[T](),
	}
}

func (apiError *Error[any]) ToGinError() gin.Error {
	return gin.Error{
		Err:  errors.New(apiError.Message),
		Type: gin.ErrorType(apiError.Code),
		Meta: apiError,
	}
}
