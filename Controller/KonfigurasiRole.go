package Controllers

import (
	"go-gin-mysql/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteRolePegawai(c *gin.Context) {
	var id Models.IdRequest

	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := Models.DeleteRolePegawai(id.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "gagal menghapus role pada pegawai",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil menghapus role pada pegawai",
	})
}
