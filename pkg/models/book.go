package models

// create model (author, title, publication)
// connect to the db
// and create init function to auto migrate the database
import (
	"github.com/bishalr0y/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type Book struct {
	gorm.Model
	Author      string `json:"author"`
	Title       string `json:"title"`
	Publication string `json:"publication"`
}

func Init() {
	db = config.ConnectDb()

	var book Book

	db.AutoMigrate(&book)

}
