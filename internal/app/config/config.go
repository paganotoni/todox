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

	SessionName   = envor.Get("SESSION_NAME", "todox_session")
	SessionSecret = envor.Get("SESSION_SECRET", "secret_key")

	// Tailo Options
	TailoOptions = []tailo.Option{
		tailo.UseInputPath("internal/assets/application.css"),
		tailo.UseOutputPath("internal/app/public/application.css"),
		tailo.UseConfigPath("tailwind.config.js"),
	}

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
