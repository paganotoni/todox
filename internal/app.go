package internal

import (
	"cmp"
	"net/http"
	"os"
	"todox/internal/todos"
	"todox/public"

	"github.com/leapkit/leapkit/core/assets"
	"github.com/leapkit/leapkit/core/db"
	"github.com/leapkit/leapkit/core/server"

	"github.com/leapkit/leapkit/core/session"
)

var (
	// DB is the database connection builder function
	// that will be used by the application based on the driver and
	// connection string.
	DB = db.ConnectionFn(
		cmp.Or(os.Getenv("DATABASE_URL"), "./todox.db"),
		db.WithDriver("sqlite3"),
	)
)

type Server interface {
	Addr() string
	Handler() http.Handler
}

// AddRoutes mounts the routes for the application,
// it assumes that the base services have been injected
// in the creation of the server instance.
func New() Server {
	r := server.New(
		server.WithHost(cmp.Or(os.Getenv("HOST"), "0.0.0.0")),
		server.WithPort(cmp.Or(os.Getenv("PORT"), "3000")),
	)

	// Session middleware to be used by the application
	// to store session data.
	r.Use(session.Middleware(
		cmp.Or(os.Getenv("SESSION_SECRET"), "secret_key"),
		cmp.Or(os.Getenv("SESSION_NAME"), "todox_session"),
	))

	// Inject the todoService into the context
	r.Use(server.InCtxMiddleware("todoService", todos.NewService(DB)))

	r.HandleFunc("GET /{$}", todos.Index)
	r.HandleFunc("GET /search", todos.Search)
	r.HandleFunc("POST /{$}", todos.Create)

	r.Group("/{id}/", func(wid server.Router) {
		wid.HandleFunc("GET /edit", todos.Edit)
		wid.HandleFunc("GET /show", todos.Show)
		wid.HandleFunc("DELETE /{$}", todos.Delete)
		wid.HandleFunc("PUT /{$}", todos.Update)
		wid.HandleFunc("PUT /complete", todos.Complete)
	})

	// Mounting the assets manager at the end of the routes
	// so that it can serve the public assets.
	assetManager := assets.NewManager(public.Files)
	r.HandleFunc(assetManager.HandlerPattern(), assetManager.HandlerFn)

	return r
}
