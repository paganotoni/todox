package todo

import (
	"html/template"
	"net/http"
	"paganotoni/todox"

	"github.com/go-chi/chi/v5"
)

func Complete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	r.ParseForm()

	var todo *todox.Todo
	for index, tx := range todox.List {
		if id == tx.ID.String() {
			todox.List[index].Completed = (r.FormValue("completed") == "on")
			todo = &todox.List[index]

			break
		}
	}

	if todo == nil {
		http.Error(w, "Not found", http.StatusNotFound)

		return
	}

	tmpl := template.New("updated")
	tmpl, err := tmpl.ParseFS(todox.Templates, "todo/todo.html")
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
