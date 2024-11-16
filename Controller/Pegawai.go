package Controllers

import (
	"encoding/json"
	"go-gin-mysql/Models"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RespPegawaiGetAll struct {
	ID        int    `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	IdPegawai string `gorm:"type:varchar(100);not null" json:"id_pegawai"`
	Nama      string `gorm:"type:varchar(255);not null" json:"nama"`
	Email     string `gorm:"unique;type:varchar(100)" json:"email"`
	Password  string `json:"password" gorm:"type:varchar(255)"`
	Aktif     bool   `json:"aktif"`

	DaftarRole []Models.Role `json:"daftar_role" gorm:"-"`
	// Pembimbing       []DataSiswa        `gorm:"foreignKey:FKIdPembimbing;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	// Fasilitator      []DataSiswa        `gorm:"foreignKey:FKIdFasilitator;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
}

func GetAllPegawai(c *gin.Context) {

	data, err := Models.GetPegawai()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var response []RespPegawaiGetAll
	for _, item := range data {

		dataRole, err := Models.GetRoleByIdPegawai(item.ID)
		if err != nil {
			dataRole = nil
		}

		response = append(response, RespPegawaiGetAll{
			ID:         item.ID,
			IdPegawai:  item.IdPegawai,
			Nama:       item.Nama,
			Email:      item.Email,
			Password:   item.Password,
			Aktif:      item.Aktif,
			DaftarRole: dataRole,
			CreatedAt:  item.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": response,
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
	var pegawaiModel []Models.Pegawai

	rawData, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "gagal membaca data",
			"error":   err.Error(),
		})
		return
	}

	if err := json.Unmarshal(rawData, &pegawaiModel); err != nil {
		var pegawai Models.Pegawai
		if err := json.Unmarshal(rawData, &pegawai); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"massage": "format data tidak valid",
				"error":   err.Error(),
			})
			return
		}
		pegawai.Password = HashPassword(pegawai.Password)
		pegawaiModel = append(pegawaiModel, pegawai)
	}

	for _, pegawai := range pegawaiModel {

		pegawai.Password = HashPassword(pegawai.Password)
		errCreate := Models.CreatePegawai(&pegawai)
		if errCreate != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": errCreate.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data pegawai berhasil ditambahkan",
	})

}

func DeletePegawai(c *gin.Context) {
	var deleteReq Models.DeletePegawaiReq

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
		errDeletePegaawi := Models.DeletePegawai(id)
		if errDeletePegaawi != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   errDeletePegaawi.Error(),
				"message": "Gagal menghapus id " + strconv.Itoa(id),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil menghapus pegawai",
	})
}

func UpdatePegawai(c *gin.Context) {
	var pegawai Models.UpdatePegawaiReq

	errBindJson := c.ShouldBindJSON(&pegawai)
	if errBindJson != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errBindJson.Error(),
		})
		return
	}

	aktifValue, okAktif := pegawai.Payload["aktif"]
	passwordValue, okPass := pegawai.Payload["password"]

	// cek aktif dan pass ada  payload
	if okAktif && okPass {

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
			pegawai.Payload["aktif"] = aktif == 1
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "'aktif' harus berupa true, false, 1, atau 0",
			})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Kolom aktif dan password tidak boleh kososng",
		})
		return
	}

	passwordValue = HashPassword(passwordValue.(string))

	switch ids := pegawai.ID.(type) {
	case float64:
		err := Models.UpdateSinglePegawai(int(ids), pegawai.Payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Pegawai berhasil diupdate",
			"id":      int(ids),
		})
	case []interface{}:

		intIDs := make([]int, 0)

		for _, id := range ids {
			if idFloat, ok := id.(float64); ok {
				intIDs = append(intIDs, int(idFloat))
			}
		}

		err := Models.UpdateMultiplePegawai(intIDs, pegawai.Payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Pegawai berhasil diupdate",
			"ids":     intIDs,
		})
	}

}

func AssignRole(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var data map[string]interface{}

	if errUnmarshal := json.Unmarshal(body, &data); errUnmarshal != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errUnmarshal.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

	// Mengakses field "id"
	id, ok := data["id"].(float64) // id adalah angka (float64 setelah unmarshal)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	// Mengakses field "payload"
	payload, ok := data["payload"].(map[string]interface{}) // "payload" adalah objek
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload is required"})
		return
	}

	// Mengakses "id_role" dan "aktif" dalam payload
	idRole, ok := payload["id_role"].(float64) // id_role adalah angka (float64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id_role is required"})
		return
	}

	aktif, ok := payload["aktif"].(bool) // aktif adalah boolean
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "aktif is required"})
		return
	}

	idRoleInt := int(idRole)
	konfigurasiRole := Models.KonfigurasiRoles{
		FKIdPegawai: int(id),
		FKIdRole:    &idRoleInt,
	}

	errAddKonfigurasiRole := Models.AddKonfigurasiRole(&konfigurasiRole)
	if errAddKonfigurasiRole != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errAddKonfigurasiRole.Error(),
		})
		return
	}

	errUpdateAktifPegawai := Models.UpdateAktifPegawai(int(id), aktif)
	if errUpdateAktifPegawai != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errUpdateAktifPegawai.Error(),
		})
		return
	}

}
