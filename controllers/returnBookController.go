package controllers

import (
	"lms/models"
	"lms/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReturnBook allows users to return a borrowed book
func ReturnBook(c *gin.Context) {
	var returnRequest struct {
		BookID int `json:"book_id"`
	}

	// Parse the book_id from the request body
	if err := c.ShouldBindJSON(&returnRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the user ID from the JWT token (set in middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Find the borrowed book record for this user and book
	var borrowedBook models.BorrowedBook
	if err := utils.DB.Where("user_id = ? AND book_id = ?", userID, returnRequest.BookID).First(&borrowedBook).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Borrowing record not found"})
		return
	}

	// Delete the borrowed book record (or update its return status, depending on your preference)
	utils.DB.Delete(&borrowedBook)

	// Increase the stock of the book
	var book models.Book
	if err := utils.DB.First(&book, returnRequest.BookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	book.Stock++
	utils.DB.Save(&book)

	c.JSON(http.StatusOK, gin.H{"message": "Book returned successfully"})
}
