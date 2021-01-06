package errors

import "errors"

var (
	ErrPasswordInvalid       = errors.New("serv.auth/authService: invalid user password")
	ErrTokenInvalid          = errors.New("serv.auth/authService: invalid auth token")
	ErrNotConfirmed          = errors.New("serv.auth/authService: user email not confirmed")
	ErrInvalidRemoteID       = errors.New("serv.auth/authService: invalid remote OAuth identifier")
	ErrInvalidRemoteProvider = errors.New("serv.auth/authService: invalid remote OAuth provider")
	ErrEmailInvalid  = errors.New("user email could not be empty")
)