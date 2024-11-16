package Controllers

import (
	"go-gin-mysql/Models"
	"log"
	"net/http"

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
