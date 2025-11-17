package todos

import (
	"net/http"

	"github.com/gofrs/uuid/v5"
	"go.leapkit.dev/core/server"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	hx "maragu.dev/gomponents-htmx"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	todos := r.Context().Value("todoService").(*service)

	id := uuid.FromStringOrNil(r.PathValue("id"))
	todo, err := todos.Find(id)
	if err != nil {
		server.Error(w, err, http.StatusInternalServerError)

		return
	}

	el := Li(
		Class("gap-3 bg-white items-center rounded gap-2 p-3"),

		hx.Get("/todos/"+todo.ID.String()+"/show"),
		hx.Swap("outerHTML"),
		hx.Trigger("keyup[event.keyCode==27] from:window"),

		Form(
			Class("flex flex-row gap-3 mb-0"),

			hx.Put("/todos/"+todo.ID.String()),
			hx.Swap("outerHTML"),
			hx.Target("closest li"),

			Input(
				Value(todo.Content), Type("text"), Name("content"), Class("p-2 border rounded flex-grow"),
			),
			Button(
				Class("p-2 px-3 bg-green-500 text-white rounded"),
				Text("Save"),
			),
		),
	)

	el.Render(w)
}
