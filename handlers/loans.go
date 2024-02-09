package handlers

import (
	"encoding/json"
	"libraryManagement/db"
	"libraryManagement/types"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandleLoans(r chi.Router) {
	r.Post("/{id}", makeHttpFunc(CompleteLoan))
	r.Post("/", makeHttpFunc(PostLoan))
}

func PostLoan(w http.ResponseWriter, r *http.Request) error {
	var a *types.Loan

	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		return err
	}

	if err := db.CreateNewLoan(a); err != nil {
		return err
	}

	writeJson(w, http.StatusOK, a)
	return nil
}

func CompleteLoan(w http.ResponseWriter, r *http.Request) error {
	var a *types.Loan

	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		return err
	}

	if err := db.CompleteLoan(a); err != nil {
		return err
	}

	writeJson(w, http.StatusOK, a)
	return nil

}
