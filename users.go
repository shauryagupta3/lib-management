package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var Users = []User{
	{
		Username: "sg",
		Password: "sg",
	},
}

func handleUser(r chi.Router) {
	r.Get("/", makeHttpFunc(GetUser))
	r.Post("/", makeHttpFunc(CreateUser))
	r.Delete("/", makeHttpFunc(DeleteUser))
}

func GetUser(w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	var u User
	if err := decoder.Decode(&u); err != nil {
		return err
	}
	fmt.Println(u.Username)
	var actualUser User

	for _, user := range Users {
		if user.Username == u.Username && user.Password == u.Password {
			actualUser = user
			break
		}
	}

	if actualUser.Username == "" {
		return errors.New("no user found")
	}

	writeJson(w, r, actualUser, 200)
	return nil
}

func CreateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func DeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
