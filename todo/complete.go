package todo

import (
	"html/template"
	"net/http"
	"paganotoni/todox"
	"paganotoni/todox/database"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
)

func Complete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id := uuid.FromStringOrNil(chi.URLParam(r, "id"))
	conn := database.FromContext(r.Context())

	todo, err := find(conn, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	todo.Completed = (r.FormValue("completed") == "on")

	if err = complete(conn, todo); err != nil {
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
