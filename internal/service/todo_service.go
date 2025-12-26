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

func (s *TodoService) CreateTodo(req models.CreateTodoInput) (*models.Todo, error) {
	if req.Title == "" {
		return nil, &ValidationError{Message: "Title is required"}
	}
	todo := &models.Todo{
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
	}
	return s.storage.CreateTodo(todo)
}

func (s *TodoService) ListTodos() ([]*models.Todo, error) {
	return s.storage.ListTodos()
}

func (s *TodoService) GetTodo(id int) (*models.Todo, error) {
	todo, err := s.storage.GetTodo(id)
	if err != nil {
		return nil, err
	}
	if todo == nil {
		return nil, &NotFoundError{ID: id}
	}
	return todo, nil
}

func (s *TodoService) DeleteTodo(id int) error {
	_, err := s.GetTodo(id)
	if err != nil {
		return err
	}
	return s.storage.DeleteTodo(id)
}

func (s *TodoService) UpdateTodo(id int, req models.UpdateTodoInput) (*models.Todo, error) {
	if req.Title == "" {
		return nil, &ValidationError{Message: "Title is required"}
	}
	oldtodo, err := s.GetTodo(id)
	if err != nil {
		return nil, err
	}
	newtodo := &models.Todo{
		ID:          oldtodo.ID,
		Title:       req.Title,
		Description: req.Description,
		Completed:   req.Completed,
	}
	return s.storage.UpdateTodo(id, newtodo)
}
