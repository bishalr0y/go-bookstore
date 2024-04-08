package routers

import (
	"github.com/bishalr0y/go-bookstore/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterBookstoreRoutes(router *gin.Engine) {

	router.GET("/book", controllers.GetBooks)
	router.POST("/book", controllers.CreateBook)
	router.GET("/book/:id", controllers.GetBookById)
	router.PUT("/book/:id", controllers.UpdateBook)
	router.DELETE("/book/:id", controllers.DeleteBook)
}
