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

	hubin := route.Group("/sipkl/v1/data/hubin/")
	{
		hubin.GET("/role", Controllers.GetAllRole)
		hubin.POST("/role", Controllers.CreateNewRole)
		hubin.DELETE("/role", Controllers.DeleteRole)
		hubin.PUT("/role", Controllers.UpdateRole)

		hubin.GET("/industri", Controllers.GetAllIndustri)
		hubin.POST("/industri", Controllers.CreateIndustri)
		hubin.DELETE("/industri", Controllers.DeleteIndustri)
		hubin.PUT("/industri", Controllers.UpdateIndustri)

		hubin.GET("/pegawai", Controllers.GetAllPegawai)
		hubin.DELETE("/pegawai", Controllers.DeletePegawai)
		hubin.PUT("/pegawai", Controllers.UpdatePegawai)
		hubin.POST("/pegawai", Controllers.CreatePegawai)

		hubin.GET("/pkl", Controllers.GetDataPkl)
		hubin.POST("/pkl", Controllers.NewDataPkl)
		hubin.PUT("/pkl/tanggal_masuk", Controllers.UpdateTanggalMasuk)
		hubin.PUT("/pkl/tanggal_keluar", Controllers.UpdateTanggalKeluar)

		hubin.POST("/pegawai/role", Controllers.AssignRole)
		hubin.DELETE("/pegawai/role", Controllers.DeleteRolePegawai)

	}
	return route
}
