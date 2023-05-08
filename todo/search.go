package todo

import (
	"net/http"
	"paganotoni/todox/database"
	"paganotoni/todox/internal"
)

func Search(w http.ResponseWriter, r *http.Request) {
	conn := database.FromContext(r.Context())

	list, err := search(conn, r.FormValue("keyword"))
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
