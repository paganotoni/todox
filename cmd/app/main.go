package main

import (
	"fmt"
	"os"
	"paganotoni/todox/internal/app"
	"paganotoni/todox/internal/app/config"

	"github.com/leapkit/core/server"
)

func main() {
	s := server.New(
		"Todox",

		server.WithPort(config.Port),
		server.WithHost(config.Host),
	)

	// Application services
	if err := app.AddServices(s); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Application routes
	if err := app.AddRoutes(s); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := s.Start(); err != nil {
		fmt.Println(err)
	}
}
