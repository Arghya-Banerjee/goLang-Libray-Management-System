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
	result := utils.DB.Raw("EXEC usp_GetAllBooks").Scan(&books)
	if result.Error != nil {
		// fmt.Println("Error executing stored procedure:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
	} else {
		// fmt.Println("Queried books successfully")
		c.JSON(http.StatusOK, books)
	}
}

// AddBook adds a new book to the database
func AddBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var newBook models.Book
	result := utils.DB.Raw(
		"EXEC usp_AddBook @Title = ?, @Author = ?, @Genre = ?, @Stock = ?, @Rating = ?;",
		book.Title, book.Author, book.Genre, book.Stock, book.Rating,
	).Scan(&newBook)

	if result.Error != nil {
		// fmt.Println("Error executing stored procedure:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
	} else {
		// fmt.Println("Added Book:", newBook)
		c.JSON(http.StatusOK, newBook)
	}
}

// UpdateBook updates the details of an existing book by ID
func UpdateBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var newBook models.Book
	result := utils.DB.Raw(
		"EXEC usp_UpdateBook @BookId = ?, @Title = ?, @Author = ?, @Genre = ?, @Stock = ?, @Rating = ?;",
		id, book.Title, book.Author, book.Genre, book.Stock, book.Rating,
	).Scan(&newBook)

	if result.Error != nil {
		// fmt.Println("Error executing stored procedure:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
	} else {
		// fmt.Println("Updated Book:", newBook)
		c.JSON(http.StatusOK, newBook)
	}
}

// DeleteBook deletes a book by ID
func DeleteBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := utils.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found!"})
		return
	}
	var msg string
	result := utils.DB.Raw(
		"EXEC usp_DeleteBook @BookId = ?;",
		id,
	).Scan(&msg)

	if result.Error != nil {
		// fmt.Println("Error executing stored procedure:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
	} else {
		// fmt.Println("Book Deleted Successfully")
		c.JSON(http.StatusOK, msg)
	}
}
