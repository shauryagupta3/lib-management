package handlers

import (
	"libraryManagement/db"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func HandleMember(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		writeJson(w, http.StatusOK, "hello world")
	})
	r.Get("/{id}", makeHttpFunc(getMember))
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
