package todo

import (
	"net/http"
	"paganotoni/todox/database"

	"github.com/go-chi/chi/v5"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	conn := database.FromContext(r.Context())

	_, err := conn.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
