package todo

import (
	"net/http"
	"paganotoni/todox"
	"paganotoni/todox/internal"

	"github.com/go-chi/chi/v5"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var todo *todox.Todo
	for _, t := range todox.List {
		if id == t.ID.String() {
			todo = &t

			break
		}
	}

	if todo == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	err := internal.Render(w, "main", todo, "todo/edit.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
