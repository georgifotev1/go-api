package handlers

import (
	"time"

	"github.com/georgifotev1/go-api/database/sqlc"
	"github.com/lib/pq"
)

type fUser struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Token     string    `json:"token"`
}

func formatUser(user sqlc.User, token string) fUser {
	return fUser{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Token:     token,
	}
}

func formatUniqueConstrainErr(err *pq.Error) string {
	if err.Code == "23505" {
		switch err.Constraint {
		case "users_username_key":
			return "username is already taken"
		case "users_email_key":
			return "email is already taken"
		default:
			return "unique constraint violation"
		}
	}
	return "failed to create user"
}
