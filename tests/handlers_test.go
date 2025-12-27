package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"todos_manager/internal/errs"
	"todos_manager/internal/handlers"
	"todos_manager/internal/models"
)

func TestHandlers_CreateTodo_Success(t *testing.T) {
	svc := &MockService{
		CreateTodoFunc: func(req models.CreateTodoInput) (*models.Todo, error) {
			return &models.Todo{ID: 52, Title: req.Title}, nil
		},
	}
	handler := handlers.NewTodoHandler(svc)

	body, _ := json.Marshal(models.CreateTodoInput{
		Title:       "title",
		Description: "desc",
	})

	mux := http.NewServeMux()
	mux.HandleFunc("POST /todos", handler.CreateTodo)

	req := httptest.NewRequest("POST", "/todos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}

	var response models.Todo
	json.Unmarshal(w.Body.Bytes(), &response)

	if response.ID != 52 {
		t.Errorf("Expected id 52, got %d", response.ID)
	}
}

func TestHandlers_CreateTodo_InvalidJSON(t *testing.T) {
	svc := &MockService{}
	handler := handlers.NewTodoHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /todos", handler.CreateTodo)

	req := httptest.NewRequest("POST", "/todos", bytes.NewReader([]byte("{title")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 invalid JSON, got %d", w.Code)
	}
}

func TestHandlers_CreateTodo_ValidationError(t *testing.T) {
	svc := &MockService{
		CreateTodoFunc: func(models.CreateTodoInput) (*models.Todo, error) {
			return nil, errs.ValidationError
		},
	}
	handler := handlers.NewTodoHandler(svc)

	body, _ := json.Marshal(models.CreateTodoInput{Title: "", Description: "desc"})

	mux := http.NewServeMux()
	mux.HandleFunc("POST /todos", handler.CreateTodo)

	req := httptest.NewRequest("POST", "/todos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", w.Code)
	}
}

func TestHandlers_CreateTodo_InternalError(t *testing.T) {
	svc := &MockService{
		CreateTodoFunc: func(models.CreateTodoInput) (*models.Todo, error) {
			return nil, errors.New("internal error")
		},
	}
	handler := handlers.NewTodoHandler(svc)

	body, _ := json.Marshal(models.CreateTodoInput{Title: "Test", Description: "Desc"})

	mux := http.NewServeMux()
	mux.HandleFunc("POST /todos", handler.CreateTodo)

	req := httptest.NewRequest("POST", "/todos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected 500, got %d", w.Code)
	}
}

func TestHandlers_GetTodo_Success(t *testing.T) {
	svc := &MockService{
		GetTodoFunc: func(id int) (*models.Todo, error) {
			return &models.Todo{ID: id, Title: "asvas"}, nil
		},
	}
	handler := handlers.NewTodoHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /todos/{id}", handler.GetTodo)

	req := httptest.NewRequest("GET", "/todos/52", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}

	var response models.Todo
	json.Unmarshal(w.Body.Bytes(), &response)

	if response.ID != 52 {
		t.Errorf("Expected id 52, got %d", response.ID)
	}
}

func TestHandlers_GetTodo_InvalidID(t *testing.T) {
	svc := &MockService{}
	handler := handlers.NewTodoHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /todos/{id}", handler.GetTodo)

	req := httptest.NewRequest("GET", "/todos/notnumber", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", w.Code)
	}
}

func TestHandlers_GetTodo_NotFound(t *testing.T) {
	svc := &MockService{
		GetTodoFunc: func(int) (*models.Todo, error) {
			return nil, errs.ErrNotFound
		},
	}
	handler := handlers.NewTodoHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /todos/{id}", handler.GetTodo)

	req := httptest.NewRequest("GET", "/todos/52", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", w.Code)
	}
}

func TestHandlers_UpdateTodo_Success(t *testing.T) {
	svc := &MockService{
		UpdateTodoFunc: func(id int, req models.UpdateTodoInput) (*models.Todo, error) {
			return &models.Todo{ID: id, Title: req.Title, Completed: req.Completed}, nil
		},
	}
	handler := handlers.NewTodoHandler(svc)

	body, _ := json.Marshal(models.UpdateTodoInput{
		Title:       "new",
		Description: "new",
		Completed:   true,
	})

	mux := http.NewServeMux()
	mux.HandleFunc("PUT /todos/{id}", handler.UpdateTodo)

	req := httptest.NewRequest("PUT", "/todos/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}

	var response models.Todo
	json.Unmarshal(w.Body.Bytes(), &response)

	if !response.Completed {
		t.Error("Expected completed = true")
	}
}

func TestHandlers_UpdateTodo_InvalidJSON(t *testing.T) {
	svc := &MockService{}
	handler := handlers.NewTodoHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("PUT /todos/{id}", handler.UpdateTodo)

	req := httptest.NewRequest("PUT", "/todos/1", bytes.NewReader([]byte("{title")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", w.Code)
	}
}

func TestHandlers_UpdateTodo_ValidationError(t *testing.T) {
	svc := &MockService{
		UpdateTodoFunc: func(int, models.UpdateTodoInput) (*models.Todo, error) {
			return nil, errs.ValidationError
		},
	}
	handler := handlers.NewTodoHandler(svc)

	body, _ := json.Marshal(models.UpdateTodoInput{
		Title:       "",
		Description: "Desc",
		Completed:   false,
	})

	mux := http.NewServeMux()
	mux.HandleFunc("PUT /todos/{id}", handler.UpdateTodo)

	req := httptest.NewRequest("PUT", "/todos/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", w.Code)
	}
}

func TestHandlers_DeleteTodo_Success(t *testing.T) {
	svc := &MockService{
		DeleteTodoFunc: func(int) error {
			return nil
		},
	}
	handler := handlers.NewTodoHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /todos/{id}", handler.DeleteTodo)

	req := httptest.NewRequest("DELETE", "/todos/1", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected 204, got %d", w.Code)
	}
}

func TestHandlers_DeleteTodo_InvalidID(t *testing.T) {
	svc := &MockService{}
	handler := handlers.NewTodoHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /todos/{id}", handler.DeleteTodo)

	req := httptest.NewRequest("DELETE", "/todos/notnumber", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", w.Code)
	}
}

func TestHandlers_DeleteTodo_NotFound(t *testing.T) {
	svc := &MockService{
		DeleteTodoFunc: func(int) error {
			return errs.ErrNotFound
		},
	}
	handler := handlers.NewTodoHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /todos/{id}", handler.DeleteTodo)

	req := httptest.NewRequest("DELETE", "/todos/52", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", w.Code)
	}
}

func TestHandlers_ListTodos_Success(t *testing.T) {
	svc := &MockService{
		ListTodosFunc: func() ([]*models.Todo, error) {
			return []*models.Todo{
				{ID: 1, Title: "task 1"},
				{ID: 2, Title: "task 2"},
			}, nil
		},
	}
	handler := handlers.NewTodoHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /todos", handler.ListTodos)

	req := httptest.NewRequest("GET", "/todos", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}

	var response []models.Todo
	json.Unmarshal(w.Body.Bytes(), &response)

	if len(response) != 2 {
		t.Errorf("Expected 2 todos, got %d", len(response))
	}
}

func TestHandlers_CompleteTodo_Success(t *testing.T) {
	svc := &MockService{
		CompleteTodoFunc: func(id int, completed bool) (*models.Todo, error) {
			return &models.Todo{
				ID:          id,
				Title:       "title",
				Description: "desc",
				Completed:   completed,
			}, nil
		},
	}
	handler := handlers.NewTodoHandler(svc)

	body, _ := json.Marshal(models.CompleteInput{Completed: true})

	mux := http.NewServeMux()
	mux.HandleFunc("PATCH /todos/{id}/complete", handler.CompleteTodo)

	req := httptest.NewRequest("PATCH", "/todos/1/complete", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}

	var response models.Todo
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if !response.Completed {
		t.Error("Expected Completed = true")
	}
	if response.ID != 1 {
		t.Errorf("Expected ID 1, got %d", response.ID)
	}
}

func TestHandlers_CompleteTodo_NotFound(t *testing.T) {
	svc := &MockService{
		CompleteTodoFunc: func(id int, completed bool) (*models.Todo, error) {
			return nil, errs.ErrNotFound
		},
	}
	handler := handlers.NewTodoHandler(svc)

	body, _ := json.Marshal(models.CompleteInput{Completed: true})

	mux := http.NewServeMux()
	mux.HandleFunc("PATCH /todos/{id}/complete", handler.CompleteTodo)

	req := httptest.NewRequest("PATCH", "/todos/52/complete", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected 404 Not Found, got %d", w.Code)
	}
}
