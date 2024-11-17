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
