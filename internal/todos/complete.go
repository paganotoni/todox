package todos

import (
	"net/http"

	"github.com/gofrs/uuid/v5"
	"github.com/leapkit/leapkit/core/server"
)

func Complete(w http.ResponseWriter, r *http.Request) {
	todos := r.Context().Value("todoService").(*service)

	id := uuid.FromStringOrNil(r.PathValue("id"))
	todo, err := todos.Find(id)
	if err != nil {
		server.Error(w, err, http.StatusInternalServerError)

		return
	}

	todo.Completed = r.FormValue("Completed") != "false"
	err = todos.SetCompleted(todo.ID, todo.Completed)
	if err != nil {
		server.Error(w, err, http.StatusInternalServerError)

		return
	}

	todoHTML(todo).Render(w)
}
