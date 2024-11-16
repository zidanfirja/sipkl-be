package Database

import (
	"fmt"
	"log"

	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

func LoadDBConfig() DBConfig {

	if err := godotenv.Load(); err != nil {
		fmt.Println("Tidak ada file ENV")
	}

	return DBConfig{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
	}

}

var Database *gorm.DB

func ConnetDB() {
	dbConfig := LoadDBConfig()

	var err error

	// postgrest for prod

	// Ambil nilai-nilai dari variabel lingkungan
	dbUsername := dbConfig.Username
	dbPassword := dbConfig.Password
	dbHost := dbConfig.Host
	dbPort := dbConfig.Port
	dbName := dbConfig.DBName

	dsn := "host=" + dbHost + " user=" + dbUsername + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable"
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // Nonaktifkan pluralisasi nama tabel
		},
	})
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}
	log.Println("Berhasil terhubung ke database PostgreSQL")

	// mysql for development
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	dbConfig.Username,
	// 	dbConfig.Password,
	// 	dbConfig.Host,
	// 	dbConfig.Port,
	// 	dbConfig.DBName,
	// )

	// Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	NamingStrategy: schema.NamingStrategy{
	// 		SingularTable: true,
	// 	},
	// })

	// if err != nil {
	// 	panic("Failed to connect to database!")
	// }

}

func AutoMigrate(models ...interface{}) {
	if err := Database.AutoMigrate(models...); err != nil {
		panic("Failed to migrate database!")
	}
}
