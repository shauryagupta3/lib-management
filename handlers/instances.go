package handlers

import (
	"encoding/json"
	"libraryManagement/db"
	"libraryManagement/types"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandleInstances(r chi.Router) {
	r.Post("/", makeHttpFunc(PostInstance))
}

func PostInstance(w http.ResponseWriter, r *http.Request) error {
	// create new instance of a particular book
	var a *types.Instance

	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		return err
	}

	if err := db.CreateNewInstance(a); err != nil {
		return err
	}
	return nil
}

func DeleteInstance(w http.ResponseWriter, r *http.Request) error {
	var a *types.Instance

	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		return err
	}

	if err := db.DeleteInstance(a); err != nil {
		return err
	}

	return nil
}
