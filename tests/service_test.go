package tests

import (
	"errors"
	"testing"
	"todos_manager/internal/errs"
	"todos_manager/internal/models"
	"todos_manager/internal/service"
)

func TestService_CerateTodo_Success(t *testing.T) {
	storage := &MockStorage{
		CreateTodoFunc: func(todo *models.Todo) (*models.Todo, error) {
			return &models.Todo{ID: 1, Title: todo.Title, Description: todo.Description}, nil
		},
	}

	svc := service.NewTodoService(storage)
	todo, err := svc.CreateTodo(models.CreateTodoInput{
		Title:       "aaa",
		Description: "bbb",
	})

	if err != nil {
		t.Fatalf("Failed to create todo: %v", err)
	}

	if todo.ID != 1 {
		t.Errorf("Expected id 1, got %d", todo.ID)
	}
}

func TestService_CreateTodo_NoTitle(t *testing.T) {
	storage := &MockStorage{}
	svc := service.NewTodoService(storage)
	_, err := svc.CreateTodo(models.CreateTodoInput{
		Title:       "",
		Description: "123",
	})

	if !errors.Is(err, errs.ValidationError) {
		t.Error("Expected ValidationError")
	}
}

func TestService_GetTodo_Success(t *testing.T) {
	storage := &MockStorage{
		GetTodoFunc: func(id int) (*models.Todo, error) {
			return &models.Todo{ID: id, Title: "qwert"}, nil
		},
	}
	svc := service.NewTodoService(storage)

	todo, err := svc.GetTodo(52)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if todo.ID != 52 {
		t.Errorf("Expected id 52, got %d", todo.ID)
	}
}

func TestService_GetTodo_NotFound(t *testing.T) {
	storage := &MockStorage{
		GetTodoFunc: func(id int) (*models.Todo, error) {
			return nil, errs.ErrNotFound
		},
	}
	svc := service.NewTodoService(storage)

	_, err := svc.GetTodo(52)

	if !errors.Is(err, errs.ErrNotFound) {
		t.Error("Expected ErrNotFound")
	}
}

func TestService_UpdateTodo_Success(t *testing.T) {
	storage := &MockStorage{
		GetTodoFunc: func(id int) (*models.Todo, error) {
			return &models.Todo{ID: id, Title: "old", Completed: false}, nil
		},
		UpdateTodoFunc: func(id int, updated *models.Todo) (*models.Todo, error) {
			return updated, nil
		},
	}
	svc := service.NewTodoService(storage)

	todo, err := svc.UpdateTodo(1, models.UpdateTodoInput{
		Title:       "new",
		Description: "new",
		Completed:   true,
	})
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	if todo.Title != "new" {
		t.Errorf("Expected title 'new', got %s", todo.Title)
	}
}

func TestService_UpdateTodo_NoTitle(t *testing.T) {
	storage := &MockStorage{}
	svc := service.NewTodoService(storage)
	_, err := svc.UpdateTodo(1, models.UpdateTodoInput{
		Title:       "",
		Description: "qgw",
		Completed:   true,
	})

	if !errors.Is(err, errs.ValidationError) {
		t.Error("Expected ValidationError")
	}
}

func TestService_UpdateTodo_NotFound(t *testing.T) {
	storage := &MockStorage{
		GetTodoFunc: func(int) (*models.Todo, error) {
			return nil, errs.ErrNotFound
		},
	}
	svc := service.NewTodoService(storage)

	_, err := svc.UpdateTodo(421, models.UpdateTodoInput{
		Title:       "new",
		Description: "new",
		Completed:   true,
	})

	if !errors.Is(err, errs.ErrNotFound) {
		t.Error("Expected ErrNotFound")
	}
}

func TestService_ListTodos_Success(t *testing.T) {
	storage := &MockStorage{
		ListTodosFunc: func() ([]*models.Todo, error) {
			return []*models.Todo{
				{ID: 1, Title: "first"},
				{ID: 2, Title: "second"},
			}, nil
		},
	}
	svc := service.NewTodoService(storage)

	todos, err := svc.ListTodos()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(todos) != 2 {
		t.Errorf("Expected 2 todos, got %d", len(todos))
	}
}

func TestService_DeleteTodo_NotFound(t *testing.T) {
	storage := &MockStorage{
		GetTodoFunc: func(int) (*models.Todo, error) {
			return nil, errs.ErrNotFound
		},
	}
	svc := service.NewTodoService(storage)

	err := svc.DeleteTodo(65)

	if !errors.Is(err, errs.ErrNotFound) {
		t.Error("Expected ErrNotFound")
	}
}

func TestService_DeleteTodo_Success(t *testing.T) {
	mock := &MockStorage{
		GetTodoFunc: func(id int) (*models.Todo, error) {
			return &models.Todo{ID: id, Title: "ecom"}, nil
		},
		DeleteTodoFunc: func(id int) error {
			return nil
		},
	}
	svc := service.NewTodoService(mock)

	err := svc.DeleteTodo(1)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
