package todos

import (
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	todos := r.Context().Value("todoService").(*service)

	list, err := todos.Search(r.FormValue("keyword"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	html := todoListHTML(list)
	html.Render(w)
}
