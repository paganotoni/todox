package main

import (
	"fmt"
	"os"

	"paganotoni/todox/internal/app/config"
	"paganotoni/todox/internal/migrations"
	"paganotoni/todox/internal/sqlite"

	"github.com/leapkit/core/db"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tools database <command>")
		fmt.Println("Available commands:")
		fmt.Println(" - migrate")
		fmt.Println(" - create")
		fmt.Println(" - drop")

		return
	}

	switch os.Args[1] {
	case "migrate":
		conn, err := sqlite.Connection()
		if err != nil {
			fmt.Println(err)
			return
		}

		err = db.RunMigrations(migrations.All, conn)
		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Println("✅ Migrations ran successfully")
	case "create":
		err := db.Create(config.DatabaseURL)
		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Println("✅ Database created successfully")

	case "drop":
		err := db.Drop(config.DatabaseURL)
		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Println("✅ Database dropped successfully")

	default:
		fmt.Println("command not found")

		return
	}
}
