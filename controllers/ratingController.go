package controllers

import (
	"lms/models"
	"lms/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RateBook allows a user to rate a book and updates the book's average rating
func RateBook(c *gin.Context) {
	var rating models.Rating

	// Bind the JSON input to the Rating model
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the user ID from the JWT token (set by middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Set the user ID in the rating
	rating.UserID = userID.(int)

	// Save the rating in the database
	if err := utils.DB.Create(&rating).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save rating"})
		return
	}

	// Calculate the new average rating for the book as a float
	var avgRating float64
	utils.DB.Table("ratings").Where("book_id = ?", rating.BookID).Select("AVG(CAST(rating AS FLOAT))").Scan(&avgRating)

	// Update the book's average rating in the books table
	if err := utils.DB.Model(&models.Book{}).Where("id = ?", rating.BookID).Update("rating", avgRating).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book rating"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Rating submitted successfully!"})
}
