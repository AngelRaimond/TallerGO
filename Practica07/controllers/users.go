package controllers

import (
    "gin/cmd/services"
    "io"
    "net/http"
    "github.com/gin-gonic/gin"
)

type UserController struct {
    userService *services.UserService
}

func GetUsers(c *gin.Context) {
    c.JSON(http.StatusOK, []string{"user1", "user2"})
}

func CreateUser(c *gin.Context) {
    body, err := io.ReadAll(c.Request.Body)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "error parsing body",
        })
        return
    }
    c.JSON(http.StatusOK, body)
}

func UpdateUser(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{
        "message": "user updated",
        "id":      id,
    })
}

func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{
        "message": "user deleted",
        "id":      id,
    })
}
