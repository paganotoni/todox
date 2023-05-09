package main

import (
	"fmt"
	"os"
	"paganotoni/todox/database"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		help()
		return
	}

	switch args[1] {
	case "migrate":
		fmt.Println("ℹ️ Running migrations")
		err := database.Migrate()

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("✅ Migrations complete")

		return
	case "generate":
		if len(args) < 4 {
			fmt.Println("Please provide a valid command: generate migration <name>")
		}

		//TODO: Generate migration.

	default:
		help()
	}
}

func help() {
	fmt.Println("Please provide a valid command: migrate")
}
