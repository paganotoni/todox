package todos

import (
	"net/http"

	"github.com/gofrs/uuid/v5"
)

func Show(w http.ResponseWriter, r *http.Request) {
	todos := r.Context().Value("todoService").(*service)
	id := uuid.FromStringOrNil(r.PathValue("id"))
	todo, err := todos.Find(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	todoHTML(todo).Render(w)
}
