package todo

import (
	"net/http"
	"paganotoni/todox"
	"paganotoni/todox/database"
	"paganotoni/todox/internal"

	"github.com/go-chi/chi/v5"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	r.ParseForm()

	var todo todox.Todo
	conn := database.FromContext(r.Context())
	err := conn.Get(&todo, "SELECT * FROM todos WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	todo.Content = r.FormValue("content")

	_, err = conn.NamedExec(`UPDATE todos SET content = :content WHERE id = :id`, todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = internal.Render(w, "todo", todo, "todo/todo.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
