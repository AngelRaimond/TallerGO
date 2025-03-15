package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Usuario struct {
	ID     int    `json:"id"`
	Nombre string `json:"name"`
	Email  string `json:"email"`
}

var usuarios []Usuario

// Usuarios

func Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("test-header", "header")
		json.NewEncoder(w).Encode(usuarios)
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error al leer el body", http.StatusBadRequest)
		}
		var user Usuario
		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, "Error parseando el JSON", http.StatusBadRequest)
		}
		user.ID = len(usuarios) + 1
		usuarios = append(usuarios, user)

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("test-header", "header")
		json.NewEncoder(w).Encode(user)
	case http.MethodDelete:
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "ID no proporcionado", http.StatusBadRequest)
			return
		}
		var idInt int
		_, err := fmt.Sscanf(id, "%d", &idInt)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		for i, user := range usuarios {
			if user.ID == idInt {
				usuarios = append(usuarios[:i], usuarios[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)

	case http.MethodPut:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error al leer el body", http.StatusBadRequest)
			return
		}

		var userUpdate Usuario
		err = json.Unmarshal(body, &userUpdate)
		if err != nil {
			http.Error(w, "Error parseando el JSON", http.StatusBadRequest)
			return
		}

		for i, user := range usuarios {
			if user.ID == userUpdate.ID {
				usuarios[i] = userUpdate
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(userUpdate)
				return
			}
		}
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
	default:
		http.Error(w, "Método no permitido", 405)
	}
}

func Ping(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "pong")
	default:
		http.Error(w, "Método no permitido", 405)
	}
}
func Index(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile("./public/index.html")
	if err != nil {
		fmt.Fprintln(w, "error leyendo el html")
		return
	}
	fmt.Fprintln(w, string(content))
}
func main() {
	usuarios = append(usuarios, Usuario{
		ID:     1,
		Nombre: "Alfredo",
		Email:  "Alfredo@mail.com",
	})
	http.HandleFunc("/ping", Ping)
	http.HandleFunc("/v1/users", Users)
	http.HandleFunc("/", Index)

	fmt.Println("Servidor escuchando en el puerto 3000")
	http.ListenAndServe(":3000", nil)
}
