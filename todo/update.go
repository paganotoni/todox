package todo

import (
	"net/http"
	"paganotoni/todox"
	"paganotoni/todox/internal"

	"github.com/go-chi/chi/v5"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	r.ParseForm()

	var todo *todox.Todo
	for index, tx := range todox.List {
		if id == tx.ID.String() {
			todox.List[index].Content = r.FormValue("content")
			todo = &todox.List[index]

			break
		}
	}

	if todo == nil {
		http.Error(w, "Not found", http.StatusNotFound)

		return
	}

	err := internal.Render(w, "todo", todo, "todo/todo.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
