package Controllers

import (
	"go-gin-mysql/Models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetAllPegawai(c *gin.Context) {

	data, err := Models.GetPegawai()
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

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func CreatePegawai(c *gin.Context) {
	var pegawaiModel Models.Pegawai

	if errBindJson := c.ShouldBindJSON(&pegawaiModel); errBindJson != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errBindJson.Error(),
		})
	}

	pegawaiModel.Password = HashPassword(pegawaiModel.Password)

	if err := Models.CreatePegawai(&pegawaiModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data pegawai berhasil ditambahkan",
	})

}

func DeletePegawai(c *gin.Context) {
	var deleteReq Models.DeletePegawaiReq

	errDeleteReq := c.ShouldBindJSON(&deleteReq)
	if errDeleteReq != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   errDeleteReq.Error(),
			"message": "Format request tidak valid",
		})
		return
	}
	var ids []int

	switch dataId := deleteReq.ID.(type) {
	case float64:
		ids = append(ids, int(dataId))
	case []interface{}:
		for _, item := range dataId {
			if id, ok := item.(float64); ok {
				ids = append(ids, int(id))
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Semua ID harus berupa angka",
				})
				return
			}
		}
	}

	if len(ids) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Harap masukkan setidaknya satu ID untuk dihapus",
		})
		return
	}

	for _, id := range ids {
		errDeletePegaawi := Models.DeletePegawai(id)
		if errDeletePegaawi != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   errDeletePegaawi.Error(),
				"message": "Gagal menghapus id " + strconv.Itoa(id),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil menghapus pegawai",
	})
}
