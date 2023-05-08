package todo

import (
	"net/http"
	"paganotoni/todox/database"
	"paganotoni/todox/internal"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	conn := database.FromContext(r.Context())
	todo, err := find(conn, uuid.FromStringOrNil(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	err = internal.Render(w, "main", todo, "todo/edit.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
