package Controllers

import (
	"go-gin-mysql/Database"
	"go-gin-mysql/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []Models.DataSiswa
	Database.Database.Find(&users)
	c.JSON(http.StatusOK, users)
}
