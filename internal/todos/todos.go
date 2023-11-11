package todos

import "github.com/gofrs/uuid/v5"

// Instance is a model that represents a todo item
// in the database
type Instance struct {
	ID        uuid.UUID
	Content   string
	Completed bool
}

// TodoService is the interface that wraps the basic CRUD operations
// for the Todo model
type Service interface {
	Find(id uuid.UUID) (Instance, error)
	Create(todo *Instance) error
	List() ([]Instance, error)
	Search(term string) ([]Instance, error)
	Delete(id uuid.UUID) error
	Update(todo *Instance) error
	SetCompleted(id uuid.UUID, completed bool) error
}
