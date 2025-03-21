package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	ID     int    `json:"id"`
	Nombre string `json:"name"`
	Email  string `json:"email"`
}

var usuarios []Usuario

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, usuarios)
}

func CreateUser(c *gin.Context) {
	var user Usuario
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parseando el JSON"})
		return
	}
	user.ID = len(usuarios) + 1
	usuarios = append(usuarios, user)
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var idInt int
	if _, err := fmt.Sscanf(id, "%d", &idInt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	for i, user := range usuarios {
		if user.ID == idInt {
			usuarios = append(usuarios[:i], usuarios[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var idInt int
	if _, err := fmt.Sscanf(id, "%d", &idInt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var userUpdate Usuario
	if err := c.BindJSON(&userUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parseando el JSON"})
		return
	}
	userUpdate.ID = idInt

	for i, user := range usuarios {
		if user.ID == idInt {
			usuarios[i] = userUpdate
			c.JSON(http.StatusOK, userUpdate)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func Index(c *gin.Context) {
	c.File("./public/index.html")
}

func main() {
	usuarios = append(usuarios, Usuario{
		ID:     1,
		Nombre: "Alfredo",
		Email:  "Alfredo@mail.com",
	})

	r := gin.Default()
	r.GET("/ping", Ping)
	r.GET("/", Index)

	v1 := r.Group("/v1")
	{
		v1.GET("/users", GetUsers)
		v1.POST("/users", CreateUser)
		v1.DELETE("/users/:id", DeleteUser)
		v1.PUT("/users/:id", UpdateUser)
	}

	fmt.Println("Servidor escuchando en el puerto 3000")
	r.Run(":3000")
}
