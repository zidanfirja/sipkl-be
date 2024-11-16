package Routers

import (
	Controllers "go-gin-mysql/Controller"
	"go-gin-mysql/Middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	route := gin.Default()

	route.Use(Middleware.CORSMiddleware())

	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"massage": "halaman pertama",
		})
	})

	v1 := route.Group("/sipkl/v1/data/hubin/")
	{
		v1.GET("/role", Controllers.GetAllRole)
		v1.POST("/role", Controllers.CreateNewRole)
		v1.DELETE("/role", Controllers.DeleteRole)
		v1.PUT("/role", Controllers.UpdateRole)

		v1.GET("/industri", Controllers.GetAllIndustri)
		v1.POST("/industri", Controllers.CreateIndustri)
		v1.DELETE("/industri", Controllers.DeleteIndustri)
		v1.PUT("/industri", Controllers.UpdateIndustri)

		v1.GET("/pegawai", Controllers.GetAllPegawai)
		v1.DELETE("/pegawai", Controllers.DeletePegawai)
		v1.PUT("/pegawai", Controllers.UpdatePegawai)

		v1.POST("/pegawai", Controllers.CreatePegawai)
		v1.POST("/pegawai/assign-role", Controllers.AssignRole)

	}
	return route
}
