package todos

import (
	"net/http"

	"go.leapkit.dev/core/server"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	hx "maragu.dev/gomponents-htmx"
)

func Index(w http.ResponseWriter, r *http.Request) {
	todos := r.Context().Value("todoService").(*service)

	list, err := todos.List()
	if err != nil {
		server.Error(w, err, http.StatusInternalServerError)

		return
	}

	p := page(
		Div(
			Input(
				Class("p-2 px-3 mb-2 w-full rounded border"),
				ID("search"), Type("search"), Name("keyword"),
				hx.Trigger("keyup delay:200ms"), hx.Get("/todos/search"), hx.Target("#todoList"), hx.Swap("innerHTML"),
				Placeholder("Type in to search"),
			),

			Hr(Class("border mb-2")),
			Ul(
				ID("todoList"),
				Class("flex flex-col gap-2 mb-2"),
				todoListHTML(list),
			),

			Div(
				Class("p-3 border bg-white items-center rounded gap-2"),
				Attr("_", "on htmx:afterRequest if detail.successful tell [#content,#search] set you.value to ''"),
				Form(
					Class("flex flex-row gap-2 mb-0 items-center"),
					hx.Post("/todos/"), hx.Target("#todoList"),
					Input(Type("text"), ID("content"), Name("Content"), Class("p-2 border rounded flex-grow"), Placeholder("TODO Content"), Attr("autofocus")),
					Button(Class("bg-blue-500 rounded p-2 px-4 text-white"), Text("Create")),
				),
			),
		),
	)

	p.Render(w)
}
