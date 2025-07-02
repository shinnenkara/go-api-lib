package api

import "github.com/gin-gonic/gin"

type Controller interface {
	BindRoutes(Router *gin.RouterGroup)
}
