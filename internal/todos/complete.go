package todos

import (
	"net/http"

	"github.com/gofrs/uuid/v5"
	"github.com/leapkit/leapkit/core/form"
)

func Complete(w http.ResponseWriter, r *http.Request) {
	todos := r.Context().Value("todoService").(*service)

	id := uuid.FromStringOrNil(r.PathValue("id"))
	todo, err := todos.Find(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)

		return
	}

	err = form.Decode(r, &todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = todos.SetCompleted(todo.ID, todo.Completed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	todoHTML(todo).Render(w)
}
