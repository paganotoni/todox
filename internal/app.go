package internal

import (
	"cmp"
	"os"
	"todox/internal/todos"
	"todox/public"

	"github.com/jmoiron/sqlx"
	"github.com/leapkit/core/assets"
	"github.com/leapkit/core/db"
	"github.com/leapkit/core/server"
	"github.com/leapkit/core/tools/rebuilder"
	"github.com/paganotoni/tailo"

	"github.com/leapkit/core/session"
)

var (
	// DatabaseURL is the connection string for the database
	// that will be used by the application.
	DatabaseURL = cmp.Or(os.Getenv("DATABASE_URL"), "./todox.db")

	// DB is the database connection builder function
	// that will be used by the application based on the driver and
	// connection string.
	DB = db.ConnectionFn(DatabaseURL, db.WithDriver("sqlite3"))

	// Assets is the manager for the public assets
	// it allows to watch for changes and reload the assets
	// when changes are made.
	Assets = assets.NewManager(public.Files)

	// TailoOptions allow to define how to compile
	// the tailwind css files, which is the input and
	// what will be the output.
	TailoOptions = []tailo.Option{
		tailo.UseInputPath("internal/assets/application.css"),
		tailo.UseOutputPath("public/application.css"),
		tailo.UseConfigPath("tailwind.config.js"),
	}

	// GlovesOptions are the options that will be used by the gloves
	// tool to hot reload the application.
	GlovesOptions = []rebuilder.Option{
		// Run the tailo watcher so when changes are made to
		// the html code it rebuilds css.
		rebuilder.WithRunner(tailo.WatcherFn(TailoOptions...)),
		rebuilder.WithRunner(Assets.Watch),
		rebuilder.WatchExtension(".go", ".css", ".js"),
	}
)

// AddRoutes mounts the routes for the application,
// it assumes that the base services have been injected
// in the creation of the server instance.
func SetupRoutes(r server.Router, db *sqlx.DB) error {
	// Session middleware to be used by the application
	// to store session data.
	ssecret := cmp.Or(os.Getenv("SESSION_SECRET"), "secret_key")
	sname := cmp.Or(os.Getenv("SESSION_NAME"), "todox_session")
	r.Use(session.Middleware(ssecret, sname))

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
	r.HandleFunc(Assets.HandlerPattern(), Assets.HandlerFn)

	return nil
}
