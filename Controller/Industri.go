package Controllers

import (
	"go-gin-mysql/Models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type IndustriResponse struct {
	ID        int       `json:"id"`
	Nama      string    `json:"nama"`
	Alamat    string    `json:"alamat"`
	Jurusan   string    `json:"jurusan"`
	CreatedAt time.Time `json:"created_at"`
}

func GetAllIndustri(c *gin.Context) {
	data, err := Models.GetIdustri()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengambil data",
			"error":   err.Error(),
		})
		return
	}

	var response []IndustriResponse
	for _, item := range data {
		response = append(response, IndustriResponse{
			ID:        item.ID,
			Nama:      item.Nama,
			Alamat:    item.Alamat,
			Jurusan:   item.Jurusan,
			CreatedAt: item.CreatedAt,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})

}

func CreateIndustri(c *gin.Context) {

	var industri Models.Industri

	err := c.ShouldBindJSON(&industri)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "data yang dimuat gagal",
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
		"message": "berhasil menambahkan ",
	})
}

func DeleteIndustri(c *gin.Context) {

	var deleteReq Models.DeleteIndustriReq

	err := c.ShouldBindJSON(&deleteReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "masukan data yang seusai",
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
					"message": "input data yang sesuai",
				})
				return
			}
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "data id tidak sesuai",
		})
		return
	}

	if len(ids) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "masukan setidaknya 1 industri",
		})
		return
	}

	for _, id := range ids {

		errDeleteIndustri := Models.DeleteIndustri(id)
		if errDeleteIndustri != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   errDeleteIndustri.Error(),
				"message": "Gagal menghapus id " + strconv.Itoa(id),
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil menghapus Industri",
	})

}

func UpdateIndustri(c *gin.Context) {
	var industriReq Models.UpdateIndustriReq

	errBindJson := c.ShouldBindJSON(&industriReq)
	if errBindJson != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error saat binding data",
			"error":   errBindJson.Error(),
		})
		return
	}

	jurusan, ok := industriReq.Payload["jurusan"]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Kolom jurusan tidak boleh kososng",
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
				"message": "Gagal update data",
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
					"message": "format data id salah",
				})
				return
			}
		}

		err := Models.UpdateMultipleIndustri(intIds, jurusanStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error saat update multiple",
				"error":   err.Error(),
			})
			return
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "data berhasil di update",
	})

}
