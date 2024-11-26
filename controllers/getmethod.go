package controllers

import (
	"net/http"

	"github.com/DeoMandupp/Library-Management/models"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book

	res := models.DB.Find(&books)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch"})
		return
	}
	c.JSON(http.StatusOK, books)
}
