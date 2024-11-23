package Controllers

import (
	"go-gin-mysql/Models"
	"net/http"
	"strconv"

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
