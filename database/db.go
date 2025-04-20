package database

import (
	"fmt"
	"log"

	"github.com/GAURAV/BookApiTask/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	pass   = "root"
	dbName = "Book_Mangement_System"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbName, pass)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Assign to global DB
	DB = db

	// Auto migrate your models
	err = DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Book{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database connected and migrated successfully ✅")
}
