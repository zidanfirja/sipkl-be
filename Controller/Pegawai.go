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
	ID         int                   `json:"id"`
	IdPegawai  string                `json:"id_pegawai"`
	Nama       string                `json:"nama"`
	Email      string                `json:"email"`
	Password   string                `json:"password"`
	Aktif      bool                  `json:"aktif"`
	DaftarRole []Models.RespGetRoles `json:"daftar_role"`
	CreatedAt  time.Time             `json:"created_at"`
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
		// Buat daftar role dari KonfigurasiRoles
		var roles []Models.RespGetRoles
		for _, konfigurasiRole := range item.KonfigurasiRoles {
			roles = append(roles, Models.RespGetRoles{
				IDKonRole: konfigurasiRole.ID,
				IDRole:    konfigurasiRole.Role.ID,
				Nama:      konfigurasiRole.Role.Nama,
				Aktif:     konfigurasiRole.Role.Aktif,
				CreatedAt: konfigurasiRole.Role.CreatedAt,
			})
		}

		// Tambahkan pegawai ke respons
		response = append(response, RespPegawaiGetAll{
			ID:         item.ID,
			IdPegawai:  item.IdPegawai,
			Nama:       item.Nama,
			Email:      item.Email,
			Password:   item.Password,
			Aktif:      item.Aktif,
			DaftarRole: roles,
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

	bcryptPass := HashPassword(passwordValue.(string))

	pegawai.Payload["password"] = bcryptPass

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

	okPayloadAktif := payload["aktif"].(bool) // aktif adalah boolean
	if !okPayloadAktif {
		c.JSON(http.StatusBadRequest, gin.H{"error": "aktif is required"})
		return
	}

	if ok := Models.CekRolePegawai(int(id), int(idRole)); ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Pegawai dengan role tersebut sudah ada",
		})
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
	c.JSON(http.StatusOK, gin.H{
		"message": "role berhasil di assign",
	})

}
