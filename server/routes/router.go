package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hjunor/api-rest-golang.git/controllers"
	"github.com/hjunor/api-rest-golang.git/server/middlewares"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/api/v1")
	{
		users := main.Group("/user")
		{
			users.POST("/", controllers.CreateUser)
			users.POST("/login", controllers.Login)
		}
		books := main.Group("/books", middlewares.Auth())
		{
			books.GET("/:id", controllers.ShowBook)
			books.GET("/", controllers.ShowBooks)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}
	return router
}
