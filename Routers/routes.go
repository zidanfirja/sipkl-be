package Routers

import (
	Controllers "go-gin-mysql/Controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	route := gin.Default()

	v1 := route.Group("/sipkl/v1/data/hubin/")
	{
		v1.GET("/role", Controllers.GetAllRole)
		v1.POST("/role", Controllers.CreateNewRole)
	}
	return route
}
