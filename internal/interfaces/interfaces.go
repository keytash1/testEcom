package interfaces

import "todos_manager/internal/models"

type Storage interface {
	CreateTodo(todo *models.Todo) (*models.Todo, error)
	ListTodos() ([]*models.Todo, error)
	GetTodo(id int) (*models.Todo, error)
	UpdateTodo(id int, newTodo *models.Todo) (*models.Todo, error)
	DeleteTodo(id int) error
}

type Service interface {
	CreateTodo(req models.CreateTodoInput) (*models.Todo, error)
	ListTodos() ([]*models.Todo, error)
	GetTodo(id int) (*models.Todo, error)
	UpdateTodo(id int, req models.UpdateTodoInput) (*models.Todo, error)
	DeleteTodo(id int) error
}
