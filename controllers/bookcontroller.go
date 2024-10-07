package controllers

import (
	"lms/models"
	"lms/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetBooks retrieves all books from the database
func GetBooks(c *gin.Context) {
	var books []models.Book
	utils.DB.Find(&books) // Retrieve all books
	c.JSON(http.StatusOK, books)
}

// AddBook adds a new book to the database
func AddBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	utils.DB.Create(&book) // Insert the new book into the database
	c.JSON(http.StatusOK, book)
}

// UpdateBook updates the details of an existing book by ID
func UpdateBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := utils.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found!"})
		return
	}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	utils.DB.Save(&book) // Update the book details in the database
	c.JSON(http.StatusOK, book)
}

// DeleteBook deletes a book by ID
func DeleteBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := utils.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found!"})
		return
	}
	utils.DB.Delete(&book) // Delete the book from the database
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully!"})
}
