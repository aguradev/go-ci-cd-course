package configs

import (
	"fmt"
	"log"
	"os"
	"praktikum/models"
	books "praktikum/models/Books"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadDatabase() {

	var DBException error

	configuration := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	DB, DBException = gorm.Open(mysql.Open(configuration), &gorm.Config{})

	if DBException != nil {
		log.Println(DBException)
		panic("Failed connect to database")
	}

	ConfigMigrations()
}

func ConfigMigrations() {
	DB.AutoMigrate(&models.Users{}, &books.Books{}, &models.Blogs{})
}
