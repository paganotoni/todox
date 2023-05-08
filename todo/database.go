package todo

import (
	"paganotoni/todox"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

func find(db *sqlx.DB, id uuid.UUID) (todox.Todo, error) {
	var t todox.Todo
	err := db.Get(&t, "SELECT * FROM todos WHERE id = $1", id)
	return t, err
}

func create(db *sqlx.DB, todo todox.Todo) error {
	_, err := db.NamedExec(`INSERT INTO todos (id, content, completed) VALUES (:id, :content, :completed)`, todo)
	return err
}

func list(db *sqlx.DB) ([]todox.Todo, error) {
	var list []todox.Todo
	err := db.Select(&list, "SELECT * FROM todos")
	return list, err
}

func search(db *sqlx.DB, term string) ([]todox.Todo, error) {
	var list []todox.Todo
	err := db.Select(&list, "SELECT * FROM todos WHERE content LIKE $1", "%"+term+"%")
	return list, err
}

func delete(db *sqlx.DB, id uuid.UUID) error {
	_, err := db.Exec("DELETE FROM todos WHERE id = $1", id)
	return err
}

func complete(db *sqlx.DB, todo todox.Todo) error {
	_, err := db.NamedExec("UPDATE todos SET completed = :completed WHERE id = :id", todo)
	return err
}
