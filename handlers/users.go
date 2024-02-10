package handlers

import (
	"encoding/json"
	"errors"
	"libraryManagement/db"
	"libraryManagement/types"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
)

func HandleUsers(r chi.Router) {
	r.Get("/", makeHttpFunc(GetUser))
	r.Post("/", makeHttpFunc(PostUser))
}

var JWT_SECRET []byte = []byte("thisismyjwtsecret")

func GetUser(w http.ResponseWriter, r *http.Request) error {
	var a *types.User

	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		return err
	}
	err := db.AuthUser(a)
	if err != nil {
		return err
	}

	token, err := CreateJWTtoken(a.Email)
	if err != nil {
		return err
	}

	writeJson(w, http.StatusOK, token)

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

func CreateJWTtoken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix()})
	tokenString, err := token.SignedString(JWT_SECRET)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return JWT_SECRET, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("please try again!")
	}

	return nil
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return errors.New("Missing authorization header")
	}
	err := verifyToken(tokenString)
	if err != nil {
		return err
	}
	return nil
}
