package todos

import (
	"net/http"

	"github.com/gofrs/uuid/v5"

	. "github.com/delaneyj/gostar/elements"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	todos := r.Context().Value("todoService").(*service)

	id := uuid.FromStringOrNil(r.PathValue("id"))
	todo, err := todos.Find(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	el := LI().CLASS("gap-3 bg-white items-center rounded gap-2 p-3").Children(
		FORM().CLASS("flex flex-row gap-3 mb-0").Children(
			INPUT().VALUE(todo.Content).TYPE("text").NAME("content").CLASS("p-2 border rounded flex-grow"),
			BUTTON().CLASS("p-2 px-3 bg-green-500 text-white rounded").Text("Save"),
		).Attr("hx-put", "/todos/"+todo.ID.String()).Attr("hx-swap", "outerHTML").Attr("hx-target", "closest li"),
	).Attr("hx-get", "/todos/"+todo.ID.String()+"/show").Attr("hx-swap", "outerHTML").Attr("hx-trigger", "keyup[event.keyCode==27] from:window")
	el.Render(w)
}
