package main

import (
	"fmt"

	"todox/internal"
	"todox/internal/migrations"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"go.leapkit.dev/core/db"
)

// The migrate command is used to ship our application
// with the latest database schema migrator. which can be invoked
// by running `migrate`.
func main() {
	conn, err := internal.DB()
	if err != nil {
		fmt.Println("Error connecting to the database: ", err)
	}

	err = db.RunMigrations(migrations.All, conn)
	if err != nil {
		fmt.Println("Error running migrations: ", err)
	}
}
