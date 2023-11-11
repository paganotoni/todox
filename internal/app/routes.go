package app

import (
	"paganotoni/todox/internal"
	"paganotoni/todox/internal/app/config"
	"paganotoni/todox/internal/app/public"
	"paganotoni/todox/internal/todos"

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

	rengine := render.NewEngine(internal.Templates)
	r.Use(render.Middleware(rengine))
	r.Use(session.Middleware(config.SessionSecret, config.SessionName))

	// Todo actions
	r.Get("/", todos.Index)
	r.Get("/search", todos.Search)
	r.Get("/{id}/edit", todos.Edit)
	r.Post("/", todos.Create)
	r.Delete("/{id}", todos.Delete)
	r.Put("/{id}", todos.Update)
	r.Put("/{id}/complete", todos.Complete)

	// Mount the public folder to be served openly
	r.Folder("/public/", public.Folder)

	return nil
}
