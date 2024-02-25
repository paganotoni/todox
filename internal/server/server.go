package server

import (
	"log/slog"
	"net/http"
)

type Root struct {
	*RouteGroup

	host string
	port string
}

func New(options ...Option) *Root {
	ss := &Root{
		RouteGroup: &RouteGroup{
			prefix: "",
			mux:    http.NewServeMux(),
			middleware: []func(http.Handler) http.Handler{
				setValuer,
				recoverer,
				logger,
			},
		},

		host: "0.0.0.0",
		port: "3000",
	}

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
