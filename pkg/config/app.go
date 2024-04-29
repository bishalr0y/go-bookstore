package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {

	dsn := "host=postgres-bookstore user=postgres password=postgres dbname=bookstore port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	// dsn := "host=localhost user=postgres password=postgres dbname=bookstore port=5432 sslmode=disable TimeZone=Asia/Kolkata"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Error connecting to the database!")
	}
	return db

}
