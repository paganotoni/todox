package todos

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	hx "maragu.dev/gomponents-htmx"
)

func todoHTML(t Instance) Node {
	return Li(
		Class("bg-white p-3 border flex flex-row items-center rounded gap-2"),
		If(t.Completed, Class("bg-gray-50")),
		If(!t.Completed, Class("bg-white")),

		Span(
			Class("flex h-6 items-center"),
			Div(
				Input(
					Name("Completed"), If(t.Completed, Checked()), Type("checkbox"), ID("complete_"+t.ID.String()),
					hx.Put("/todos/"+t.ID.String()+"/complete"),
					hx.Include("#not_complete_"+t.ID.String()),
					hx.Target("closest li"),
					hx.Swap("outerHTML"),

					Class("h-5 w-5 rounded border-gray-300 text-indigo-600 focus:ring-indigo-600"),
				),
				Input(
					Name("Completed"), Type("hidden"), Value("false"), ID("not_complete_"+t.ID.String()),
				),
			),
		),

		Span(
			Class("flex-grow items-center"),
			If(t.Completed, Span(Class("line-through"), Text(t.Content))),
			If(!t.Completed,
				Span(
					Class("cursor-pointer underline flex flex-row gap-2 items-center"),

					Attr("_", "on htmx:afterRequest if detail.successful tell #content set you.value to ''"),
					hx.Get("/todos/"+t.ID.String()+"/edit"),
					hx.Target("closest li"),
					hx.Swap("outerHTML"),

					Span(Text(t.Content)),
					Raw(
						`<svg xmlns="http://www.w3.org/2000/svg" height="18px" viewBox="0 -960 960 960" width="18px" fill="#5f6368">
							<path d="M200-200h57l391-391-57-57-391 391v57Zm-80 80v-170l528-527q12-11 26.5-17t30.5-6q16 0 31 6t26 18l55 56q12 11 17.5 26t5.5 30q0 16-5.5 30.5T817-647L290-120H120Zm640-584-56-56 56 56Zm-141 85-28-29 57 57-29-28Z"/>
						</svg>`,
					),
				),
			),
		),

		Span(
			Class("bg-red-500 text-white py-2 px-3 rounded cursor-pointer"),
			hx.Confirm("Are you sure you wish to delete this TODO?"),
			hx.Delete("/todos/"+t.ID.String()),
			hx.Swap("outerHTML"),
			hx.Target("closest li"),
			Text("Delete"),

			Attr("_", "on htmx:afterRequest if detail.successful send keyup to #search"),
		),
	)
}

func todoListHTML(todos []Instance) Group {
	if len(todos) == 0 {
		return []Node{
			Li(
				Class("p-10 bg-white rounded border text-center"),
				Text("No todos found."),
			),
		}
	}

	els := []Node{}
	for _, todo := range todos {
		els = append(els, todoHTML(todo))
	}

	return els
}
