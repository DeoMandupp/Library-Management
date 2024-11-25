package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type books struct {
	ID         int    `json:"ID" gorm:"primaryKey"`
	AuthorName string `json:"AuthorName"`
	BookName   string `json:"BookName"`
	ISBN       int    `json:"ISBN"`
}

var db *gorm.DB

func ConnectToDB() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DB_URL")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to Database")
	}

	//migrate the book model
	err = db.AutoMigrate(&books{})
	if err != nil {
		log.Fatalf("failed to migrate %v", err)
	}
	log.Println("database connected and migrated ")

}

func getBooks(c *gin.Context) {
	var books []books
	res := db.Find(&books)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch"})
		return
	}
	c.JSON(http.StatusOK, books)
}

func main() {

	ConnectToDB()

	r := gin.Default() //gin router

	r.GET("/books", getBooks)
	r.Run(":8080")

}
