package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"todox/internal"

	// Load environment variables
	_ "github.com/leapkit/leapkit/core/tools/envload"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

var server = internal.New()

func main() {
	slog.Info(fmt.Sprintf("> Starting todox at %s", server.Addr()))
	err := http.ListenAndServe(server.Addr(), server.Handler())
	if err != nil {
		slog.Error(fmt.Sprintf("Server terminated: %v", err.Error()))
	}
}
