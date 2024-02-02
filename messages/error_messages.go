package messages

const (
	ErrInternalServer       = "internal serve error"
	ErrInvalidJSON          = "bad request: invalid JSON"
	ErrInvalidInput         = "bad request: invalid input"
	ErrAuthenticationFailed = "authentication failed"
	ErrSessionExists        = "session already exists"
	ErrBadUsername          = "username is already taken"
	ErrBadEmail             = "email is already taken"
	ErrUniqueConstraint     = "unique constraint violation"
	ErrFailedRegistration   = "failed to create user"
)
