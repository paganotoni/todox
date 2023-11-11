package todos

import (
	"net/http"

	"github.com/leapkit/core/render"
)

func Index(w http.ResponseWriter, r *http.Request) {
	todos := r.Context().Value("todoService").(Service)

	list, err := todos.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rw := render.FromCtx(r.Context())
	rw.Set("list", list)

	err = rw.Render("todos/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
