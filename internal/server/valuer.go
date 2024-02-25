package server

import (
	"context"
	"net/http"
	"sync"
)

// values is a map where we can store values for the request context
// these values will then be available for other components such as
// the render engine.
type valuer struct {
	data map[string]any
	moot sync.Mutex
}

// Value returns the value for the key specified.
func (v *valuer) Value(key string) any {
	return v.data[key]
}

// Values returns the values stored in the valuer.
func (v *valuer) Values() map[string]any {
	return v.data
}

// Set sets the value for the key specified.
func (v *valuer) Set(key string, value any) {
	v.moot.Lock()
	defer v.moot.Unlock()

	v.data[key] = value
}

// For each of the requests we want to have a valuer intance so
// that we can store values in the context that can be used by
// other components.
func setValuer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vlr := &valuer{
			data: map[string]any{
				// Adding base values that are useful for the handlers.
				"request":    r,
				"currentURL": r.URL.String(),
			},
		}

		r = r.WithContext(context.WithValue(r.Context(), "valuer", vlr))
		next.ServeHTTP(w, r)
	})
}
