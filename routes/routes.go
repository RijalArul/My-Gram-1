package routes

import (
	"my-gram-1/configs"
	"my-gram-1/handlers"
	"my-gram-1/middlewares"
	"my-gram-1/repositories"
	"my-gram-1/services"

	"github.com/gin-gonic/gin"
)

func MainRouter() {
	r := gin.Default()
	db := configs.GetDB()

	userRepository := repositories.NewUserReposittory(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	userRouter := r.Group("/users")

	{
		userRouter.POST("/register", userHandler.Register)
		userRouter.POST("/login", userHandler.Login)
		userRouter.Use(middlewares.Authenthication())
		userRouter.GET("/profile", userHandler.Profile)
		userRouter.PUT("/", userHandler.Update)
		userRouter.DELETE("/", userHandler.Delete)
	}
	r.Run()
}
