package models

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDatabase() {
	var err error

	DB, error := gorm.open(sqlite.Open("test.db"), &gorm.Config{})
	if err := nil {
		log.Fatal("Failed to connect to database", err)
	}

	DB.AutoMigrate(&models.Product{})
}