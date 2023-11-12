package todos

import (
	"github.com/gofrs/uuid/v5"
	"github.com/jmoiron/sqlx"
)

var _ Service = (*service)(nil)

type service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) *service {
	return &service{db: db}
}

func (s *service) Find(id uuid.UUID) (t Instance, err error) {
	err = s.db.Get(&t, "SELECT * FROM todos WHERE id = $1", id)
	return t, err
}

func (s *service) Create(todo *Instance) error {
	todo.ID = uuid.Must(uuid.NewV4())
	_, err := s.db.NamedExec(`INSERT INTO todos (id, content, completed) VALUES (:id, :content, :completed)`, todo)
	return err
}

func (s *service) List() ([]Instance, error) {
	var list []Instance
	err := s.db.Select(&list, "SELECT * FROM todos ORDER BY completed ASC")
	return list, err
}

func (s *service) Search(term string) ([]Instance, error) {
	var list []Instance
	err := s.db.Select(&list, "SELECT * FROM todos WHERE content LIKE $1", "%"+term+"%")
	return list, err
}

func (s *service) Delete(id uuid.UUID) error {
	_, err := s.db.Exec("DELETE FROM todos WHERE id = $1", id)
	return err
}

func (s *service) Update(todo *Instance) error {
	_, err := s.db.NamedExec("UPDATE todos SET content = :content WHERE id = :id", todo)
	return err
}

func (s *service) SetCompleted(id uuid.UUID, completed bool) error {
	_, err := s.db.Exec("UPDATE todos SET completed = $1 WHERE id = $2", completed, id)
	return err
}
