package controllers

import (
	"fmt"
	"lms/models"
	"lms/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterUser handles the user registration process without password hashing
func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the password field is populated
	fmt.Println("Password from request:", user.Password)

	// Save the user to the database (including the plain text password)
	if err := utils.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully!"})
}
