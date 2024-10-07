package routes

import (
	"lms/controllers"
	"lms/middleware" // Import the JWT middleware

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Public Routes
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)

	// Protected Routes (require JWT token)
	protected := router.Group("/")
	protected.Use(middleware.JWTAuthMiddleware())
	protected.POST("/borrow", controllers.BorrowBook)
	protected.POST("/return", controllers.ReturnBook)

	// Book Routes (add these as needed)
	router.GET("/books", controllers.GetBooks)
	router.POST("/books", controllers.AddBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)
}
