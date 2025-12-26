package service

import (
	"fmt"
	"todos_manager/internal/models"
	"todos_manager/internal/storage"
)

type TodoService struct {
	storage *storage.Storage
}

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

type NotFoundError struct {
	ID int
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Todo %d not found", e.ID)
}

func NewTodoService(storage *storage.Storage) *TodoService {
	return &TodoService{storage: storage}
}

func (s *TodoService) CreateTodo(req models.CreateTodoInput) (models.Todo, error) {

}

func (s *TodoService) ListTodos() ([]models.Todo, error) {

}

func (s *TodoService) GetTodo(id int) (models.Todo, error) {

}

func (s *TodoService) DeleteTodo(id int) error {

}

func (s *TodoService) UpdateTodo(id int, req models.UpdateTodoInput) (models.Todo, error) {

}
