package server

import (
	"net/http"
	"path"
	"strings"

	"io/fs"
)

// HandlerGroup is a group of routes with a common prefix and middleware
// that should be executed for all the handlers in the group
type HandlerGroup struct {
	prefix     string
	mux        *http.ServeMux
	middleware []Middleware
}

// Use allows to specify a middleware that should be executed for all the handlers
// in the group
func (rg *HandlerGroup) Use(middleware Middleware) {
	// Add the middleware to the beginning of the middleware chain
	// so that it is executed first
	rg.middleware = append([]Middleware{middleware}, rg.middleware...)
}

// Handle allows to register a new handler for a specific pattern
// in the group with the middleware that should be executed for the handler
// specified in the group.
func (rg *HandlerGroup) Handle(pattern string, handler http.Handler) {
	for _, v := range rg.middleware {
		handler = v(handler)
	}

	method := ""
	route := pattern

	if parts := strings.Split(pattern, " "); len(parts) > 1 {
		method = parts[0]
		route = parts[1]
	}

	pattern = strings.Join([]string{method, rg.prefix + route}, " ")
	rg.mux.Handle(pattern, handler)
}

// HandleFunc allows to register a new handler function for a specific pattern
// in the group with the middleware that should be executed for the handler
// specified in the group.
func (rg *HandlerGroup) HandleFunc(pattern string, handler http.HandlerFunc) {
	rg.Handle(pattern, http.HandlerFunc(handler))
}

// Folder allows to serve static files from a directory
func (rg *HandlerGroup) Folder(prefix string, fs fs.FS) {
	rg.mux.Handle(
		"GET "+prefix+"/*",
		http.StripPrefix(prefix, http.FileServer(http.FS(fs))),
	)
}

// Group allows to create a new group of routes with a common prefix
// and middleware that should be executed for all the handlers in the group
func (rg *HandlerGroup) Group(prefix string, rfn func(rg *HandlerGroup)) {
	group := &HandlerGroup{
		prefix:     path.Join(rg.prefix, prefix),
		mux:        http.NewServeMux(),
		middleware: rg.middleware,
	}

	rfn(group)
	rg.mux.Handle(prefix, group.mux)
}
