package todos

import (
	"net/http"

	"github.com/leapkit/leapkit/core/server"
)

func Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		server.Error(w, err, http.StatusInternalServerError)

		return
	}

	todo := Instance{
		Content: r.FormValue("Content"),
	}

	todo.Completed = false
	todos := r.Context().Value("todoService").(*service)

	err = todos.Create(&todo)
	if err != nil {
		server.Error(w, err, http.StatusInternalServerError)

		return
	}

	list, err := todos.List()
	if err != nil {
		server.Error(w, err, http.StatusInternalServerError)

		return
	}

	todoListHTML(list).Render(w)
}
