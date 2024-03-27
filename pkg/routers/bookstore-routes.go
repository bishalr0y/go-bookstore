package router

import (
	"github.com/bishalr0y/go-bookstore/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterBookstoreRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/book", controllers.GetBooks)
	router.POST("/book", controllers.CreateBook)
	router.GET("/book/:id", controllers.GetBookById)
	router.PUT("/book/:id", controllers.UpdateBook)
	router.DELETE("/book/:id", controllers.DeleteBook)

	return router
}
