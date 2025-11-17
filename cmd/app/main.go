package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"todox/internal"

	// Load environment variables
	_ "go.leapkit.dev/core/tools/envload"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	addr, server := internal.New()
	slog.Info(fmt.Sprintf("> Starting todox at %s", addr))
	http.ListenAndServe(addr, server)
}
