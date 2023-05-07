package todo

import (
	"net/http"
	"paganotoni/todox"

	"github.com/go-chi/chi/v5"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	for index, todo := range todox.List {
		if id == todo.ID.String() {
			todox.List = append(todox.List[:index], todox.List[index+1:]...)

			break
		}
	}
}
