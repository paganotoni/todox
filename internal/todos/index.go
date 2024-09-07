package todos

import (
	"net/http"

	. "github.com/delaneyj/gostar/elements"
)

func Index(w http.ResponseWriter, r *http.Request) {
	todos := r.Context().Value("todoService").(*service)

	list, err := todos.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	page(
		Group(
			INPUT().ID("search").TYPE("search").NAME("keyword").Attr("hx-get", "/todos/search").Attr("hx-target", "#todoList").Attr("hx-swap", "innerHTML").Attr("hx-trigger", "keyup delay:200ms").PLACEHOLDER("Type in to search").CLASS("p-2 px-3 mb-2 w-full rounded border"),
			HR().CLASS("border mb-2"),
			UL().ID("todoList").CLASS("flex flex-col gap-2 mb-2").Children(
				todoListHTML(list),
			),
			DIV().CLASS("p-3 border bg-white items-center rounded gap-2").Children(
				FORM().Attr("hx-post", "/todos/").Attr("hx-target", "#todoList").CLASS("flex flex-row gap-2 mb-0").Children(
					INPUT().TYPE("text").ID("content").NAME("Content").CLASS("p-2 border rounded flex-grow").PLACEHOLDER("TODO Content").AUTOFOCUS(),
					BUTTON().CLASS("bg-blue-500 rounded p-2 px-4 text-white").Text("Create"),
				),
			).Attr("_", "on htmx:afterRequest if detail.successful tell [#content,#search] set you.value to ''"),
		),
	).Render(w)

}
