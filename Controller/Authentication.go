package Controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-gin-mysql/Models"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  os.Getenv("RedirectURL"),
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

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

func LoginOAuth(c *gin.Context) {
	state := "smkpu-negerijabarbandung" // Ganti dengan generator state yang aman
	url := googleOauthConfig.AuthCodeURL(state)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func Callback(c *gin.Context) {
	state := c.Query("state")
	if state != "smkpu-negerijabarbandung" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state"})
		return
	}

	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Code not found"})
		return
	}

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
		return
	}

	// ini mengambil data user dari google , nama , email,foto dll
	client := googleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}
	defer resp.Body.Close()

	// ini untuk parse data dari google diatas ke struct yang ada
	var user Models.UserInfoOAuth
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user info"})
		return
	}

	//masukan data ke credential
	var cred Models.Credential
	cred.Email = user.Email

	//cek ada tidaknya user dengan models.AuthenticateUserCekEmail()
	dataUser, err := Models.AuthenticateUserCekEmail(&cred)
	if err != nil {
		// jika tidak ada
		// -- redirect ke halaman sebelumnya
		urlFailedLogin := os.Getenv("URL_PAGE_FAILED_LOGIN")
		c.Redirect(http.StatusFound, urlFailedLogin)
	}

	// jika ada :
	// -- create jwt
	stringTkn, err := CreateJwt(dataUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Gagal membuat create token",
		})
		return
	}
	// -- redirect ke halamn selanjutnya
	urlSuccesLogin := fmt.Sprintf(os.Getenv("URL_PAGE_SUCCESS_LOGIN")+"%s", stringTkn)
	c.Redirect(http.StatusFound, urlSuccesLogin)

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
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
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

type NewDataPayloadToken struct {
	Data Models.ClaimsUser `json:"data"`
}

type NewDataPayload struct {
	User        Models.Userdata   `json:"userdata"`
	CurrentRole Models.DataRole   `json:"current_role"`
	DaftarRole  []Models.DataRole `json:"daftar_role"`
	jwt.RegisteredClaims
}

func PayloadLogin(c *gin.Context) {

	payload, ok := c.Get("payload")
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "data tidak ditemukan",
		})
		return
	}

	payloadMap, ok := payload.(*Models.ClaimsUser)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "format payload tidak valid",
		})
		return
	}

	var newDataPayload Models.ClaimsUser
	id_role := c.Query("id_role")

	if id_role != "" {
		idInt, err := strconv.Atoi(id_role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "id role tidak bisa di kenversi ke int",
			})
			return
		}

		for _, dataRole := range payloadMap.DaftarRole {
			if dataRole.IDRole == idInt {
				newDataPayload = Models.ClaimsUser{
					User: payloadMap.User,
					CurrentRole: Models.DataRole{
						IDRole:   idInt,
						NamaRole: dataRole.NamaRole,
					},
					DaftarRole: payloadMap.DaftarRole,
					RegisteredClaims: jwt.RegisteredClaims{
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 2)),
						Issuer:    "sipkl-smkpu",
					},
				}
				break
			}
		}
		payloadMap = &newDataPayload

	}

	//membuat toke baru / refresh token
	SecKey := GetSecKey()
	var jwtKey = []byte(SecKey)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payloadMap)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "gagal",
			"error":   "gagal membuat refresh token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "sukses",
		"data":      payloadMap,
		"jwt_token": tokenString,
	})

}
