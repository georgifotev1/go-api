package handlers

import (
	"net/http"

	"github.com/georgifotev1/go-api/database/sqlc"
)

type User struct {
	Storage *sqlc.Queries
}

func (u *User) Register(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	p := params{}
	if err := ReadJSON(r.Body, &p); err != nil {
		WriteError(w, http.StatusBadRequest, "bad request")
	}
}

func (u *User) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Register"))
}

func (u *User) SignOut(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout"))
}
