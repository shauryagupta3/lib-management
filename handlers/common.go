package handlers

import (
	"encoding/json"
	"libraryManagement/types"
	"net/http"
)

type DbConn struct {
	conn types.DB
}

type funcWithError func(w http.ResponseWriter, r *http.Request) error

func makeHttpFunc(f funcWithError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJson(w, http.StatusBadGateway, err)
		}
	}
}

func writeJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
