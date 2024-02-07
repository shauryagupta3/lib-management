package handlers

import (
	"encoding/json"
	"libraryManagement/db"
	"libraryManagement/types"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func HandleBook(r chi.Router) {
	r.Post("/", makeHttpFunc(createBook))
	r.Get("/author/{id}", makeHttpFunc(HandleGetBookByAuthorID))
	r.Get("/{id}", makeHttpFunc(getBookByID))
}

func createBook(w http.ResponseWriter, r *http.Request) error {
	var a *types.Book
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		return err
	}
	err := db.CreateBook(a)
	if err != nil {
		return err
	}
	writeJson(w, http.StatusOK, a)
	return nil
}

func getBookByID(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	book, err := db.GetBook(i)
	if err != nil {
		return err
	}
	writeJson(w, http.StatusOK, book)
	return nil
}

func HandleGetBookByAuthorID(w http.ResponseWriter, r *http.Request) error {

	id := chi.URLParam(r, "id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	books, err := db.GetBooksByAuthorID(i)
	if err != nil {
		return err
	}
	writeJson(w, http.StatusOK, books)

	return nil
}

