package main

import (
	"fmt"
	"net/http"
	"os"
	"paganotoni/todox/internal/database"
	"paganotoni/todox/internal/fs"
	"paganotoni/todox/public"
	"paganotoni/todox/todo"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Get the port from the environment
	addr := ":3000"
	if e := os.Getenv("PORT"); e != "" {
		addr = ":" + e
	}

	// Start the server
	fmt.Println("Server listening on", addr)
	http.ListenAndServe(addr, router())
}

func router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(database.Connection)

	// Mounting the profiler at /debug to :eyes: the app
	r.Mount("/debug", middleware.Profiler())

	r.Get("/", todo.Index)
	r.Get("/search", todo.Search)
	r.Get("/{id}/edit", todo.Edit)
	r.Post("/", todo.Create)
	r.Delete("/{id}", todo.Delete)
	r.Put("/{id}", todo.Update)
	r.Put("/{id}/complete", todo.Complete)

	// Static files like css, images and so on.
	// TODO: review name of the folder
	publicFolder := http.FS(fs.NewFallback(public.Folder, "public/"))
	r.Handle("/*", http.StripPrefix("/", http.FileServer(publicFolder)))

	return r
}
