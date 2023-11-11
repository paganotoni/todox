package config

import (
	"github.com/leapkit/core/envor"
	"github.com/leapkit/core/gloves"
	"github.com/paganotoni/tailo"
)

var (
	// DatabaseURL is the connection string for the database
	// that will be used by the application.
	DatabaseURL = envor.Get("DATABASE_URL", "./todox.db")
	Environment = envor.Get("GO_ENV", "development")

	Port = envor.Get("PORT", "3000")
	Host = envor.Get("HOST", "0.0.0.0")

	// Session options.
	SessionName   = envor.Get("SESSION_NAME", "todox_session")
	SessionSecret = envor.Get("SESSION_SECRET", "secret_key")

	// TailoOptions allow to define how to compile
	// the tailwind css files, which is the input and
	// what will be the output.
	TailoOptions = []tailo.Option{
		tailo.UseInputPath("internal/assets/application.css"),
		tailo.UseOutputPath("internal/app/public/application.css"),
		tailo.UseConfigPath("tailwind.config.js"),
	}

	// GlovesOptions are the options that will be used by the gloves
	// tool to hot reload the application.
	GlovesOptions = []gloves.Option{
		// Run the tailo watcher so when changes are made to
		// the html code it rebuilds css.
		gloves.WithRunner(func() {
			tailo.Watch(TailoOptions...)
		}),

		gloves.WatchExtension(".go", ".html", ".css", ".js"),
		gloves.ExcludePaths(""), // Add paths to exclude here.
	}
)
