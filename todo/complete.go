package todo

import (
	"net/http"
	"paganotoni/todox/database"
	"paganotoni/todox/internal"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
)

func Complete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id := uuid.FromStringOrNil(chi.URLParam(r, "id"))
	conn := database.FromContext(r.Context())

	todo, err := find(conn, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	todo.Completed = (r.FormValue("completed") == "on")

	if err = complete(conn, todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	err = internal.Render(w, "todo", todo, "todo/todo.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
