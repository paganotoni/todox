package main

import (
	"cmp"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"todox/internal"

	"github.com/leapkit/core/server"

	// Load environment variables
	_ "github.com/leapkit/core/tools/envload"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := internal.DB()
	if err != nil {
		panic(err)
	}

	server := server.New(
		server.WithHost(cmp.Or(os.Getenv("HOST"), "0.0.0.0")),
		server.WithPort(cmp.Or(os.Getenv("PORT"), "3000")),
	)

	// Application routes
	if err := internal.SetupRoutes(server, db); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	slog.Info(fmt.Sprintf("> Starting todox at %s", server.Addr()))
	if err := http.ListenAndServe(server.Addr(), server.Handler()); err != nil {
		slog.Error(fmt.Sprintf("Server terminated: %v", err.Error()))
	}
}
