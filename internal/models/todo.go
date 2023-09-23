package models

import (
	"github.com/gofrs/uuid/v5"
)

// Todo is a model that represents a todo item
// in the database
type Todo struct {
	ID        uuid.UUID
	Content   string
	Completed bool
}

// TodoService is the interface that wraps the basic CRUD operations
// for the Todo model
type TodoService interface {
	Find(id uuid.UUID) (Todo, error)
	Create(todo *Todo) error
	List() ([]Todo, error)
	Search(term string) ([]Todo, error)
	Delete(id uuid.UUID) error
	Update(todo *Todo) error
	SetCompleted(id uuid.UUID, completed bool) error
}
