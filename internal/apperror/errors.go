package apperror

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrContextNotEmpty = errors.New("cannot delete non-empty context")
	ErrUnathorized = errors.New("unauthorized")
	ErrBadCredentials = errors.New("invalid password or email")
	ErrTaskIsOverdue = errors.New("task deadline has passed")
	ErrToken = errors.New("couldn't generate token")
	ErrEmailExists = errors.New("email already exists")
	ErrInvalidToken = errors.New("invalid token")
)

