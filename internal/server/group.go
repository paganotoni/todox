package server

import (
	"net/http"

	"io/fs"
)

type RouteGroup struct {
	prefix     string
	mux        *http.ServeMux
	middleware []func(http.Handler) http.Handler
}

func (rg *RouteGroup) Use(middleware func(http.Handler) http.Handler) {
	rg.middleware = append([]func(http.Handler) http.Handler{middleware}, rg.middleware...)
}

func (rg *RouteGroup) Handle(pattern string, handler http.Handler) {
	for _, v := range rg.middleware {
		handler = v(handler)
	}

	rg.mux.Handle(pattern, handler)
}

func (rg *RouteGroup) HandleFunc(pattern string, handler http.HandlerFunc) {
	rg.Handle(pattern, http.HandlerFunc(handler))
}

func (rg *RouteGroup) Folder(prefix string, fs fs.FS) {
	rg.mux.Handle(
		"GET "+prefix+"/*",
		http.StripPrefix(prefix, http.FileServer(http.FS(fs))),
	)
}
