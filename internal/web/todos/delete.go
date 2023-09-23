package todos

import (
	"net/http"
	"paganotoni/todox/internal/models"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	todos := r.Context().Value("todoService").(models.TodoService)

	id := uuid.FromStringOrNil(chi.URLParam(r, "id"))
	err := todos.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
