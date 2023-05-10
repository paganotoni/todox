package todo

import (
	"net/http"
	"paganotoni/todox/internal/database"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	conn := database.FromContext(r.Context())
	err := delete(conn, uuid.FromStringOrNil(chi.URLParam(r, "id")))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
