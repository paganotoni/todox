package internal

import (
	"todox/internal/todos"

	"github.com/leapkit/core/server"

	lserver "github.com/leapkit/core/server"
)

// AddServices is a function that will be called by the server
// to inject services in the context.
func AddServices(r server.Router) error {
	conn, err := DB()
	if err != nil {
		return err
	}

	// Services that will be injected in the context
	r.Use(lserver.InCtxMiddleware("todoService", todos.NewService(conn)))

	return nil
}
