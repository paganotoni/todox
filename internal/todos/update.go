package todos

import (
	"net/http"

	"github.com/gofrs/uuid/v5"
	"github.com/leapkit/leapkit/core/server"
)

func Update(w http.ResponseWriter, r *http.Request) {
	todos := r.Context().Value("todoService").(*service)

	id := uuid.FromStringOrNil(r.PathValue("id"))
	todo, err := todos.Find(id)
	if err != nil {
		server.Error(w, err, http.StatusInternalServerError)

		return
	}

	r.ParseForm()
	todo.Content = r.FormValue("content")

	err = todos.Update(&todo)
	if err != nil {
		server.Error(w, err, http.StatusInternalServerError)

		return
	}

	todoHTML(todo).Render(w)
}
