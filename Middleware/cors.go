package Middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware untuk mengaktifkan CORS
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		origin := c.Request.Header.Get("Origin")

		// Izinkan hanya domain tertentu
		if origin == "http://localhost:3000" || origin == "https://sipkl.smkpunegerijabar.sch.id" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
