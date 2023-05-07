package todo

import (
	"net/http"
	"paganotoni/todox"
	"paganotoni/todox/internal"
)

func Index(w http.ResponseWriter, r *http.Request) {
	templates := []string{
		"todo/index.html",
		"todo/todo.html",
		"application.html",
	}

	err := internal.Render(w, "page", todox.List, templates...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
