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
	DBUrl    string
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
		DBUrl:    os.Getenv("DATABASE_URL"),
	}

}

var Database *gorm.DB

func ConnetDB() {
	dbConfig := LoadDBConfig()

	var err error

	// postgrest for prod

	dsn := dbConfig.DBUrl + "?pgbouncer=true&connection_limit=1"

	// dsn := "host=" + dbHost + " user=" + dbUsername + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable" + " pg_stmtcache.mode=describe"
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // Nonaktifkan pluralisasi nama tabel
		},
		PrepareStmt: false, // Nonaktifkan prepared statement cache (untuk seeding)
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

func CloseDB() {
	db, err := Database.DB()
	if err != nil {
		log.Printf("Gagal mendapatkan objek *sql.DB: %v", err)
		return
	}

	if err := db.Close(); err != nil {
		log.Printf("Gagal menutup koneksi database: %v", err)
	} else {
		log.Println("Koneksi database berhasil ditutup")
	}
}

func AutoMigrate(models ...interface{}) {
	if err := Database.AutoMigrate(models...); err != nil {
		panic("Failed to migrate database!")
	}
}
