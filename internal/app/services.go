package app

import (
	"todox/internal/app/database"
	"todox/internal/todos"

	"github.com/leapkit/core/server"
)

// AddServices is a function that will be called by the server
// to inject services in the context.
func AddServices(r *server.Instance) error {
	conn, err := database.Connection()
	if err != nil {
		return err
	}

	// Services that will be injected in the context
	r.Use(server.InCtxMiddleware("todoService", todos.NewService(conn)))

	return nil
}
