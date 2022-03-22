package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	controllers "Gin/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/books", controllers.GetBook)
	router.POST("/books", controllers.AddBook)
	router.PUT("/books", controllers.UpdateBook)
	router.DELETE("/books", controllers.DeleteBook) //delete dari param

	router.Run("localhost:8080")
}
