package todo

import (
	"net/http"
	"paganotoni/todox"
	"paganotoni/todox/internal"

	"github.com/gofrs/uuid"
)

func Create(w http.ResponseWriter, r *http.Request) {
	todo := todox.Todo{
		ID:        uuid.Must(uuid.NewV4()),
		Content:   r.FormValue("content"),
		Completed: false,
	}

	// Append new TODO to the list
	todox.List = append(todox.List, todo)

	err := internal.Render(w, "todo", todo, "todo/todo.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
