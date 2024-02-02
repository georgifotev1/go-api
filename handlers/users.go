package handlers

import (
	"net/http"

	"github.com/georgifotev1/go-api/database/sqlc"
	"github.com/georgifotev1/go-api/helpers"
	"github.com/georgifotev1/go-api/messages"
	"github.com/lib/pq"
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
	if err := helpers.ReadJSON(r.Body, &params); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, messages.ErrInvalidJSON)
		return
	}

	if !isEmail(params.Email) || !isAlphanumeric(params.Username) || !isAlphanumeric(params.Password) {
		helpers.WriteError(w, http.StatusBadRequest, messages.ErrInvalidInput)
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		helpers.WriteError(w, http.StatusInternalServerError, messages.ErrInternalServer)
		return
	}

	user, err := u.Storage.CreateUser(r.Context(), sqlc.CreateUserParams{
		Username: params.Username,
		Email:    params.Email,
		Password: string(hashedPass),
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			msg := helpers.FormatUniqueConstrainErr(pqErr)
			helpers.WriteError(w, http.StatusForbidden, msg)
			return
		}
		helpers.WriteError(w, http.StatusInternalServerError, messages.ErrInternalServer)
		return
	}

	tokenString, err := helpers.CreateToken(user.ID)
	if err != nil {
		helpers.WriteError(w, http.StatusInternalServerError, messages.ErrInternalServer)
		return
	}

	err = helpers.WriteJSON(w, http.StatusCreated, helpers.FormatUser(user, tokenString))
	if err != nil {
		helpers.WriteError(w, http.StatusInternalServerError, messages.ErrInternalServer)
		return
	}
}

func (u *User) SignIn(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	params := parameters{}
	if err := helpers.ReadJSON(r.Body, &params); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, messages.ErrInvalidJSON)
		return
	}

	if !isEmail(params.Email) || !isAlphanumeric(params.Password) {
		helpers.WriteError(w, http.StatusBadRequest, messages.ErrInvalidInput)
		return
	}

	user, err := u.Storage.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		if _, ok := err.(*pq.Error); ok {
			//TODO fix error
			helpers.WriteError(w, http.StatusForbidden, messages.ErrInvalidInput)
			return
		}
		helpers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		helpers.WriteError(w, http.StatusBadRequest, messages.ErrWrongPassword)
		return
	}

	tokenString, err := helpers.CreateToken(user.ID)
	if err != nil {
		helpers.WriteError(w, http.StatusInternalServerError, messages.ErrInternalServer)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, helpers.FormatUser(user, tokenString))
	if err != nil {
		helpers.WriteError(w, http.StatusInternalServerError, messages.ErrInternalServer)
		return
	}
}

func (u *User) SignOut(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	helpers.BlacklistToken(tokenString)
	w.WriteHeader(http.StatusOK)
}
