package todo

import (
	"net/http"
	"paganotoni/todox"
	"paganotoni/todox/database"
	"paganotoni/todox/internal"

	"github.com/gofrs/uuid"
)

func Create(w http.ResponseWriter, r *http.Request) {
	todo := todox.Todo{
		ID:        uuid.Must(uuid.NewV4()),
		Content:   r.FormValue("content"),
		Completed: false,
	}

	conn := database.FromContext(r.Context())
	_, err := conn.NamedExec("INSERT INTO todos (id, content, completed) VALUES (:id, :content, :completed)", todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	var list []todox.Todo
	err = conn.Select(&list, "SELECT * FROM todos")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	err = internal.Render(w, "list", list, "todo/index.html", "todo/todo.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
