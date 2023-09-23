package config

var (
	// DatabaseURL is the connection string for the database
	// that will be used by the application.
	DatabaseURL             = "./todox.db"
	Environment             = "development"
	SessionName             = "todox"
	SessionSecret           = "secret"
	GlovesExtensionsToWatch = []string{".go", ".html", ".css", ".js"}
)
