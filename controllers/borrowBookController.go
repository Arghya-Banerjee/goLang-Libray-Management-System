package controllers

import (
	"lms/models"
	"lms/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// BorrowBook allows authenticated users to borrow a book
func BorrowBook(c *gin.Context) {
	var borrowRequest struct {
		BookID int `json:"book_id"`
	}

	// Parse the book_id from the request body
	if err := c.ShouldBindJSON(&borrowRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the user ID from the JWT token (set in middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Find the book the user wants to borrow
	var book models.Book
	if err := utils.DB.First(&book, borrowRequest.BookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Check if the book is available in stock
	if book.Stock <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book out of stock"})
		return
	}

	// Decrease the stock of the book
	book.Stock--
	utils.DB.Save(&book)

	// Record the borrowing in the BorrowedBook table
	borrowedBook := models.BorrowedBook{
		UserID:     userID.(int), // Convert to int
		BookID:     book.ID,
		BorrowedAt: time.Now(),
		DueDate:    time.Now().AddDate(0, 1, 0), // Due date is 1 month from now
	}
	utils.DB.Create(&borrowedBook)

	c.JSON(http.StatusOK, gin.H{"message": "Book borrowed successfully", "due_date": borrowedBook.DueDate})
}
