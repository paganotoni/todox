package main

import (
	"fmt"
	"net/http"
	"os"
	"paganotoni/todox/database"
	"paganotoni/todox/internal/envor"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		err := database.Migrate()
		if err != nil {
			fmt.Println(err)
		}

		return
	}

	// Start the server
	addr := ":" + envor.Get("PORT", "3000")
	fmt.Println("Server listening on", addr)
	http.ListenAndServe(addr, buildServer())
}
