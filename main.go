package main

import (
	"go-gin-mysql/Database"
	"go-gin-mysql/Models"
	"go-gin-mysql/Routers"
	"go-gin-mysql/Seed"
	"os"
)

func main() {
	Database.ConnetDB()

	// migrate database
	Database.AutoMigrate(&Models.Role{}, &Models.Pegawai{}, &Models.KonfigurasiRoles{}, &Models.Industri{}, &Models.DataSiswa{})
	// Database.DB.Migrator().DropTable(&Models.Role{}, &Models.Pegawai{}, &Models.KonfigurasiRoles{}, &Models.Industri{}, &Models.DataSiswa{})

	// seed data
	Seed.SeedRole()
	Seed.SeedIndustri()
	Seed.SeedPegawai()
	Seed.SeedKonfigurasiRoles()
	Seed.SeedDataSiswa()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := Routers.SetupRouter()
	router.Run(":" + port)

}
