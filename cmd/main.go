package main

import (
	"fmt"

	"github.com/bishalr0y/go-bookstore/pkg/models"
	"github.com/bishalr0y/go-bookstore/pkg/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to Go Bookstore!")
	r := gin.Default()

	models.Init()
	routers.RegisterBookstoreRoutes(r)

	r.Run(":9010")
}
