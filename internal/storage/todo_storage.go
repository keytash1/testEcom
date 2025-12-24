package storage

import (
	"sync"
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
