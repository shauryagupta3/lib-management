package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *API) handleBooks(r chi.Router) {
	r.Post("/", makeHttpFunc(s.handleCreateBook))
}

func (s *API) handleCreateBook(w http.ResponseWriter, r *http.Request) error {
	var a Book
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		return err
	}

	if err := s.db.createBook(&a); err != nil {
		return err
	}
	writeJson(w, r, a, http.StatusOK)
	return nil
}

func getBook(w http.ResponseWriter, r *http.Request) error { return nil }

func (s *dbPgx) createBooksTable() error {
	_, err := s.conn.Query(context.Background(), "create table if not exists books(id serial primary key, title varchar(255) not null unique, release_year int not null, genre varchar(255))")
	if err != nil {
		return err
	}
	return nil
}

func (s *dbPgx) getBookById(id int) (*Book, error) {
	var a Book

	err := s.conn.QueryRow(context.Background(), "select id,title,year,genre from books where id=$1", id).Scan(&a.Id, &a.Title, &a.Year, &a.Genre)
	if err != nil {
		return nil, err
	}

	return &a, err
}

func (s *dbPgx) createBook(a *Book) error {

	_, err := s.conn.Query(context.Background(), "insert into books(title,release_year,genre) values ($1,$2,$3)", a.Title, a.Year, a.Genre)
	if err != nil {
		return err
	}

	for _, author := range a.Authors {
		if err := s.createAuthor(&author); err != nil {
			return err
		}
	}

	return nil
}
