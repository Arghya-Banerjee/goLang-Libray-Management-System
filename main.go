package main

import (
	"fmt"
	"lms/utils" // Import your utils package
	"log"
)

func main() {
	// Initialize the database connection
	db := utils.InitDB()

	// Defer closing the connection until the main function finishes
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal("Error closing the database: ", err)
		} else {
			fmt.Println("Database connection closed.")
		}
	}()

	fmt.Println("Database connection successful!")
}
