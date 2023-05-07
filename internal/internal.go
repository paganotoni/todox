package internal

import (
	"net/http"
	"paganotoni/todox"
	"text/template"
)

// internal.Render renders a template with the given data considering the templates FS.
func Render(w http.ResponseWriter, partial string, data any, templates ...string) error {
	tmpl := template.New("xx")
	tmpl, err := tmpl.ParseFS(todox.Templates, templates...)
	if err != nil {
		return err
	}

	err = tmpl.ExecuteTemplate(w, partial, data)
	if err != nil {
		return err
	}

	return nil
}
