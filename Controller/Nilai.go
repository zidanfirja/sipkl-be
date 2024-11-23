package Controllers

import (
	"go-gin-mysql/Models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetListIndustriPembimbing(c *gin.Context) {
	param := c.Param("id_pembimbing")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "id hasul",
		})
		return
	}

	data, err := Models.GetIndustriPembimbing(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func GetListIndustriFasilitator(c *gin.Context) {
	param := c.Param("id_fasilitator")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "id tidak valid",
		})
		return
	}

	data, err := Models.GetIndustriFasilitator(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

type NilaiPklPembimbing struct {
	ID            int                              `json:"id_perusahaan" gorm:"column:id_perusahaan"`
	Nama          string                           `json:"nama_perusahaan" gorm:"column:nama_perusahaan"`
	Alamat        string                           `json:"alamat_perusahaan" gorm:"column:alamat_perusahaan"`
	TanggalMasuk  time.Time                        `json:"tanggal_masuk" gorm:"column:tanggal_masuk"`
	TanggalKeluar time.Time                        `json:"tanggal_keluar" gorm:"column:tanggal_keluar"`
	DaftarSiswa   []Models.NilaiSiswaPklPembimbing `json:"daftar_siswa"`
}

func GetNilaiPembimbing(c *gin.Context) {
	param_pembimbing := c.Param("id_pembimbing")

	param_industri := c.Param("id_industri")

	id_pembimbing, err := strconv.Atoi(param_pembimbing)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "id tidak valid",
		})
		return
	}
	id_industri, err := strconv.Atoi(param_industri)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"id":    id_industri,
			"error": "id tidak valid",
		})
		return
	}

	data_industri, err := Models.GetIndustri(id_industri)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	data_nilai, err := Models.GetNilaiByPemb(id_pembimbing, id_industri)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	nilai := NilaiPklPembimbing{
		ID:            data_industri.ID,
		Nama:          data_industri.Nama,
		Alamat:        data_industri.Alamat,
		TanggalMasuk:  data_industri.TanggalKeluar,
		TanggalKeluar: data_industri.TanggalKeluar,
		DaftarSiswa:   data_nilai,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": nilai,
	})

}

type NilaiPklFasilitator struct {
	ID            int                               `json:"id_perusahaan" gorm:"column:id_perusahaan"`
	Nama          string                            `json:"nama_perusahaan" gorm:"column:nama_perusahaan"`
	Alamat        string                            `json:"alamat_perusahaan" gorm:"column:alamat_perusahaan"`
	TanggalMasuk  time.Time                         `json:"tanggal_masuk" gorm:"column:tanggal_masuk"`
	TanggalKeluar time.Time                         `json:"tanggal_keluar" gorm:"column:tanggal_keluar"`
	DaftarSiswa   []Models.NilaiSiswaPklFasilitator `json:"daftar_siswa"`
}

func GetNilaiFasilitator(c *gin.Context) {
	param_fasil := c.Param("id_fasilitator")

	param_industri := c.Param("id_industri")

	id_fasil, err := strconv.Atoi(param_fasil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "id fasil tidak valid",
		})
		return
	}
	id_industri, err := strconv.Atoi(param_industri)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"id":    id_industri,
			"error": "id industri	 tidak valid",
		})
		return
	}

	data_industri, err := Models.GetIndustri(id_industri)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	data_nilai, err := Models.GetNilaiByFasil(id_fasil, id_industri)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	nilai := NilaiPklFasilitator{
		ID:            data_industri.ID,
		Nama:          data_industri.Nama,
		Alamat:        data_industri.Alamat,
		TanggalMasuk:  data_industri.TanggalKeluar,
		TanggalKeluar: data_industri.TanggalKeluar,
		DaftarSiswa:   data_nilai,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": nilai,
	})

}
