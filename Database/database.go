package Database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
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

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)

	Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("Failed to connect to database!")
	}

}

func AutoMigrate(models ...interface{}) {
	if err := Database.AutoMigrate(models...); err != nil {
		panic("Failed to migrate database!")
	}
}
