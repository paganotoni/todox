package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

func requestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(context.WithValue(r.Context(), "requestID", time.Now().UnixNano()))
		next.ServeHTTP(w, r)
	})
}

// logger is a middleware that logs the request method and URL
// and the time it took to process the request.
func logger(next http.Handler) http.Handler {
	// TODO: allow logger.With to be passed as an option
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		slog.Info("", "method", r.Method, "url", r.URL.Path, "took", time.Since(start))
	})
}

// recoverer is a middleware that recovers from panics and logs the error.
func recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				slog.Error("panic", "error", err, "method", r.Method, "url", r.URL.Path)

				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
