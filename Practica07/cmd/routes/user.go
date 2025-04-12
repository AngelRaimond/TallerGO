package routes

import (
    "gin/cmd/controllers"
    "gin/cmd/middleware"
    "gin/cmd/services"

    "github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, userService *services.UserService) {
    admin := r.Group("/admin")
    admin.Use(middleware.APIKeyAuthMiddleware())
    // Controller
    userController := controllers.NewUserController(userService)

    admin.GET("/users", userController.GetUsers)
    admin.POST("/users", userController.CreateUser)
    admin.PUT("/users/:id", userController.UpdateUser)
    admin.DELETE("/users/:id", userController.DeleteUser)
}
