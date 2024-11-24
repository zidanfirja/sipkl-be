package Middleware

import (
	"fmt"
	"go-gin-mysql/Models"
	"net/http"
	"os"
	"strings"

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

func CheckAuthToken() gin.HandlerFunc {

	return func(c *gin.Context) {
		// Ambil token dari header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
			c.Abort()
			return
		}

		// Hapus prefix "Bearer "
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse token dan verifikasi
		token, err := jwt.ParseWithClaims(tokenString, &Models.ClaimsUser{}, func(token *jwt.Token) (interface{}, error) {
			// Pastikan algoritma sesuai
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			SecKey := GetSecKey()
			var jwtKey = []byte(SecKey)
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Ambil data claims jika token valid
		if claims, ok := token.Claims.(*Models.ClaimsUser); ok {
			// Set data ke context untuk digunakan di handler berikutnya
			c.Set("payload", claims)

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Lanjut ke handler berikutnya
		c.Next()
	}

	// return func(c *gin.Context) {
	// 	authHeader := c.GetHeader("Authorization")
	// 	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
	// 		c.JSON(http.StatusUnauthorized, gin.H{
	// 			"error":   "data header tidak ditemukan",
	// 			"message": "anda perlu login terlebih dahulu",
	// 		})
	// 		c.Abort()
	// 		return
	// 	}

	// 	rawToken := strings.TrimPrefix(authHeader, "Bearer ")

	// 	token, err := jwt.ParseWithClaims(rawToken, &Models.ClaimsUser{}, func(t *jwt.Token) (interface{}, error) {
	// 		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
	// 			return nil, fmt.Errorf("mehtode signing tidak diketahui: %v", t.Header["alg"])
	// 		}
	// 		secretKey := GetSecKey()
	// 		return secretKey, nil
	// 	})

	// 	// cek token yang sudah di parse
	// 	if err != nil || !token.Valid {
	// 		c.JSON(http.StatusUnauthorized, gin.H{
	// 			"error":   err,
	// 			"message": "invalid token",
	// 		})
	// 		c.Abort()
	// 		return
	// 	}

	// 	// ini memasukan data ke set yang ada di gin.Context unutk handler atau middleware selanjutnya
	// 	// if claims, ok := token.Claims.(*Models.ClaimsUser); !ok {
	// 	// 	c.Set("nama_pegawai", claims.User.Nama)

	// 	// } else {
	// 	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
	// 	// 	c.Abort()
	// 	// 	return
	// 	// }

	// 	c.Next()

	// }

}
