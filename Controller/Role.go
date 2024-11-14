package Controllers

import (
	"go-gin-mysql/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllRole(c *gin.Context) {

	roles, err := Models.GetRoles()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Failed to retrieve upcoming shows",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": roles,
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
		"massage": "Role berhasil ditambahkan",
	})

}

// type roleController struct {
// 	studentRepository repository.StudentRepository
// }

// func AddRole(c *gin.Context) {
// 	var roles Models.Role
// 	err := c.ShouldBindJSON(&roles)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	}
// }

// func GetRoles(c *gin.Context) {
// 	roles, err := Models.GetRoles()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data roles"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, roles)
// }

// type roleController struct {
// 	studentRepository repository.StudentRepository
// }

// func NewRoleController(studentRepository repository.StudentRepository) StudentService {
// 	return &studentService{studentRepository}
// }
