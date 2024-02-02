package helpers

import (
	"time"

	"github.com/georgifotev1/go-api/database/sqlc"
	"github.com/georgifotev1/go-api/messages"
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

func FormatUser(user sqlc.User, token string) fUser {
	return fUser{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Token:     token,
	}
}

func FormatUniqueConstrainErr(err *pq.Error) string {
	if err.Code == "23505" {
		switch err.Constraint {
		case "users_username_key":
			return messages.ErrBadUsername
		case "users_email_key":
			return messages.ErrBadEmail
		default:
			return messages.ErrUniqueConstraint
		}
	}
	return messages.ErrFailedRegistration
}
