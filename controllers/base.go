package controllers

import (
	config "gin-crud/configs"
	"gin-crud/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Serve struct {
	DB *gorm.DB
}

func (serve *Serve) Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var variables config.Config
	variables.DB = os.Getenv("DB_NAME")
	variables.Host = os.Getenv("DB_HOST")
	variables.Port = os.Getenv("DB_PORT")
	variables.Username = os.Getenv("DB_USER")
	variables.Passwrod = os.Getenv("DB_PASSWORD")
	serve.DB = config.InitDatabase(variables)
	Migration(serve.DB)
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Category{}, &models.Author{}, &models.Book{}, &models.BAndR{})
}
