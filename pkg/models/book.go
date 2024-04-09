package models

import (
	"github.com/bishalr0y/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Author      string `json:"author" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Publication string `json:"publication" validate:"required"`
}

func Init() {
	db = config.ConnectDb()

	var book Book

	db.AutoMigrate(&book)

}
