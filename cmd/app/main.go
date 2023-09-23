package main

import (
	"fmt"
	"paganotoni/todox/internal/sqlite"
	"paganotoni/todox/internal/web"

	"github.com/leapkit/core/server"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	conn, err := sqlite.Connection()
	if err != nil {
		panic(err)
	}

	s := server.New(
		"Todox",
		// Services to be injected in the context.
		server.WithCtxVal("todoService", sqlite.NewTodoService(conn)),

		// Routes are defined in internal/web/routes.go
		// we pass these to the newly created server
		// as an option.
		server.WithRoutesFn(web.Routes),
	)

	if err := s.Start(); err != nil {
		fmt.Println(err)
	}
}
