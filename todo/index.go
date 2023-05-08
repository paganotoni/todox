package todo

import (
	"net/http"
	"paganotoni/todox"
	"paganotoni/todox/database"
	"paganotoni/todox/internal"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var list []todox.Todo
	conn := database.FromContext(r.Context())

	err := conn.Select(&list, "SELECT * FROM todos")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	templates := []string{
		"todo/index.html",
		"todo/todo.html",
		"application.html",
	}

	err = internal.Render(w, "page", list, templates...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
