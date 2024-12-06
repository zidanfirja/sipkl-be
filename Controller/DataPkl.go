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
			// created_at := time.Now()
			dataSiswa.CreatedAt = time.Now()
			dataSiswa.UpdatedAtNilaiPembimbing = nil
			dataSiswa.UpdatedAtNilaiFasilitator = nil

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
		dataSiswa.UpdatedAtNilaiPembimbing = nil
		dataSiswa.UpdatedAtNilaiFasilitator = nil

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
	var dataTanggalMasuk Models.ReqUpdateDataPkl

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
	var dataTanggalKeluar Models.ReqUpdateDataPkl

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

func UpdatePetugasPkl(c *gin.Context) {
	var data Models.ReqUpdateDataPkl

	errBindJson := c.ShouldBindJSON(&data)
	if errBindJson != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errBindJson.Error(),
		})
		return
	}

	// ini unutk mengatur NIS nya
	var listNis []string
	switch dataNis := data.NIS.(type) {
	case string:
		listNis = append(listNis, dataNis)
	case []interface{}:
		for _, item := range dataNis {
			if nis, ok := item.(string); ok {
				listNis = append(listNis, nis)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "data nis invalid",
				})
			}

		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "NIS harus string, atau array string",
		})
		return
	}

	var listUpdatePegawai []Models.UpdatePetugas

	for _, nis := range listNis {
		fkIdPembimbing, ok := data.Payload["fk_id_pembimbing"].(float64)
		// log.Println(reflect.TypeOf(fkIdPembimbing), fkIdPembimbing)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "format id pembimbing tidak valid",
			})
			return
		}
		fkIdFasilitator, ok := data.Payload["fk_id_fasilitator"].(float64)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "format id fasilitator  tidak valid",
			})
			return
		}
		fkIdIndustri, ok := data.Payload["fk_id_industri"].(float64)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "format id industri  tidak valid",
			})
			return
		}
		aktif, ok := data.Payload["aktif"].(bool)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "format aktif  tidak valid",
			})
			return
		}

		updatePetugas := Models.UpdatePetugas{
			NIS:             nis,
			FKIdPembimbing:  int(fkIdPembimbing),
			FKIdFasilitator: int(fkIdFasilitator),
			FKIdIndustri:    int(fkIdIndustri),
			Aktif:           aktif,
		}

		listUpdatePegawai = append(listUpdatePegawai, updatePetugas)
	}

	err := Models.UpdatePengurusPkl(&listUpdatePegawai)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err,
			"message": "gagal update data petugas pkl",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil update data",
	})

}

func DeleteDataSiswaPkl(c *gin.Context) {
	var data Models.ReqDataNis

	errBindJson := c.ShouldBindJSON(&data)
	if errBindJson != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   errBindJson.Error(),
			"massage": "data tidak sesuai singga tidak bisa di binding",
		})
		return
	}

	var listNist []Models.DataNis
	switch dataNis := data.NIS.(type) {
	case string:

		newDataNis := Models.DataNis{
			NIS: dataNis,
		}

		listNist = append(listNist, newDataNis)

	case []interface{}:
		for _, itemNis := range dataNis {
			nis, ok := itemNis.(string)
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "data nis tidak valid",
				})
				return
			}

			newDataNis := Models.DataNis{
				NIS: nis,
			}

			listNist = append(listNist, newDataNis)
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "NIS harus string, atau array string",
		})
		return

	}

	err := Models.DeleteSiswaPkl(&listNist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
		return
	}

	// dataNis = append(dataNis, listNist)
	c.JSON(http.StatusOK, gin.H{
		"massage": "berhasil hapus data",
	})

}

func RawDataPkl(c *gin.Context) {
	data, err := Models.GetRawDataPkl()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}
