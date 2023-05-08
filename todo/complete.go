package todo

import (
	"html/template"
	"net/http"
	"paganotoni/todox"
	"paganotoni/todox/database"

	"github.com/go-chi/chi/v5"
)

func Complete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	r.ParseForm()

	var todo todox.Todo
	conn := database.FromContext(r.Context())
	err := conn.Get(&todo, "SELECT * FROM todos WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	todo.Completed = (r.FormValue("completed") == "on")

	_, err = conn.NamedExec("UPDATE todos SET completed = :completed WHERE id = :id", todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("updated").ParseFS(todox.Templates, "todo/todo.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "todo", todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
