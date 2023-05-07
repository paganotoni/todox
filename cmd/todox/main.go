package main

import (
	"net/http"
	"paganotoni/todox/internal/fs"
	"paganotoni/todox/public"
	"paganotoni/todox/todo"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// dd, _ := public.Folder.ReadDir(".")
	// for _, v := range dd {
	// 	println(v.Name())
	// }

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

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

	http.ListenAndServe(":3000", router)
}
