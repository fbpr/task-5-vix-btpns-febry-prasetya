package router

import (
	"github.com/fbpr/task-5-vix-btpns-febry-prasetya/controllers"
	"github.com/fbpr/task-5-vix-btpns-febry-prasetya/middlewares"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controllers.CreateUser)
		userRouter.POST("/login", controllers.Login)
		userRouter.PUT("/:userId", controllers.UpdateUserById)
		userRouter.DELETE("/:userId", controllers.DeleteUserById)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middlewares.Auth())

		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.GetPhoto)
		photoRouter.PUT("/:photoId", controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", controllers.DeletePhoto)
	}

	return router
}