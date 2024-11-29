package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/DeoMandupp/Library-Management/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBooks(c *gin.Context) {
	// Get the user ID from query parameters
	userID := c.Query("id")

	var books []models.Book
	var result *gorm.DB

	if userID != "" {
		fmt.Println("using ID")
		// If a user ID is provided, then fetching the specific book
		result = models.DB.Where("id = ?", userID).Find(&books)
	} else {
		fmt.Println("Not using ID")
		// If user ID is not provided, the fetching all the books
		result = models.DB.Find(&books)
	}

	// Check for database errors
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch books",
		})
		return
	}

	// Check if no books were found
	if len(books) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No books found",
		})
		return
	}

	// Return the books
	c.JSON(http.StatusOK, books)
}

// func GetBooks(c *gin.Context) {
// 	var books []models.Book

// 	res := models.DB.Find(&books)
// 	if res.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "Failed to fetch"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, books)
// }

func AddBooks(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Input",
		})
		return
	}

	// Add books to Database
	result := models.DB.Create(&book)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add book"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book added successfully", "book": book})
}

// Delete books from data base using ID

func DeleteBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	// Delete the book from the database
	result := models.DB.Delete(&models.Book{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

// Fetch a book by ID
func GetBookByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var book models.Book
	result := models.DB.First(&book, id) // Fetch the book using the primary key (ID)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch book"})
		}
		return
	}

	c.JSON(http.StatusOK, book)
}

// Modify book details by ID
func ModifyBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	// Fetch the book by ID
	var book models.Book
	if err := models.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Bind the updated data from the request body
	var updatedData models.Book
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Update only the provided fields
	if err := models.DB.Model(&book).Updates(updatedData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully", "book": book})
}
