package todos

import (
	"net/http"
	"paganotoni/todox/internal/models"

	"github.com/leapkit/core/render"
)

func Search(w http.ResponseWriter, r *http.Request) {
	todos := r.Context().Value("todoService").(models.TodoService)

	list, err := todos.Search(r.FormValue("keyword"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rw := render.FromCtx(r.Context())
	rw.Set("list", list)

	err = rw.RenderClean("todos/list.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
