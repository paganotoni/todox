package internal

import (
	"github.com/leapkit/core/db"
	_ "github.com/mattn/go-sqlite3"
)

// Connection is the database connection builder function
// that will be used by the application based on the driver and
// connection string.
var Connection = db.ConnectionFn(DatabaseURL, db.WithDriver("sqlite3"))
