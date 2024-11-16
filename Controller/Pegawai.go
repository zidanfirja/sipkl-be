package Controllers

import (
	"go-gin-mysql/Models"
	"net/http"

	"github.com/gin-gonic/gin"
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
