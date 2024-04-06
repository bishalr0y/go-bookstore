package main

import (
	"fmt"

	"github.com/bishalr0y/go-bookstore/pkg/config"
	"github.com/bishalr0y/go-bookstore/pkg/models"
	"github.com/bishalr0y/go-bookstore/pkg/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to Go Bookstore!")
	r := gin.Default()
	db := config.ConnectDb()

	models.Init()

	// do something with the database
	fmt.Println(db)

	routers.RegisterBookstoreRoutes(r)

	r.Run(":9010")

}

//! Calling the router and running the server on the port 9010
