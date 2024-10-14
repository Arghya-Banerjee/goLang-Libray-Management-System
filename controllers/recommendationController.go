package controllers

import (
	"fmt"
	"lms/models"
	"lms/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RecommendBooks returns a list of top-rated books
func RecommendBooks(c *gin.Context) {
	var books []models.Book

	// Fetch books with the highest ratings (you can limit this to top 5 or 10)
	result := utils.DB.Raw(
		"EXEC usp_RecommendBooks;",
	).Scan(&books)

	if result.Error != nil {
		fmt.Println("Error executing stored procedure: ", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
	} else {
		fmt.Println("Recommended books successfully")
		c.JSON(http.StatusOK, books)
	}
}
