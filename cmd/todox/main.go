package main

import (
	"fmt"
	"net/http"
	"os"
	"paganotoni/todox/database"
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
	http.ListenAndServe(addr, buildServer())
}

func buildServer() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(database.Connection)

	router.Get("/", todo.Index)
	router.Get("/search", todo.Search)
	router.Get("/{id}/edit", todo.Edit)
	router.Post("/", todo.Create)
	router.Delete("/{id}", todo.Delete)
	router.Put("/{id}", todo.Update)
	router.Put("/{id}/complete", todo.Complete)

	// Static files like css, images and so on.
	// TODO: review name of the folder
	publicFolder := http.FS(fs.NewFallback(public.Folder, "public/"))
	router.Handle("/*", http.StripPrefix("/", http.FileServer(publicFolder)))

	return router
}
