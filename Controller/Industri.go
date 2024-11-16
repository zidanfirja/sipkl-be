package Controllers

import (
	"go-gin-mysql/Models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllIndustri(c *gin.Context) {
	data, err := Models.GetIdustri()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Gagal mengambil data",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}

func CreateIndustri(c *gin.Context) {

	var industri Models.Industri

	err := c.ShouldBindJSON(&industri)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"massage": "data yang dimuat gagal",
			"error":   err.Error(),
		})
		return
	}

	errCreate := Models.CreateIndustri(&industri)
	if errCreate != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errCreate.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"massage": "berhasil menambahkan ",
	})
}

func DeleteIndustri(c *gin.Context) {

	var deleteReq Models.DeleteIndustriReq

	err := c.ShouldBindJSON(&deleteReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "masukan data yang seusai",
			"error":   err.Error(),
		})
		return
	}

	var ids []int

	// [2,3,4,5,"six"]

	switch dataId := deleteReq.ID.(type) {
	case float64:
		ids = append(ids, int(dataId))
	case []interface{}:
		for _, item := range dataId {
			if id, ok := item.(float64); ok {
				ids = append(ids, int(id))
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"massage": "input data yang sesuai",
				})
				return
			}
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "data id tidak sesuai",
		})
		return
	}

	if len(ids) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "masukan setidaknya 1 industri",
		})
		return
	}

	for _, id := range ids {

		errDeleteIndustri := Models.DeleteIndustri(id)
		if errDeleteIndustri != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   errDeleteIndustri.Error(),
				"massage": "Gagal menghapus id " + strconv.Itoa(id),
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"massage": "Berhasil menghapus Industri",
	})

}

func UpdateIndustri(c *gin.Context) {
	var industriReq Models.UpdateIndustriReq

	errBindJson := c.ShouldBindJSON(&industriReq)
	if errBindJson != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "error saat binding data",
			"error":   errBindJson.Error(),
		})
		return
	}

	jurusan, ok := industriReq.Payload["jurusan"]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "Kolom jurusan tidak boleh kososng",
		})
		return
	}

	jurusanStr, okCheckStr := jurusan.(string)
	if !okCheckStr || jurusanStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Field 'jurusan' harus string.",
		})
		return
	}

	switch dataId := industriReq.ID.(type) {
	case float64:
		err := Models.UpdateSingleIndustri(int(dataId), industriReq.Payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"massage": "Gagal update data",
				"error":   err.Error(),
			})
			return
		}
	case []interface{}:
		var intIds []int
		for _, item := range dataId {
			if idFloat, ok := item.(float64); ok {
				intIds = append(intIds, int(idFloat))
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"massage": "format data id salah",
				})
				return
			}
		}

		err := Models.UpdateMultipleIndustri(intIds, jurusanStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"massage": "error saat update multiple",
				"error":   err.Error(),
			})
			return
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"massage": "data berhasil di update",
	})

}
