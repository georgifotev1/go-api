package messages

const (
	ErrInternalServer       = "internal serve error"
	ErrInvalidJSON          = "bad request: invalid JSON"
	ErrInvalidInput         = "bad request: invalid input"
	ErrWrongPassword        = "bad request: wrong password"
	ErrAuthenticationFailed = "authentication failed"
	ErrBadUsername          = "username is already taken"
	ErrBadEmail             = "email is already taken"
	ErrUniqueConstraint     = "unique constraint violation"
	ErrFailedRegistration   = "failed to create user"
)
