package Controllers

import (
	"fmt"
	"go-gin-mysql/Models"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func GetSecKey() string {

	if err := godotenv.Load(); err != nil {
		fmt.Println("Tidak ada file ENV")
		return ""
	}
	return os.Getenv("SECRET_KEY")

}

func Login(c *gin.Context) {

	// cek credential login
	var cred Models.Credential

	errBindJson := c.ShouldBindJSON(&cred)
	if errBindJson != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   errBindJson.Error(),
			"message": "data login tidak valid",
		})
		return
	}

	user, err := Models.AuthenticateUser(&cred)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "error di AuthenticateUser",
		})
		return
	}

	// create token
	stringTkn, err := CreateJwt(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Gagal membuat create token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "berhasil login",
		"jwt_token": stringTkn,
	})

}

func CreateJwt(user *Models.Pegawai) (string, error) {

	var current_role Models.DataRole
	var daftar_role []Models.DataRole

	// cari role pegawai
	role_pegawai, err := Models.GetRoleByIdPegawai(user.ID)
	if err != nil {
		return "", fmt.Errorf("anda belum memiliki role. error: %s", err.Error())
	}

	// masukan ke daftar role
	for _, data := range role_pegawai {
		newDataRole := Models.DataRole{
			IDRole:   data.IDRole,
			NamaRole: data.Nama,
		}

		daftar_role = append(daftar_role, newDataRole)
	}

	//masukan current role
	current_role.IDRole = daftar_role[0].IDRole
	current_role.NamaRole = daftar_role[0].NamaRole

	claims := Models.ClaimsUser{
		User: Models.Userdata{
			ID:        user.ID,
			IdPegawai: user.IdPegawai,
			Nama:      user.Nama,
			Email:     user.Email,
		},
		CurrentRole: current_role,
		DaftarRole:  daftar_role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 2)),
			Issuer:    "sipkl-smkpu",
		},
	}

	SecKey := GetSecKey()
	var jwtKey = []byte(SecKey)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func PayloadLogin(c *gin.Context) {

	payload, ok := c.Get("payload")
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "data tidak ditemukan",
		})
		return
	}
	data, ok := payload.(Models.ClaimsUser)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "gagal parse data",
		})
		return
	}

	id_role := c.Query("id_role")
	if id_role != "" {
		idInt, err := strconv.Atoi(id_role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "id role tidak bisa di kenversi ke int",
			})
			return
		}

		for _, dataRole := range data.DaftarRole {
			if dataRole.IDRole == idInt {
				data.CurrentRole.IDRole = idInt
				data.CurrentRole.NamaRole = dataRole.NamaRole
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}
