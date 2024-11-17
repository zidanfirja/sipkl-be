package Controllers

import (
	"go-gin-mysql/Models"
	"net/http"

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
	var dataSiswa Models.DataSiswa

	if err := c.ShouldBindJSON(&dataSiswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "input data yang valid",
			"error":   err.Error(),
		})
		return
	}
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
