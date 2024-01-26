package aplication

import (
	"net/http"
	"strconv"

	"github.com/georgifotev1/go-api/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt"
)

func (a *App) authMiddleware(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			handlers.RespondWithError(w, http.StatusUnauthorized, handlers.ErrAuthenticationFailed)
			return
		}

		token, err := handlers.VerifyToken(tokenString)
		if err != nil {
			handlers.RespondWithError(w, http.StatusUnauthorized, handlers.ErrAuthenticationFailed)
			return
		}

		if !token.Valid {
			handlers.RespondWithError(w, http.StatusUnauthorized, handlers.ErrAuthenticationFailed)
			return
		}

		stringId := chi.URLParam(r, "userId")
		userId, err := strconv.ParseInt(stringId, 10, 64)
		if err != nil {
			handlers.RespondWithError(w, http.StatusUnauthorized, handlers.ErrAuthenticationFailed)
			return
		}

		user, err := a.db.GetUserById(r.Context(), userId)
		if err != nil {
			handlers.RespondWithError(w, http.StatusInternalServerError, handlers.ErrAuthenticationFailed)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		if user.ID != int64(claims["id"].(float64)) {
			handlers.RespondWithError(w, http.StatusUnauthorized, handlers.ErrAuthenticationFailed)
		}

		handlerFunc(w, r)
	}
}
