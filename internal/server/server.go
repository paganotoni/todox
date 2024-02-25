package server

import (
	"log/slog"
	"net/http"
)

// Rood routeGroup is a group of routes with a common prefix and middleware
// it also has a host and port as well as a Start method as it is the root of the server
// that should be executed for all the handlers in the group.
type Root struct {
	*RouteGroup

	host string
	port string
}

func New(options ...Option) *Root {
	ss := &Root{
		RouteGroup: &RouteGroup{
			prefix:     "",
			mux:        http.NewServeMux(),
			middleware: []func(http.Handler) http.Handler{},
		},

		host: "0.0.0.0",
		port: "3000",
	}

	ss.Use(logger)
	ss.Use(recoverer)
	ss.Use(requestID)
	ss.Use(setValuer)

	for _, option := range options {
		option(ss)
	}

	return ss
}

func (s Root) Start() error {
	slog.Info("> Starting server on port " + s.port)

	fhp := s.host + ":" + s.port
	return http.ListenAndServe(fhp, s.mux)
}
