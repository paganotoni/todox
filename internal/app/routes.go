package app

import (
	"todox/internal"
	"todox/internal/app/config"
	"todox/internal/app/public"
	"todox/internal/todos"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/leapkit/core/render"
	"github.com/leapkit/core/server"
	"github.com/leapkit/core/session"
)

// AddRoutes mounts the routes for the application,
// it assumes that the base services have been injected
// in the creation of the server instance.
func AddRoutes(r *server.Instance) error {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(render.Middleware(render.NewEngine(internal.Templates)))
	r.Use(session.Middleware(config.SessionSecret, config.SessionName))

	// Todo actions
	r.Get("/", todos.Index)
	r.Get("/search", todos.Search)
	r.Get("/{id}/edit", todos.Edit)
	r.Get("/{id}/show", todos.Show)
	r.Post("/", todos.Create)
	r.Delete("/{id}", todos.Delete)
	r.Put("/{id}", todos.Update)
	r.Put("/{id}/complete", todos.Complete)

	// Mount the public folder to be served openly
	r.Folder("/public/", public.Folder)

	return nil
}
