package controllers

import (
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/bishalr0y/go-bookstore/pkg/config"
	"github.com/bishalr0y/go-bookstore/pkg/models"
	"github.com/gin-gonic/gin"
)

// var book models.Book

var db = config.ConnectDb()
var validate = validator.New(validator.WithRequiredStructEnabled())

func GetBooks(c *gin.Context) {
	var books []models.Book
	db.Find(&books)
	c.JSON(http.StatusOK, gin.H{
		"books": books,
	})

}

func CreateBook(c *gin.Context) {
	var book models.Book
	c.BindJSON(&book)

	if err := validate.Struct(book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.Create(&book)
	c.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

func GetBookById(c *gin.Context) {
	var book models.Book
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id query not provided",
		})
		return
	}
	db.First(&book, id)

	if book.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "book with such id doesn't exists",
			"note":  "book id starts from 1",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"book": book,
	})

}

func UpdateBook(c *gin.Context) {

}

func DeleteBook(c *gin.Context) {

}
