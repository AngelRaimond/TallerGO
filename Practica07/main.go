package main

import (
	"fmt"
	"practica07/controllers"
	"practica07/middleware"
	"practica07/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())

	userService := services.NewUserService()
	userController := controllers.NewUserController(userService)

	r.Static("/public", "./public")
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/", func(c *gin.Context) {
		c.File("./public/index.html")
	})

	v1 := r.Group("/v1")
	v1.Use(middleware.APIKeyAuthMiddleware())
	{
		v1.GET("/users", userController.GetUsers)
		v1.POST("/users", userController.CreateUser)
		v1.DELETE("/users/:id", userController.DeleteUser)
		v1.PUT("/users/:id", userController.UpdateUser)
	}

	fmt.Println("Server listening on port 3000")
	r.Run(":3000")
}
