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
			"message": "Failed to retrieve upcoming shows",
			"error":   err.Error(),
		})
		return
	}

	var response []Models.RespGetRoles
	for _, item := range roles {
		response = append(response, Models.RespGetRoles{
			ID:        item.ID,
			Nama:      item.Nama,
			Aktif:     item.Aktif,
			CreatedAt: item.CreatedAt,
		})

	}

	c.JSON(http.StatusOK, gin.H{
		"data": response,
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
		"message": "Role berhasil ditambahkan",
	})

}

func DeleteRole(c *gin.Context) {
	var deleteReq Models.DeleteRoleReq

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
		errDeleteRole := Models.DeleteRole(id)
		if errDeleteRole != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   errDeleteRole.Error(),
				"message": "Gagal menghapus id " + strconv.Itoa(id),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil menghapus role",
	})
}

func UpdateRole(c *gin.Context) {
	var role Models.UpdateRoleReq

	errBindJson := c.ShouldBindJSON(&role)
	if errBindJson != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errBindJson.Error(),
		})
		return
	}

	// cek 'aktif' ada  payload
	if aktifValue, ok := role.Payload["aktif"]; ok {
		switch aktif := aktifValue.(type) {
		case bool:
			// dikosongkan karena jika ini bool juga betul
		case float64:

			if aktif != 0 && aktif != 1 {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "'aktif' harus berupa true, false, 1, atau 0",
				})
				return
			}
			// Konversi angka ke boolean
			role.Payload["aktif"] = aktif == 1
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "'aktif' harus berupa true, false, 1, atau 0",
			})
			return
		}
	}

	switch ids := role.ID.(type) {
	case float64:
		err := Models.UpdateSingleRole(int(ids), role.Payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"id":    ids,
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Role berhasil diupdate",
			"id":      int(ids),
		})
	case []interface{}:

		intIDs := make([]int, 0)

		for _, id := range ids {
			if idFloat, ok := id.(float64); ok {
				intIDs = append(intIDs, int(idFloat))
			}
		}

		err := Models.UpdateMultipleRoles(intIDs, role.Payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Role berhasil diupdate",
			"ids":     intIDs,
		})
	}

}
