package api

import "github.com/gin-gonic/gin"

type Endpoint[T any, E any] func(c *gin.Context) (T, Error[E])

func ProcessRequest[T any, E any](context *gin.Context, successCode int, endpoint Endpoint[T, E]) {
	response, err := endpoint(context)
	if err.Code != 0 {
		ProduceError(context, err)
		return
	}

	context.JSON(successCode, response)
}

func ProduceError[T any](context *gin.Context, apiError Error[T]) {
	ginError := apiError.ToGinError()
	context.Errors = append(context.Errors, &ginError)
}
