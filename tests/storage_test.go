package tests

import (
	"testing"
	"todos_manager/internal/models"
	"todos_manager/internal/storage"
)

// дублирование id - невозможно тк сами задаем а не юзер присылает

func TestStorage_CreateTodo(t *testing.T) {
	s := storage.NewStorage()
	todo := &models.Todo{Title: "Test"}
	created, err := s.CreateTodo(todo)
	if err != nil {
		t.Fatalf("Create todo failed: %v", err)
	}
	if created.ID != 1 {
		t.Errorf("Expected id 1, got %d", created.ID)
	}
	if created.Title != "Test" {
		t.Errorf("Expected title 'Test', got %s", created.Title)
	}
}

func TestStorage_GetTodo_NotFound(t *testing.T) {
	s := storage.NewStorage()
	_, err := s.GetTodo(2)
	if err == nil {
		t.Error("Expected ErrNotFound")
	}
}

func TestStorage_UpdateTodo(t *testing.T) {
	s := storage.NewStorage()

	created, _ := s.CreateTodo(&models.Todo{Title: "old"})
	updated, _ := s.UpdateTodo(created.ID, &models.Todo{Title: "new"})

	if updated.Title != "new" {
		t.Errorf("Expected title 'new', got %s", updated.Title)
	}
}

func TestStorage_DeleteTodo(t *testing.T) {
	s := storage.NewStorage()
	created, _ := s.CreateTodo(&models.Todo{Title: "aa"})
	err := s.DeleteTodo(created.ID)
	if err != nil {
		t.Fatalf("Delete todo failed: %v", err)
	}
	_, err = s.GetTodo(created.ID)
	if err == nil {
		t.Error("todo was not deleted")
	}
}
