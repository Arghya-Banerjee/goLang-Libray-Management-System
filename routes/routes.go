package routes

import(
	"github.com/gin-gonic/gin"
	"lms/controllers"
)

func SetupRoutes(router *gin.Engine){

	router.GET("/books", controllers.GetBooks)
	router.POST("/books", controllers.AddBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

}