package database

import (
	"sync"

	"github.com/jmoiron/sqlx"
)

var (
	conn *sqlx.DB
	mx   sync.Mutex
)

// Connection returns a database connection to the database
// it reads from the DATABASE_URL variable and the DATABASE_DRIVER
// environment variables to determine the connection string and
// the driver to use. If DATABASE_URL is not set it will use the
// application name, if DATABASE_DRIVER is not set it uses
// sqlite3 as default.
func connection() (*sqlx.DB, error) {
	mx.Lock()
	defer mx.Unlock()

	if conn != nil {
		return conn, nil
	}

	return sqlx.Connect("sqlite3", "todox.db")
}
