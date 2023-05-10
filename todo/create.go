package todo

import (
	"net/http"
	"paganotoni/todox"
	"paganotoni/todox/internal"
	"paganotoni/todox/internal/database"

	"github.com/gofrs/uuid"
)

func Create(w http.ResponseWriter, r *http.Request) {
	todo := todox.Todo{
		ID:        uuid.Must(uuid.NewV4()),
		Content:   r.FormValue("content"),
		Completed: false,
	}

	conn := database.FromContext(r.Context())
	err := create(conn, todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	list, err := list(conn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	err = internal.Render(w, "list", list, "todo/index.html", "todo/todo.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
