package todos

import (
	"net/http"

	"go.leapkit.dev/core/server"
)

func Search(w http.ResponseWriter, r *http.Request) {
	todos := r.Context().Value("todoService").(*service)

	list, err := todos.Search(r.FormValue("keyword"))
	if err != nil {
		server.Error(w, err, http.StatusInternalServerError)

		return
	}

	html := todoListHTML(list)
	html.Render(w)
}
