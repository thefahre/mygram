package database

import (
	"fmt"
	"log"
	"mygram/models"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = os.Getenv("DB_HOST")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbport   = os.Getenv("DB_PORT")
	dbname   = os.Getenv("DB_NAME")
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbport)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	db.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.Socialmedia{})
}

func GetDB() *gorm.DB {
	return db
}
