package Controllers

import (
	"go-gin-mysql/Models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetDataPkl(c *gin.Context) {
	data, err := Models.GetDataPkl()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengambil data",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}

func AddDataPkl(c *gin.Context) {
	var reqDataSiswa Models.ReqAddDataSiswa
	var dataSiswa Models.DataSiswa

	if err := c.ShouldBindJSON(&reqDataSiswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "input data yang valid",
			"error":   err.Error(),
		})
		return
	}

	layout := "2006-01-02"

	// parse data tanggal masuk dari inputan string bisa masuk ke model lalu ke db
	if reqDataSiswa.TanggalMasuk == "" {
		dataSiswa.TanggalMasuk = nil

	} else {

		parsedTanggalMasuk, errParseMasuk := time.Parse(layout, reqDataSiswa.TanggalMasuk)
		if errParseMasuk != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "masukan format data tanggal masuk yang benar",
			})
		}
		dataSiswa.TanggalMasuk = &parsedTanggalMasuk
	}

	// parse data tanggal masuk dari inputan string bisa masuk ke model lalu ke db
	if reqDataSiswa.TanggalKeluar == "" {

		dataSiswa.TanggalKeluar = nil

	} else {
		parsedTanggalKeluar, errParseKeluar := time.Parse(layout, reqDataSiswa.TanggalKeluar)
		if errParseKeluar != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "masukan format data tanggal masuk yang benar",
			})
		}

		dataSiswa.TanggalKeluar = &parsedTanggalKeluar

	}

	dataSiswa.NIS = reqDataSiswa.NIS
	dataSiswa.Nama = reqDataSiswa.Nama
	dataSiswa.Kelas = reqDataSiswa.Kelas
	dataSiswa.Jurusan = reqDataSiswa.Jurusan
	dataSiswa.Rombel = reqDataSiswa.Rombel
	dataSiswa.FKIdFasilitator = reqDataSiswa.FKIdFasilitator
	dataSiswa.FKIdIndustri = reqDataSiswa.FKIdIndustri
	dataSiswa.FKIdPembimbing = reqDataSiswa.FKIdPembimbing
	dataSiswa.Aktif = true
	created_at := time.Now()
	dataSiswa.CreatedAt = created_at

	// c.JSON(http.StatusInternalServerError, gin.H{
	// 	"data": dataSiswa,
	// })
	// return

	err := Models.AddDataPkl(&dataSiswa)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal menginput data",
			"error":   err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "input data pkl berhasil",
	})

}
