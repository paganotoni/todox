package todo

import (
	"net/http"
	"paganotoni/todox/internal"
	"paganotoni/todox/internal/database"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	r.ParseForm()

	conn := database.FromContext(r.Context())
	todo, err := find(conn, uuid.FromStringOrNil(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	todo.Content = r.FormValue("content")

	err = update(conn, todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	err = internal.Render(w, "todo", todo, "todo/todo.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
