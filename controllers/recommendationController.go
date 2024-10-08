package controllers

import (
	"lms/models"
	"lms/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RecommendBooks returns a list of top-rated books
func RecommendBooks(c *gin.Context) {
	var books []models.Book

	// Fetch books with the highest ratings (you can limit this to top 5 or 10)
	utils.DB.Order("rating DESC").Limit(5).Find(&books)

	c.JSON(http.StatusOK, books)
}
