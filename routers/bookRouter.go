package routers

import (
	"challenge-3-chapter-2/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	routers := gin.Default()

	routers.GET("/books", controllers.GetBookAll)

	routers.GET("/books/:bookID", controllers.GetBookId)

	routers.POST("/books", controllers.CreateBook)

	routers.PUT("/books/:bookID", controllers.UpdateBook)

	routers.DELETE("/books/:bookID", controllers.DeleteBook)

	return routers
}
