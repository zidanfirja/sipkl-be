package main

import (
	"fmt"
	"go-gin-mysql/Database"
	"go-gin-mysql/Routers"
)

func main() {
	Database.ConnetDB()
	// Menjalankan migrasi otomatis
	// Database.AutoMigrate(&Models.DataSiswa{}, &Models.Role{}, &Models.Industri{}, &Models.Pegawai{}, &Models.KonfigurasiRoles{})

	// Database.AutoMigrate(&Models.Role{}, &Models.Pegawai{}, &Models.KonfigurasiRoles{}, &Models.Industri{}, &Models.DataSiswa{})
	// Database.DB.Migrator().DropTable(&Models.Role{}, &Models.Pegawai{}, &Models.KonfigurasiRoles{}, &Models.Industri{}, &Models.DataSiswa{})

	// roleModel := Models.RoleModel(db)
	// userService := Controller.RoleCOntroller(roleModel)

	// Mengatur router Gin
	router := Routers.SetupRouter()
	fmt.Println("Run at localhost:8080")
	router.Run(":8080")

}
