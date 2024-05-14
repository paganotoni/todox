package todos

import . "github.com/delaneyj/gostar/elements"

func todoHTML(t Instance) ElementRenderer {
	bgClass := "bg-gray-50"
	content := SPAN().CLASS("line-through").Text(t.Content)
	if !t.Completed {
		bgClass = "bg-white"
		content = SPAN().Attr("hx-get", "/"+t.ID.String()+"/edit").Attr("hx-target", "closest li").Attr("hx-swap", "outerHTML").CLASS("cursor-pointer underline flex flex-row gap-2 items-center").Attr("_", "on htmx:afterRequest if detail.successful tell #content set you.value to ''").Children(
			SPAN().Text(t.Content),
			//TODO Icon
		)
	}

	return LI().CLASS("p-3 border  flex flex-row items-center rounded gap-2", bgClass).Children(
		SPAN().CLASS("flex h-6 items-center").Children(
			DIV(
				INPUT().NAME("Completed").CHECKEDSet(t.Completed).TYPE("checkbox").ID("complete_"+t.ID.String()).Attr("hx-put", "/"+t.ID.String()+"/complete").Attr("hx-include", "#not_complete_"+t.ID.String()).Attr("hx-target", "closest li").Attr("hx-swap", "outerHTML").Attr("aria-describedby", "comments-description").CLASS("h-5 w-5 rounded border-gray-300 text-indigo-600 focus:ring-indigo-600"),
				INPUT().NAME("Completed").TYPE("hidden").VALUE("false").ID("not_complete_"+t.ID.String()),
			),
		),
		SPAN().CLASS("flex-grow items-center").Children(content),
		SPAN().CLASS("bg-red-500 text-white py-2 px-3 rounded cursor-pointer").Text("Delete").Attr("hx-delete", "/"+t.ID.String()).Attr("hx-target", "#element_"+t.ID.String()).Attr("hx-swap", "outerHTML").Attr("hx-confirm", "Are you sure you wish to delete this TODO?").Attr("_", "on htmx:afterRequest if detail.successful send keyup to #search"),
	)
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
