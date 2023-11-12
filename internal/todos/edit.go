package todos

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
	"github.com/leapkit/core/render"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	todos := r.Context().Value("todoService").(*service)

	id := uuid.FromStringOrNil(chi.URLParam(r, "id"))
	todo, err := todos.Find(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	rw := render.FromCtx(r.Context())
	rw.Set("todo", todo)

	err = rw.RenderClean("todos/edit.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
