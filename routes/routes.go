package routes

import (
	"lms/controllers"
	"lms/middleware" // Middleware is still needed for protected routes

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Public Routes (No Authentication Needed)
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)

	// Book Management (Public Routes)
	router.GET("/books", controllers.GetBooks)          // Get all books
	router.POST("/books", controllers.AddBook)          // Add a new book (public access)
	router.PUT("/books/:id", controllers.UpdateBook)    // Update a book by ID (public access)
	router.DELETE("/books/:id", controllers.DeleteBook) // Delete a book by ID (public access)

	// Recommendations (Public Routes)
	router.GET("/recommendations", controllers.RecommendBooks) // Get top-rated book recommendations

	// Protected Routes (Require Authentication for Borrowing/Returning Books and Ratings)
	protected := router.Group("/")
	protected.Use(middleware.JWTAuthMiddleware()) // Apply middleware only to protected routes

	// Borrowing and Returning books
	protected.POST("/borrow", controllers.BorrowBook) // Borrow a book (requires login)
	protected.POST("/return", controllers.ReturnBook) // Return a borrowed book (requires login)

	// Rating books
	protected.POST("/rate", controllers.RateBook) // Rate a book (requires login)
}
