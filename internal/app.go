package internal

import (
	"cmp"
	"net/http"
	"os"

	"todox/internal/todos"
	"todox/public"

	"go.leapkit.dev/core/db"
	"go.leapkit.dev/core/server"
)

// DBFn is the database connection builder function
// that will be used by the application based on the driver and
// connection string.
var DBFn = db.ConnectionFn(
	cmp.Or(os.Getenv("DATABASE_URL"), "database.db"),
	db.WithDriver("sqlite3"),

	db.Params(
		"_busy_timeout", "5000",
		"_journal_mode", "WAL",
		"_sync", "1",
		"_cache_size", "8192",
		"_txlock", "deferred",
		// journal_size_limit
		// mmap_size
	),
)

// New mounts the routes for the application,
// it assumes that the base services have been injected
// in the creation of the server instance.
func New() (string, http.Handler) {
	r := server.New(
		server.WithHost(cmp.Or(os.Getenv("HOST"), "0.0.0.0")),
		server.WithPort(cmp.Or(os.Getenv("PORT"), "3000")),
		server.WithAssets(public.Files, "/public/"),

		server.WithSession(
			cmp.Or(os.Getenv("SESSION_SECRET"), "secret_key"),
			cmp.Or(os.Getenv("SESSION_NAME"), "todox_session"),
		),
	)

	// Inject the todoService into the context
	r.Use(server.InCtxMiddleware("todoService", todos.NewService(DBFn)))

	r.HandleFunc("GET /{$}", todos.Index)
	r.HandleFunc("GET /health", health)

	r.Group("/todos", func(r server.Router) {
		r.HandleFunc("GET /search", todos.SearchItem)
		r.HandleFunc("POST /{$}", todos.Create)
		r.HandleFunc("GET /{id}/edit", todos.Edit)
		r.HandleFunc("GET /{id}/show", todos.Show)
		r.HandleFunc("DELETE /{id}/{$}", todos.Delete)
		r.HandleFunc("PUT /{id}/{$}", todos.Update)
		r.HandleFunc("PUT /{id}/complete", todos.Complete)
	})

	return r.Addr(), r.Handler()
}
