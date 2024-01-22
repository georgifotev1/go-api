package handlers

import (
	"log"
	"net/http"

	"github.com/georgifotev1/go-api/database/sqlc"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Storage *sqlc.Queries
}

func (u *User) Register(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	params := parameters{}
	if err := ReadJSON(r.Body, &params); err != nil {
		WriteError(w, http.StatusBadRequest, "bad request")
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := u.Storage.CreateUser(r.Context(), sqlc.CreateUserParams{
		Username: params.Username,
		Email:    params.Email,
		Password: string(hashedPass),
	})

	if err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	tokenString, err := createToken(user.Username)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := map[string]interface{}{
		"user":  user,
		"token": tokenString,
	}

	log.Fatal(WriteJSON(w, http.StatusOK, response))
}

func (u *User) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Register"))
}

func (u *User) SignOut(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout"))
}
