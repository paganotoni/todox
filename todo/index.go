package todo

import (
	"net/http"
	"paganotoni/todox/database"
	"paganotoni/todox/internal"
)

func Index(w http.ResponseWriter, r *http.Request) {
	conn := database.FromContext(r.Context())
	list, err := list(conn)
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
