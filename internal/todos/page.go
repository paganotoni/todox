package todos

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func page(yield Node) Node {
	return HTML(
		Head(
			Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
			Meta(Charset("utf-8")),

			TitleEl(Text("TodoX")),
			Script(Src("https://unpkg.com/htmx.org@2.0.0/dist/htmx.min.js")),
			Script(Src("https://unpkg.com/hyperscript.org@0.9.8")),
			Script(Src("https://cdn.tailwindcss.com?plugins=forms,typography,line-clamp")),
		),

		Body(
			Class("h-full bg-gray-100 pb-10 pt-10"),
			Div(
				Class("max-w-[1500px] mx-auto px-5"),
				H1(
					Class("text-2xl mb-2 font-bold"),
					Text("TodoX List"),
				),

				// here goes the thing you want to render.
				yield,
			),
		),
	)
}
