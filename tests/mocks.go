package tests

import (
	"todos_manager/internal/models"
)

type MockStorage struct {
	CreateTodoFunc func(*models.Todo) (*models.Todo, error)
	GetTodoFunc    func(int) (*models.Todo, error)
	ListTodosFunc  func() ([]*models.Todo, error)
	UpdateTodoFunc func(int, *models.Todo) (*models.Todo, error)
	DeleteTodoFunc func(int) error
}

func (m *MockStorage) CreateTodo(todo *models.Todo) (*models.Todo, error) {
	return m.CreateTodoFunc(todo)
}

func (m *MockStorage) GetTodo(id int) (*models.Todo, error) {
	return m.GetTodoFunc(id)
}

func (m *MockStorage) ListTodos() ([]*models.Todo, error) {
	return m.ListTodosFunc()
}

func (m *MockStorage) UpdateTodo(id int, updated *models.Todo) (*models.Todo, error) {
	return m.UpdateTodoFunc(id, updated)
}

func (m *MockStorage) DeleteTodo(id int) error {
	return m.DeleteTodoFunc(id)
}

type MockService struct {
	CreateTodoFunc func(models.CreateTodoInput) (*models.Todo, error)
	GetTodoFunc    func(int) (*models.Todo, error)
	ListTodosFunc  func() ([]*models.Todo, error)
	UpdateTodoFunc func(int, models.UpdateTodoInput) (*models.Todo, error)
	DeleteTodoFunc func(int) error
}

func (m *MockService) CreateTodo(req models.CreateTodoInput) (*models.Todo, error) {
	return m.CreateTodoFunc(req)
}

func (m *MockService) GetTodo(id int) (*models.Todo, error) {
	return m.GetTodoFunc(id)
}

func (m *MockService) ListTodos() ([]*models.Todo, error) {
	return m.ListTodosFunc()
}

func (m *MockService) UpdateTodo(id int, req models.UpdateTodoInput) (*models.Todo, error) {
	return m.UpdateTodoFunc(id, req)
}

func (m *MockService) DeleteTodo(id int) error {
	return m.DeleteTodoFunc(id)
}
