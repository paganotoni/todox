package sqlite

import (
	"paganotoni/todox/internal/models"

	"github.com/gofrs/uuid/v5"
	"github.com/jmoiron/sqlx"
)

var _ models.TodoService = (*TodoService)(nil)

type TodoService struct {
	db *sqlx.DB
}

func NewTodoService(db *sqlx.DB) *TodoService {
	return &TodoService{db: db}
}

func (s *TodoService) Find(id uuid.UUID) (t models.Todo, err error) {
	err = s.db.Get(&t, "SELECT * FROM todos WHERE id = $1", id)
	return t, err
}

func (s *TodoService) Create(todo *models.Todo) error {
	todo.ID = uuid.Must(uuid.NewV4())
	_, err := s.db.NamedExec(`INSERT INTO todos (id, content, completed) VALUES (:id, :content, :completed)`, todo)
	return err
}

func (s *TodoService) List() ([]models.Todo, error) {
	var list []models.Todo
	err := s.db.Select(&list, "SELECT * FROM todos")
	return list, err
}

func (s *TodoService) Search(term string) ([]models.Todo, error) {
	var list []models.Todo
	err := s.db.Select(&list, "SELECT * FROM todos WHERE content LIKE $1", "%"+term+"%")
	return list, err
}

func (s *TodoService) Delete(id uuid.UUID) error {
	_, err := s.db.Exec("DELETE FROM todos WHERE id = $1", id)
	return err
}

func (s *TodoService) Update(todo *models.Todo) error {
	_, err := s.db.NamedExec("UPDATE todos SET content = :content WHERE id = :id", todo)
	return err
}

func (s *TodoService) SetCompleted(id uuid.UUID, completed bool) error {
	_, err := s.db.Exec("UPDATE todos SET completed = $1 WHERE id = $2", completed, id)
	return err
}