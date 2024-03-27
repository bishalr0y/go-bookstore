package main

import (
	"fmt"

	"github.com/bishalr0y/go-bookstore/pkg/config"
)

func main() {
	fmt.Println("Welcome to Go Bookstore")
	db, err := config.ConnectDb()

	if err != nil {
		panic("Error connecting to the database!")
	}
	// do something with the database
	fmt.Println(db)
}
