package Controllers

import (
	"encoding/json"
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
func handleZeroTime(dateString string) (*time.Time, error) {
	// Jika string kosong, kembalikan nil
	if dateString == "" {
		return nil, nil
	}

	// Tentukan layout sesuai dengan format input (yyyy-mm-dd)
	layout := "2006-01-02"
	// Parsing string ke time.Time
	parsedTime, err := time.Parse(layout, dateString)
	if err != nil {
		return nil, err
	}
	return &parsedTime, nil
}

func NewDataPkl(c *gin.Context) {
	var data interface{}
	var siswaList []Models.DataSiswa
	var siswa Models.ReqAddDataSiswa
	var dataSiswa Models.DataSiswa

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	switch dataPkl := data.(type) {
	case []interface{}:

		for _, item := range dataPkl {
			// Convert setiap item ke JSON dan bind ke struct
			itemBytes, _ := json.Marshal(item)

			if err := json.Unmarshal(itemBytes, &siswa); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data in array"})
				return
			}

			tglMasuk, _ := handleZeroTime(siswa.TanggalMasuk)
			dataSiswa.TanggalMasuk = tglMasuk

			tglKeluar, _ := handleZeroTime(siswa.TanggalKeluar)
			dataSiswa.TanggalKeluar = tglKeluar

			dataSiswa.NIS = siswa.NIS
			dataSiswa.Nama = siswa.Nama
			dataSiswa.Kelas = siswa.Kelas
			dataSiswa.Jurusan = siswa.Jurusan
			dataSiswa.Rombel = siswa.Rombel
			dataSiswa.FKIdFasilitator = siswa.FKIdFasilitator
			dataSiswa.FKIdIndustri = siswa.FKIdIndustri
			dataSiswa.FKIdPembimbing = siswa.FKIdPembimbing
			dataSiswa.Aktif = true
			created_at := time.Now()
			dataSiswa.CreatedAt = created_at

			siswaList = append(siswaList, dataSiswa)
		}
		// Process the list of students
		err := Models.AddMultipleDataPkl(&siswaList)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Gagal menginput data",
				"error":   err.Error(),
			})
			return

		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Berhasil input data",
		})
		return

	case map[string]interface{}:
		// Handle a single object
		itemBytes, _ := json.Marshal(dataPkl)

		if err := json.Unmarshal(itemBytes, &siswa); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid single object data",
			})
			return
		}

		tglMasuk, _ := handleZeroTime(siswa.TanggalMasuk)
		dataSiswa.TanggalMasuk = tglMasuk

		tglKeluar, _ := handleZeroTime(siswa.TanggalKeluar)
		dataSiswa.TanggalKeluar = tglKeluar

		dataSiswa.NIS = siswa.NIS
		dataSiswa.Nama = siswa.Nama
		dataSiswa.Kelas = siswa.Kelas
		dataSiswa.Jurusan = siswa.Jurusan
		dataSiswa.Rombel = siswa.Rombel
		dataSiswa.FKIdFasilitator = siswa.FKIdFasilitator
		dataSiswa.FKIdIndustri = siswa.FKIdIndustri
		dataSiswa.FKIdPembimbing = siswa.FKIdPembimbing
		dataSiswa.Aktif = true
		created_at := time.Now()
		dataSiswa.CreatedAt = created_at

		// proses singele data
		err := Models.AddDataPkl(&dataSiswa)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Gagal menginput data",
				"error":   err.Error(),
			})
			return

		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Berhasil input data",
		})
		return

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

}

// func AddDataPkl(c *gin.Context) {
// 	var reqDataSiswa Models.ReqAddDataSiswa
// 	var dataSiswa Models.DataSiswa

// 	if err := c.ShouldBindJSON(&reqDataSiswa); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "input data yang valid",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	dataSiswa.TanggalMasuk, _ = handleZeroTime(reqDataSiswa.TanggalMasuk)
// 	dataSiswa.TanggalMasuk, _ = handleZeroTime(reqDataSiswa.TanggalMasuk)

// 	dataSiswa.NIS = reqDataSiswa.NIS
// 	dataSiswa.Nama = reqDataSiswa.Nama
// 	dataSiswa.Kelas = reqDataSiswa.Kelas
// 	dataSiswa.Jurusan = reqDataSiswa.Jurusan
// 	dataSiswa.Rombel = reqDataSiswa.Rombel
// 	dataSiswa.FKIdFasilitator = reqDataSiswa.FKIdFasilitator
// 	dataSiswa.FKIdIndustri = reqDataSiswa.FKIdIndustri
// 	dataSiswa.FKIdPembimbing = reqDataSiswa.FKIdPembimbing
// 	dataSiswa.Aktif = true
// 	created_at := time.Now()
// 	dataSiswa.CreatedAt = created_at

// 	err := Models.AddDataPkl(&dataSiswa)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "Gagal menginput data",
// 			"error":   err.Error(),
// 		})
// 		return

// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "input data pkl berhasil",
// 	})

// }

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
