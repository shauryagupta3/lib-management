package api

import (
	"libraryManagement/handlers"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type API struct {
	listenAddress string
}

func NewServer(address string) *API {
	return &API{
		listenAddress: address,
	}
}

func (s *API) RunServer() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/", func(r chi.Router) {
		r.Route("/book", handlers.HandleBook)
		r.Route("/member", handlers.HandleMember)
		r.Route("/user", handlers.HandleUsers)
	})

	return http.ListenAndServe(s.listenAddress, r)
}
