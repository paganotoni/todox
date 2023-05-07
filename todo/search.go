package todo

import (
	"net/http"
	"paganotoni/todox"
	"paganotoni/todox/internal"
	"strings"
)

func Search(w http.ResponseWriter, r *http.Request) {
	term := r.FormValue("keyword")
	todos := []todox.Todo{}
	for _, tx := range todox.List {
		if strings.Contains(tx.Content, term) {
			todos = append(todos, tx)
		}
	}

	err := internal.Render(w, "list", todos, "todo/index.html", "todo/todo.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
