package aplication

import (
	"fmt"
	"net/http"

	"github.com/georgifotev1/go-api/handlers"
)

func (a *App) authMiddleware(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			handlers.RespondWithError(w, http.StatusUnauthorized, handlers.ErrInvalidToken)
			return
		}

		err := handlers.VerifyToken(tokenString)
		if err != nil {
			handlers.RespondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		fmt.Println("Token is valid")
		handlerFunc(w, r)
	}
}
