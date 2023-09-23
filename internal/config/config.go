package config

import "github.com/leapkit/core/envor"

var (
	// DatabaseURL is the connection string for the database
	// that will be used by the application.
	DatabaseURL = envor.Get("DATABASE_URL", "./todox.db")
	Environment = envor.Get("GO_ENV", "development")
	Port        = envor.Get("PORT", "3000")

	SessionName   = envor.Get("SESSION_NAME", "todox_session")
	SessionSecret = envor.Get("SESSION_SECRET", "secret_key")

	GlovesExtensionsToWatch = []string{".go", ".html", ".css", ".js"}
	GlovesExcludePaths      = []string{""}
)
