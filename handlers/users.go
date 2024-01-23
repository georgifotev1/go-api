package handlers

import (
	"net/http"

	"github.com/georgifotev1/go-api/database/sqlc"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

const (
	ErrInternalServer = "internal serve error"
	ErrInvalidJSON    = "bad request: invalid JSON"
	ErrInvalidInput   = "bad request: invalid input"
	ErrWrongPassword  = "bad request: wrong password"
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
	if err := readJSON(r.Body, &params); err != nil {
		respondWithError(w, http.StatusBadRequest, ErrInvalidJSON)
		return
	}

	if !isEmail(params.Email) || !isValid(params.Username) || !isValid(params.Password) {
		respondWithError(w, http.StatusBadRequest, ErrInvalidInput)
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	user, err := u.Storage.CreateUser(r.Context(), sqlc.CreateUserParams{
		Username: params.Username,
		Email:    params.Email,
		Password: string(hashedPass),
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			msg := formatUniqueConstrainErr(pqErr)
			respondWithError(w, http.StatusForbidden, msg)
			return
		}
		respondWithError(w, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	tokenString, err := createToken(user.Username)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	err = respondWithJSON(w, http.StatusCreated, formatUser(user, tokenString))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, ErrInternalServer)
		return
	}
}

func (u *User) SignIn(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	params := parameters{}
	if err := readJSON(r.Body, &params); err != nil {
		respondWithError(w, http.StatusBadRequest, ErrInvalidJSON)
		return
	}

	if !isEmail(params.Email) || !isValid(params.Password) {
		respondWithError(w, http.StatusBadRequest, ErrInvalidInput)
		return
	}

	user, err := u.Storage.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		if _, ok := err.(*pq.Error); ok {
			//TODO fix error
			respondWithError(w, http.StatusForbidden, ErrInvalidInput)
			return
		}
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, ErrWrongPassword)
		return
	}

	tokenString, err := createToken(user.Username)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	err = respondWithJSON(w, http.StatusOK, formatUser(user, tokenString))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, ErrInternalServer)
		return
	}
}

func (u *User) SignOut(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout"))
}
