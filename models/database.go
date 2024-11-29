package models

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	ID         int    `json:"ID" gorm:"primaryKey"` // extra field is added
	AuthorName string `json:"AuthorName" gorm:"column:AuthorName; not null"`
	BookName   string `json:"BookName" gorm:"column:BookName; not null"`
	ISBN       int    `json:"ISBN" gorm:"column:ISBN; not null"`
}

var DB *gorm.DB

//initialize the connection to the database

func ConnectToDB() {
	var err error

	// Load env file
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get database url from environment
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to Database")
	}

	//Migrate the book model
	// This automagrator is adding extra field in the data base
	err = DB.AutoMigrate(&Book{})
	if err != nil {
		log.Fatalf("failed to migrate %v", err)
	}
	log.Println("database connected and migrated ")

}
