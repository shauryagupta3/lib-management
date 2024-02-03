package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (s *API) handleMember(r chi.Router) {
	r.Get("/{id}", makeHttpFunc(s.handleGetMember))
	r.Post("/",makeHttpFunc(s.handleCreateMember))
}

func (s *API) handleGetMember(w http.ResponseWriter, r *http.Request) error {
	idOfMember := chi.URLParam(r, "id")
	i, err := strconv.Atoi(idOfMember)
	if err != nil {
		return err
	}
	member, err := s.db.getMemberById(i)
	writeJson(w, r, member, http.StatusOK)
	return nil
}

func (s *API) handleCreateMember(w http.ResponseWriter, r *http.Request) error {
	var m Member
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		return err
	}
	err := s.db.createMemeber(m)
	if err != nil {
		return err
	}
	writeJson(w, r, m, http.StatusOK)
	return nil
}
