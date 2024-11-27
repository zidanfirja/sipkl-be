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

	route.POST("/sipkl/v1/auth/login", Controllers.Login)

	route.GET("/sipkl/v1/auth/googlelogin", Controllers.LoginOAuth)
	route.GET("/sipkl/v1/auth/callback", Controllers.Callback)

	route.GET("/sipkl/v1/auth/verify", Middleware.CheckAuthToken(), Controllers.PayloadLogin)

	hubin := route.Group("/sipkl/v1/data/hubin/")
	{
		hubin.Use(Middleware.CheckAuthToken())

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
		hubin.PUT("/pkl", Controllers.UpdatePetugasPkl)
		hubin.PUT("/pkl/tanggal_masuk", Controllers.UpdateTanggalMasuk)
		hubin.PUT("/pkl/tanggal_keluar", Controllers.UpdateTanggalKeluar)
		hubin.DELETE("/pkl", Controllers.DeleteDataSiswaPkl)

		hubin.POST("/pegawai/role", Controllers.AssignRole)
		hubin.DELETE("/pegawai/role", Controllers.DeleteRolePegawai)

	}

	nilai := route.Group("/sipkl/v1/data/nilai")

	{

		nilai.Use(Middleware.CheckAuthToken())

		nilai.GET("/industri-pembimbing/:id_pembimbing", Controllers.GetListIndustriPembimbing)
		nilai.GET("/nilai-pembimbing/:id_pembimbing/:id_industri", Controllers.GetNilaiPembimbing)
		nilai.PUT("/nilai-pembimbing/", Controllers.UpdateNilaiPembimbing)

		nilai.GET("/industri-fasilitator/:id_fasilitator", Controllers.GetListIndustriFasilitator)
		nilai.GET("/nilai-fasilitator/:id_fasilitator/:id_industri", Controllers.GetNilaiFasilitator)
		nilai.PUT("/nilai-fasilitator/", Controllers.UpdateNilaiFasilitator)

		nilai.GET("/nilai-walikelas/:kelas/:jurusan/:rombel", Controllers.GetNilaiPklWakel)

	}
	return route
}
