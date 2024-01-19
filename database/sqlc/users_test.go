package sqlc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUser(t *testing.T) {
	t.Run("test CreateUser", func(t *testing.T) {
		arg := CreateUserParams{
			Username: "Test",
			Email:    "test@gmail.com",
			Password: "secret",
		}
		acc, err := testQueries.CreateUser(context.Background(), arg)

		require.NoError(t, err)
		require.NotEmpty(t, acc)
		require.Equal(t, acc.Username, arg.Username)
		require.NotZero(t, acc.ID)
	})
}
