package Controllers

import (
	"go-gin-mysql/Models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllRole(c *gin.Context) {

	roles, err := Models.GetRoles()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Failed to retrieve upcoming shows",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": roles,
	})
}

func CreateNewRole(c *gin.Context) {
	var role Models.Role

	errBindJson := c.ShouldBindJSON(&role)
	if errBindJson != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errBindJson.Error(),
		})
		return
	}

	err := Models.CreateRole(&role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"massage": "Role berhasil ditambahkan",
	})

}

func DeleteRole(c *gin.Context) {
	var deleteReq Models.DeleteRoleReq

	errDeleteReq := c.ShouldBindJSON(&deleteReq)
	if errDeleteReq != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   errDeleteReq.Error(),
			"massage": "Format request tidak valid",
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
					"massage": "Semua ID harus berupa angka",
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
		errDeleteRole := Models.DeleteRole(id)
		if errDeleteRole != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   errDeleteRole.Error(),
				"massage": "Gagal menghapus id " + strconv.Itoa(id),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"massage": "Berhasil menghapus role",
	})
}
