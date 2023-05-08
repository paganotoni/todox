package database

import (
	"context"
	"net/http"
)

func Connection(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := connection()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		h.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "db", conn)))
	})
}
