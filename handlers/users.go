package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandleUsers(r chi.Router) {
	//r.Post("/",)
}

func PostUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
