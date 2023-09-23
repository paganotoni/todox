package sqlite

import (
	"embed"

	"paganotoni/todox/internal/config"

	"github.com/leapkit/core/db"
	_ "github.com/mattn/go-sqlite3"
)

var (
	//go:embed migrations/*.sql
	Migrations embed.FS

	//Connection is the database connection builder function
	//that will be used by the application based on the driver and
	//connection string.
	Connection = db.ConnectionFn(config.DatabaseURL, db.WithDriver("sqlite3"))
)
