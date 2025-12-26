package storage

import (
	"sync"
	"todos_manager/internal/errs"
	"todos_manager/internal/models"
)

type Storage struct {
	todos  map[int]*models.Todo
	mu     sync.RWMutex
	nextID int
}

func NewStorage() *Storage {
	return &Storage{
		nextID: 1,
		todos:  make(map[int]*models.Todo),
	}
}

func (s *Storage) CreateTodo(todo *models.Todo) (*models.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo.ID = s.nextID
	s.todos[todo.ID] = todo
	s.nextID++

	return todo, nil
}

func (s *Storage) ListTodos() ([]*models.Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todos := make([]*models.Todo, 0, len(s.todos))
	for _, todo := range s.todos {
		todos = append(todos, todo)
	}

	return todos, nil
}

func (s *Storage) GetTodo(id int) (*models.Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todo, exists := s.todos[id]

	if !exists {
		return nil, errs.ErrNotFound
	}

	return todo, nil
}

func (s *Storage) DeleteTodo(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.todos[id]; !exists {
		return errs.ErrNotFound
	}

	delete(s.todos, id)
	return nil
}

func (s *Storage) UpdateTodo(id int, updated *models.Todo) (*models.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	todo, exists := s.todos[id]
	if !exists {
		return nil, errs.ErrNotFound
	}
	// тк у нас put то полностью меняем (был бы patch меняли бы только то что пришло)
	todo.Title = updated.Title
	todo.Description = updated.Description
	todo.Completed = updated.Completed

	return todo, nil
}
