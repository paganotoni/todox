package todos

import (
	"net/http"

	"github.com/leapkit/core/form"
)

func Create(w http.ResponseWriter, r *http.Request) {
	todo := Instance{}
	err := form.Decode(r, &todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	todo.Completed = false
	todos := r.Context().Value("todoService").(*service)

	err = todos.Create(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	list, err := todos.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	todoListHTML(list).Render(w)
}
