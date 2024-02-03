package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type API struct {
	listenAddress string
	db            dbPgx
}

type funcHttp func(w http.ResponseWriter, r *http.Request) error

func makeHttpFunc(f funcHttp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJson(w, r, err, http.StatusBadRequest)
		}
	}
}

func writeJson(w http.ResponseWriter, r *http.Request, v any, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func NewApiServer(address string, db dbPgx) *API {
	return &API{listenAddress: address, db: db}
}

func (s *API) runServer() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api", s.handleApi)
	fmt.Printf("server running at %s \n", s.listenAddress)
	http.ListenAndServe(s.listenAddress, r)
}

func (s *API) handleApi(r chi.Router) {
	r.Route("/member", s.handleMember)
}
