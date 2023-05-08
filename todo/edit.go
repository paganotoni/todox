package todo

import (
	"net/http"
	"paganotoni/todox"
	"paganotoni/todox/database"
	"paganotoni/todox/internal"

	"github.com/go-chi/chi/v5"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	conn := database.FromContext(r.Context())
	rows, err := conn.Query("SELECT * FROM todos WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	var todo todox.Todo
	for rows.Next() {
		err = rows.Scan(&todo.ID, &todo.Content, &todo.Completed)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		break
	}

	err = internal.Render(w, "main", todo, "todo/edit.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
