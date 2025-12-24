package main

import (
	"log"
	"net/http"
	"todos_manager/internal/handlers"
	"todos_manager/internal/router"
	"todos_manager/internal/service"
	"todos_manager/internal/storage"
)

func main() {
	storage := storage.NewStorage()
	service := service.NewTodoService(storage)
	handler := handlers.NewTodoHandler(service)

	mux := router.NewRouter(handler)

	log.Println("starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("server failed:", err)
	}
}
