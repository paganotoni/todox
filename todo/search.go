package todo

import (
	"net/http"
	"paganotoni/todox"
	"paganotoni/todox/database"
	"paganotoni/todox/internal"
)

func Search(w http.ResponseWriter, r *http.Request) {
	term := r.FormValue("keyword")
	var list []todox.Todo
	conn := database.FromContext(r.Context())

	err := conn.Select(&list, "SELECT * FROM todos WHERE content LIKE $1", "%"+term+"%")
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
