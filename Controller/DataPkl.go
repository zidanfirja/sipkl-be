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

// untuk handle inputan tanggal string menjadi time.time
func handleZeroTime(tanggal string) (*time.Time, error) {
	layout := "2006-01-02"
	dataTanggal, errParse := time.Parse(layout, tanggal)
	if errParse != nil {
		return nil, errParse
	}
	return &dataTanggal, errParse

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

	// layout := "2006-01-02"

	// parse data tanggal masuk dari inputan string bisa masuk ke model lalu ke db
	// if reqDataSiswa.TanggalMasuk == "" {
	// 	dataSiswa.TanggalMasuk = nil

	// } else {

	// 	parsedTanggalMasuk, errParseMasuk := time.Parse(layout, reqDataSiswa.TanggalMasuk)
	// 	if errParseMasuk != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{
	// 			"error": "masukan format data tanggal masuk yang benar",
	// 		})
	// 	}
	// }
	// tanggal, errTanggalMasuk := handleZeroTime(reqDataSiswa.TanggalMasuk)
	//
	dataSiswa.TanggalMasuk, _ = handleZeroTime(reqDataSiswa.TanggalMasuk)

	// _, errTanggalKeluar := handleZeroTime(reqDataSiswa.TanggalKeluar)
	// if errTanggalKeluar != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "input data tidak valid",
	// 	})
	// 	return
	// }
	dataSiswa.TanggalMasuk, _ = handleZeroTime(reqDataSiswa.TanggalMasuk)

	// dataSiswa.TanggalKeluar = handleZeroTime(reqDataSiswa.TanggalKeluar)

	// // parse data tanggal masuk dari inputan string bisa masuk ke model lalu ke db
	// if reqDataSiswa.TanggalKeluar == "" {

	// 	dataSiswa.TanggalKeluar = nil

	// } else {
	// 	parsedTanggalKeluar, errParseKeluar := time.Parse(layout, reqDataSiswa.TanggalKeluar)
	// 	if errParseKeluar != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{
	// 			"error": "masukan format data tanggal masuk yang benar",
	// 		})
	// 	}

	// }

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

func UpdateTanggalMasuk(c *gin.Context) {
	var dataTanggalMasuk Models.ReqUpdateTanggal

	errBindJson := c.ShouldBindJSON(&dataTanggalMasuk)
	if errBindJson != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error saat binding data",
			"error":   errBindJson.Error(),
		})
		return
	}

	// cek payload bersi tanggal_masuk
	tanggal_masuk, ok := dataTanggalMasuk.Payload["tanggal_masuk"]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "data payload tidak berisi tanggal_masuk",
		})
		return
	}

	// cek isi tanggal masuk
	tanggal, ok := tanggal_masuk.(string)
	if !ok || tanggal == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Field 'jurusan' harus string.",
		})
		return
	}

	//parse tanggal
	tanggalTime, err := handleZeroTime(tanggal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "input data tidak valid",
		})
	}

	var daftarNis []string
	switch nis := dataTanggalMasuk.NIS.(type) {
	case string:
		daftarNis = append(daftarNis, nis)
		err := Models.UpdateTanggalMasuk(daftarNis, *tanggalTime)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Gagal update data",
				"error":   err.Error(),
			})
			return
		}
	case []interface{}:

		for _, item := range nis {
			if dataNis, ok := item.(string); ok {
				daftarNis = append(daftarNis, dataNis)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "format data id salah",
				})
				return
			}
		}
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"message": daftarNis,
		// })
		// return

		err := Models.UpdateTanggalMasuk(daftarNis, *tanggalTime)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error saat update multiple",
				"error":   err.Error(),
			})
			return
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"massage": "berhasil update data tanggal masuk",
	})

}

func UpdateTanggalKeluar(c *gin.Context) {
	var dataTanggalKeluar Models.ReqUpdateTanggal

	errBindJson := c.ShouldBindJSON(&dataTanggalKeluar)
	if errBindJson != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error saat binding data",
			"error":   errBindJson.Error(),
		})
		return
	}

	// cek payload bersi tanggal_keluar
	tanggal_keluar, ok := dataTanggalKeluar.Payload["tanggal_keluar"]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "data payload tidak berisi tanggal_keluar",
		})
		return
	}

	// cek isi tanggal keluar
	tanggal, ok := tanggal_keluar.(string)
	if !ok || tanggal == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Field 'jurusan' harus string.",
		})
		return
	}

	//parse tanggal
	tanggalTime, err := handleZeroTime(tanggal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "input data tidak valid",
		})
		return
	}

	var daftarNis []string
	switch nis := dataTanggalKeluar.NIS.(type) {
	case string:

		daftarNis = append(daftarNis, nis)
		err := Models.UpdateTanggalKeluar(daftarNis, *tanggalTime)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Gagal update data",
				"error":   err.Error(),
			})
			return
		}
	case []interface{}:

		for _, item := range nis {
			if dataNis, ok := item.(string); ok {
				daftarNis = append(daftarNis, dataNis)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "format data id salah",
				})
				return
			}
		}

		err := Models.UpdateTanggalKeluar(daftarNis, *tanggalTime)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error saat update multiple",
				"error":   err.Error(),
			})
			return
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"massage": "berhasil update data tanggal keluar",
	})

}
