package models

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

// Create Connection
func PGConnection() {
	var err error

	dsn := "host=localhost user=postgres password=1sampai10! dbname=learn-go-fiber port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database pg : ", err)
	}

	if !tableExists("categories") {
		if err := database.AutoMigrate(&Category{}); err != nil {
			log.Fatal("Failed to migrate Category:", err)
		}
	}

	if !tableExists("products") {
		if err := database.AutoMigrate(&Product{}); err != nil {
			log.Fatal("Failed to migrate Product:", err)
		}
	}
}

// Check Table Exist For Migrations
func tableExists(tableName string) bool {
	var count int64
	err := database.Raw(fmt.Sprintf("SELECT COUNT(*) FROM information_schema.tables WHERE table_name = '%s'", tableName)).Scan(&count).Error
	if err != nil {
		log.Printf("Failed to check table %s existence: %v", tableName, err)
		return false
	}
	return count > 0
}

// Function to useable use database
func PGDatabase() *gorm.DB {
	return database
}
