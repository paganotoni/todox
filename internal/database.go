package internal

import (
	"github.com/leapkit/core/db"
	"github.com/leapkit/core/envor"
	_ "github.com/mattn/go-sqlite3"
)

var (
	// DatabaseURL is the connection string for the database
	// that will be used by the application.
	DatabaseURL = envor.Get("DATABASE_URL", "./todox.db")

	// DB is the database connection builder function
	// that will be used by the application based on the driver and
	// connection string.
	DB = db.ConnectionFn(DatabaseURL, db.WithDriver("sqlite3"))
)
