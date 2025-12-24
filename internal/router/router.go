package router

import (
	"net/http"
	"todos_manager/internal/handlers"
)

func NewRouter(handler *handlers.TodoHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /todos", handler.CreateTodo)
	mux.HandleFunc("GET /todos", handler.ListTodos)
	mux.HandleFunc("GET /todos/{id}", handler.GetTodo)
	mux.HandleFunc("PUT /todos/{id}", handler.UpdateTodo)
	mux.HandleFunc("DELETE /todos/{id}", handler.DeleteTodo)
	return mux
}
