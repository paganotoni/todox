package internal

import (
	"net/http"

	"github.com/leapkit/leapkit/core/server"
)

func health(w http.ResponseWriter, r *http.Request) {
	db, err := DB()
	if err != nil {
		server.Error(w, err, http.StatusInternalServerError)
		return
	}

	err = db.Ping()
	if err != nil {
		server.Error(w, err, http.StatusInternalServerError)
		return
	}

	w.Write([]byte("OK"))
}
