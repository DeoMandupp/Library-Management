package main

import (
	"github.com/DeoMandupp/Library-Management/controllers"
	"github.com/DeoMandupp/Library-Management/models"
	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectToDB()

	r := gin.Default() //gin router

	r.GET("/books", controllers.GetBooks)
	r.Run(":8080")

}
