package handlers

import (
	"encoding/json"
	"libraryManagement/db"
	"libraryManagement/types"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandleUsers(r chi.Router) {
	r.Get("/", makeHttpFunc(GetUser))
	r.Post("/", makeHttpFunc(PostUser))
}

func GetUser(w http.ResponseWriter, r *http.Request) error {
	var a *types.User

	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		return err
	}

	err := db.AuthUser(a)
	if err != nil {
		return err
	}

	writeJson(w, http.StatusOK, "welcome")

	return nil

}

func PostUser(w http.ResponseWriter, r *http.Request) error {
	var a *types.User

	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		return err
	}

	if err := db.CreateUser(a); err != nil {
		return err
	}

	writeJson(w, http.StatusOK, "user created")

	return nil
}
