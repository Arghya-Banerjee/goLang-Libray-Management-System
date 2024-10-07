package main

import (
	"lms/routes"
	"lms/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database and migrate models
	utils.InitDB()

	// Initialize the Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router)

	// Start the server on port 8080
	router.Run(":8080")
}
