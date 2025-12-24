package service

import (
	"todos_manager/internal/storage"
)

type TodoService struct {
	storage *storage.Storage
}

func NewTodoService(storage *storage.Storage) *TodoService {
	return &TodoService{storage: storage}
}
