package models

import (
	"fmt"
	"log"

	schemasql "github.com/enrichoalkalas01/learn-go-fiber.git/models/schema-sql"
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

	if !TableExists("categories") {
		if err := database.AutoMigrate(&schemasql.Category{}); err != nil {
			log.Fatal("Failed to migrate Category:", err)
		}
	}

	if !TableExists("products") {
		if err := database.AutoMigrate(&schemasql.Product{}); err != nil {
			log.Fatal("Failed to migrate Product:", err)
		}
	}
}

// Check Table Exist For Migrations
func TableExists(tableName string) bool {
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

// Function for close connection from database
func PGCloseConnection() {
	SQLDB, err := database.DB()
	if err != nil {
		log.Fatal("Failed to get SQL DB instance : ", err)
		return
	}

	if err := SQLDB.Close(); err != nil {
		log.Fatal("Failed to cose database connection : ", err)
	} else {
		log.Println("Database connection closed successfully")
	}
}
