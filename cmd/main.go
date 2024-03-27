package main

import (
	"fmt"

	"github.com/bishalr0y/go-bookstore/pkg/config"
)

func main() {
	fmt.Println("Welcome to Go Bookstore")
	db := config.ConnectDb()
	fmt.Println(db)
}
