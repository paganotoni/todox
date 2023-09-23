package web

import (
	"embed"
	"net/http"
	"paganotoni/todox/internal/config"
	"paganotoni/todox/internal/helpers"
	"paganotoni/todox/internal/web/public"
	"paganotoni/todox/internal/web/todos"
	"path"

	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/sessions"
	"github.com/leapkit/core/mdfs"
	"github.com/leapkit/core/render"
	"github.com/leapkit/core/server"
	"github.com/leapkit/core/session"
)

var (
	//go:embed **/*.html
	tmpls     embed.FS
	templates = mdfs.New(
		tmpls,
		path.Join("internal", "web"),
		config.Environment,
	)

	//sessions store
	store = sessions.NewCookieStore([]byte(config.SessionSecret))

	// the rendering engine for the application, this
	// is used to render each of the HTML responses
	// for the application.
	renderer = render.NewEngine(
		templates,
		render.WithHelpers(helpers.All),
	)
)

func Routes(r *server.Instance) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// LeapKit Middleware
	r.Use(session.InCtx(store, config.SessionName))
	r.Use(render.InCtx(renderer))
	r.Use(session.AddHelpers)

	r.Get("/", todos.Index)
	r.Get("/search", todos.Search)
	r.Get("/{id}/edit", todos.Edit)
	r.Post("/", todos.Create)
	r.Delete("/{id}", todos.Delete)
	r.Put("/{id}", todos.Update)
	r.Put("/{id}/complete", todos.Complete)

	// Public files that include anything thats on the
	// public folder. This is useful for files and assets.
	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.FS(public.Folder))))
}
