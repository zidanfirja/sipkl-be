package main

import (
	"go-gin-mysql/Database"
	"go-gin-mysql/Models"
	"go-gin-mysql/Routers"
	"os"
)

func main() {
	Database.ConnetDB()

	Database.AutoMigrate(&Models.Role{}, &Models.Pegawai{}, &Models.KonfigurasiRoles{}, &Models.Industri{}, &Models.DataSiswa{})
	// Database.DB.Migrator().DropTable(&Models.Role{}, &Models.Pegawai{}, &Models.KonfigurasiRoles{}, &Models.Industri{}, &Models.DataSiswa{})

	// roleModel := Models.RoleModel(db)
	// userService := Controller.RoleCOntroller(roleModel)
	//
	// Mengatur router Gin

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := Routers.SetupRouter()
	router.Run(":" + port)

}
