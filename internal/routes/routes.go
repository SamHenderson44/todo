package routes

import (
	"fmt"
	"net/http"

	handlers "github.com/SamHenderson44/todo/internal/handlers"
)

func InitRoutes() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /todos", handlers.HandleGet)
	mux.HandleFunc("POST /todos", handlers.HandleCreateToDo)
	mux.HandleFunc("PATCH /todos/{id}/update", handlers.HandleUpdateStatus)

	if err := http.ListenAndServe("localhost:8081", mux); err != nil {
		fmt.Println(err.Error())
	}
}
