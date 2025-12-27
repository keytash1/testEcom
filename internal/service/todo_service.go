package service

import (
	"todos_manager/internal/errs"
	"todos_manager/internal/interfaces"
	"todos_manager/internal/models"
)

type TodoService struct {
	storage interfaces.Storage
}

func NewTodoService(storage interfaces.Storage) *TodoService {
	return &TodoService{storage: storage}
}

func (s *TodoService) CreateTodo(req models.CreateTodoInput) (*models.Todo, error) {
	if req.Title == "" {
		return nil, errs.ValidationError
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
		return nil, errs.ValidationError
	}
	oldTodo, err := s.GetTodo(id)
	if err != nil {
		return nil, err
	}
	newTodo := &models.Todo{
		ID:          oldTodo.ID,
		Title:       req.Title,
		Description: req.Description,
		Completed:   req.Completed,
	}
	return s.storage.UpdateTodo(id, newTodo)
}

func (s *TodoService) CompleteTodo(id int, completed bool) (*models.Todo, error) {
	todo, err := s.GetTodo(id)
	if err != nil {
		return nil, err
	}
	todo.Completed = completed
	return s.storage.UpdateTodo(id, todo)
}
