package aplication

import (
	"net/http"
	"strconv"

	"github.com/georgifotev1/go-api/helpers"
	"github.com/georgifotev1/go-api/messages"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt"
)

func (a *App) authMiddleware(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			helpers.WriteError(w, http.StatusUnauthorized, messages.ErrAuthenticationFailed)
			return
		}

		token, err := helpers.VerifyToken(tokenString)
		if err != nil {
			helpers.WriteError(w, http.StatusUnauthorized, messages.ErrAuthenticationFailed)
			return
		}

		if !token.Valid {
			helpers.WriteError(w, http.StatusUnauthorized, messages.ErrAuthenticationFailed)
			return
		}

		stringId := chi.URLParam(r, "userId")
		userId, err := strconv.ParseInt(stringId, 10, 64)
		if err != nil {
			helpers.WriteError(w, http.StatusUnauthorized, messages.ErrAuthenticationFailed)
			return
		}

		user, err := a.db.GetUserById(r.Context(), userId)
		if err != nil {
			helpers.WriteError(w, http.StatusInternalServerError, messages.ErrAuthenticationFailed)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		if user.ID != int64(claims["id"].(float64)) {
			helpers.WriteError(w, http.StatusUnauthorized, messages.ErrAuthenticationFailed)
		}

		handlerFunc(w, r)
	}
}
