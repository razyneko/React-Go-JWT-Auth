package database

import (
	"github.com/razyneko/React-Go-JWT-Auth/models"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){
	// Database connection string
	dsn := "root:password@tcp(127.0.0.1:3306)/moviesdb?charset=utf8mb4&parseTime=True&loc=Local"

    // Try to open a connection to the database
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    // Check if an error occurred
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err) 
	}
	DB = connection
    // If the connection is successful, log the db connection status
	log.Println("Successfully connected to the database")
	connection.AutoMigrate(&models.User{})
}