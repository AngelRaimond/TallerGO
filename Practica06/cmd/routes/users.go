package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
	

	"github.com/gin-gonic/gin"
)
type User struct {
	ID     int    `json:"id"`
	Nombre string `json:"name"`
	Email  string `json:"email"`
}

var users []User

func SetupUserRoutes(r *gin.Engine) {
    r.GET("/users", func(c *gin.Context) {
        c.JSON(http.StatusOK, users)
    })
    r.POST("/users", func(c *gin.Context) {
        body, err := io.ReadAll(c.Request.Body)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "message": "error reading the body",
            })
            return
        }
        var user User
        err = json.Unmarshal(body, &user)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "message": "error parsing body",
            })
            return
        }
        user.ID = len(users) + 1
        users = append(users, user)
        c.JSON(http.StatusOK, user)
    })
	r.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println("user id", id)



	})
	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println("user id", id)

	})
}
		
	