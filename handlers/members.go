package handlers

import (
	"encoding/json"
	"libraryManagement/db"
	"libraryManagement/types"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func HandleMember(r chi.Router) {
	r.Get("/{id}", makeHttpFunc(getMember))
	r.Post("/", makeHttpFunc(PostMember))
}

func getMember(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	member, err := db.GetMember(i)
	if err != nil {
		return err
	}
	writeJson(w, http.StatusOK, member)
	return nil
}

func PostMember(w http.ResponseWriter, r *http.Request) error {
	var a *types.Member

	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		return err
	}

	if err := db.CreateNewMember(a); err != nil {
		return err
	}

	writeJson(w, http.StatusOK, a)
	return nil
}
