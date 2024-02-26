package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"todox/internal"

	"github.com/leapkit/core/envor"
	"github.com/leapkit/core/server"
)

func main() {
	server := server.New(
		server.WithHost(envor.Get("HOST", "0.0.0.0")),
		server.WithPort(envor.Get("PORT", "3000")),
	)

	// Application services
	if err := internal.AddServices(server); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	// Application routes
	if err := internal.AddRoutes(server); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	slog.Info(fmt.Sprintf("> Starting todox at %s", server.Addr()))
	if err := http.ListenAndServe(server.Addr(), server.Handler()); err != nil {
		slog.Error(fmt.Sprintf("Server terminated: %v", err.Error()))
	}
}
