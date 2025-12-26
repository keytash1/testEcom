package service

import (
	"todos_manager/internal/errs"
	"todos_manager/internal/models"
	"todos_manager/internal/storage"
)

type TodoService struct {
	storage *storage.Storage
}

func NewTodoService(storage *storage.Storage) *TodoService {
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
