package todos

import . "github.com/delaneyj/gostar/elements"

func page(e ElementRenderer) ElementRenderer {
	return HTML().Children(
		HEAD().Children(
			META().NAME("viewport").CONTENT("width=device-width, initial-scale=1"),
			META().CHARSET("utf-8"),
			TITLE().Text("Todo"),
			SCRIPT().SRC("https://unpkg.com/htmx.org@1.9.2"),
			SCRIPT().SRC("https://unpkg.com/hyperscript.org@0.9.8"),
			LINK().REL("stylesheet").HREF("/public/application.css"),
		),
		BODY().CLASS("h-full bg-gray-100 pb-10 pt-10").Children(
			DIV().CLASS("max-w-[1500px] mx-auto px-5").Children(
				H1().CLASS("text-2xl mb-2 font-bold").Text("Todo List"),
				e,
			),
		),
	)
}
