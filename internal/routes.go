package internal

import (
	"todox/internal/server"
	"todox/internal/todos"
	"todox/public"

	"github.com/leapkit/core/envor"
	"github.com/leapkit/core/render"
	"github.com/leapkit/core/session"
)

var (
	// Session options.
	sessionName   = envor.Get("SESSION_NAME", "todox_session")
	sessionSecret = envor.Get("SESSION_SECRET", "secret_key")
)

// AddRoutes mounts the routes for the application,
// it assumes that the base services have been injected
// in the creation of the server instance.
func AddRoutes(r *server.Root) error {
	r.Use(session.Middleware(sessionSecret, sessionName))
	r.Use(render.Middleware(Templates, render.WithDefaultLayout("layout.html")))

	r.HandleFunc("GET /", todos.Index)
	r.HandleFunc("GET /search", todos.Search)
	r.HandleFunc("GET /{id}/edit", todos.Edit)
	r.HandleFunc("GET /{id}/show", todos.Show)
	r.HandleFunc("POST /", todos.Create)
	r.HandleFunc("DELETE /{id}", todos.Delete)
	r.HandleFunc("PUT /{id}", todos.Update)
	r.HandleFunc("PUT /{id}/complete", todos.Complete)

	// Mount the public folder to be served openly
	r.Folder("/public", public.Files)

	return nil
}
