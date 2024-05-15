package todos

import . "github.com/delaneyj/gostar/elements"

func todoHTML(t Instance) ElementRenderer {
	return LI().CLASS("p-3 border  flex flex-row items-center rounded gap-2").Children(
		SPAN().CLASS("flex h-6 items-center").Children(
			DIV(
				INPUT().NAME("Completed").CHECKEDSet(t.Completed).TYPE("checkbox").ID("complete_"+t.ID.String()).Attr("hx-put", "/"+t.ID.String()+"/complete").Attr("hx-include", "#not_complete_"+t.ID.String()).Attr("hx-target", "closest li").Attr("hx-swap", "outerHTML").Attr("aria-describedby", "comments-description").CLASS("h-5 w-5 rounded border-gray-300 text-indigo-600 focus:ring-indigo-600"),
				INPUT().NAME("Completed").TYPE("hidden").VALUE("false").ID("not_complete_"+t.ID.String()),
			),
		),

		SPAN().CLASS("flex-grow items-center").IfChildren(
			t.Completed,
			SPAN().CLASS("line-through").Text(t.Content),
		).IfChildren(
			!t.Completed,
			SPAN().Attr("hx-get", "/"+t.ID.String()+"/edit").Attr("hx-target", "closest li").Attr("hx-swap", "outerHTML").CLASS("cursor-pointer underline flex flex-row gap-2 items-center").Attr("_", "on htmx:afterRequest if detail.successful tell #content set you.value to ''").Children(
				SPAN().Text(t.Content),
				//TODO Icon
				// <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
				//                 <path stroke-linecap="round" stroke-linejoin="round" d="M16.862 4.487l1.687-1.688a1.875 1.875 0 112.652 2.652L6.832 19.82a4.5 4.5 0 01-1.897 1.13l-2.685.8.8-2.685a4.5 4.5 0 011.13-1.897L16.863 4.487zm0 0L19.5 7.125" />
				//             </svg>
				//
				SVG_SVG().Attr("xmlns", "http://www.w3.org/2000/svg").Attr("fill", "none").Attr("viewBox", "0 0 24 24").Attr("stroke-width", "1.5").Attr("stroke", "currentColor").CLASS("w-4 h-4").Children(
					SVG_PATH().Attr("stroke-linecap", "round").Attr("stroke-linejoin", "round").Attr("d", "M16.862 4.487l1.687-1.688a1.875 1.875 0 112.652 2.652L6.832 19.82a4.5 4.5 0 01-1.897 1.13l-2.685.8.8-2.685a4.5 4.5 0 011.13-1.897L16.863 4.487zm0 0L19.5 7.125"),
				),
			),
		),

		SPAN().CLASS("bg-red-500 text-white py-2 px-3 rounded cursor-pointer").Text("Delete").Attr("hx-delete", "/"+t.ID.String()).Attr("hx-target", "closest li").Attr("hx-swap", "outerHTML").Attr("hx-confirm", "Are you sure you wish to delete this TODO?").Attr("_", "on htmx:afterRequest if detail.successful send keyup to #search"),
	).IfCLASS(t.Completed, "bg-gray-50").IfCLASS(!t.Completed, "bg-white")
}

func todoListHTML(todos []Instance) ElementRenderer {
	if len(todos) == 0 {
		return LI().CLASS("p-10 bg-white rounded border text-center").Text("No todos found.")
	}

	els := []ElementRenderer{}
	for _, todo := range todos {
		els = append(els, todoHTML(todo))
	}

	return Group(els...)
}
