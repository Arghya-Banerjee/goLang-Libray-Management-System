package utils

import (
	"fmt"
	"lms/models" // Import your models
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {

	// Database connection details
	server := "localhost"
	port := 1433
	user := "sa"
	password := "qwaszx1234" 
	database := "library_db"

	// Connection string for MS SQL Server
	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		user, password, server, port, database)

	// Open the connection using GORM and SQL Server driver
	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	// Store the DB connection in the global variable for use in the application
	DB = db

	// Automatically migrate the models (creates tables if they don't exist)
	err = db.AutoMigrate(&models.Book{}, &models.User{}, &models.BorrowedBook{}, &models.Rating{})
	if err != nil {
		log.Fatal("Failed to migrate models: ", err)
	}

	fmt.Println("Database migration completed successfully!")
	return db
}
