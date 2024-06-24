package todos

import (
	"database/sql"

	"github.com/gofrs/uuid/v5"
	"github.com/jmoiron/sqlx"
)

type service struct {
	db func() (*sql.DB, error)
}

func NewService(db func() (*sql.DB, error)) *service {
	return &service{db: db}
}

func (s *service) Find(id uuid.UUID) (t Instance, err error) {
	cx, err := s.db()
	if err != nil {
		return t, err
	}

	conn := sqlx.NewDb(cx, "sqlite3")
	err = conn.Get(&t, "SELECT * FROM todos WHERE id = $1", id)
	return t, err
}

func (s *service) Create(todo *Instance) error {
	todo.ID = uuid.Must(uuid.NewV4())

	cx, err := s.db()
	if err != nil {
		return err
	}

	conn := sqlx.NewDb(cx, "sqlite3")
	_, err = conn.NamedExec(`INSERT INTO todos (id, content, completed) VALUES (:id, :content, :completed)`, todo)
	return err
}

func (s *service) List() ([]Instance, error) {
	var list []Instance

	cx, err := s.db()
	if err != nil {
		return list, err
	}
	conn := sqlx.NewDb(cx, "sqlite3")
	err = conn.Select(&list, "SELECT * FROM todos ORDER BY completed ASC")
	return list, err
}

func (s *service) Search(term string) ([]Instance, error) {
	var list []Instance

	cx, err := s.db()
	if err != nil {
		return list, err
	}

	conn := sqlx.NewDb(cx, "sqlite3")
	err = conn.Select(&list, "SELECT * FROM todos WHERE content LIKE $1 ORDER BY completed ASC", "%"+term+"%")
	return list, err
}

func (s *service) Delete(id uuid.UUID) error {
	cx, err := s.db()
	if err != nil {
		return err
	}

	conn := sqlx.NewDb(cx, "sqlite3")
	_, err = conn.Exec("DELETE FROM todos WHERE id = $1", id)
	return err
}

func (s *service) Update(todo *Instance) error {
	cx, err := s.db()
	if err != nil {
		return err
	}

	conn := sqlx.NewDb(cx, "sqlite3")
	_, err = conn.NamedExec("UPDATE todos SET content = :content WHERE id = :id", todo)
	return err
}

func (s *service) SetCompleted(id uuid.UUID, completed bool) error {
	cx, err := s.db()
	if err != nil {
		return err
	}

	conn := sqlx.NewDb(cx, "sqlite3")
	_, err = conn.Exec("UPDATE todos SET completed = $1 WHERE id = $2", completed, id)
	return err
}
