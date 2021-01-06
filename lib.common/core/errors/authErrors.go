package errors

import "errors"

var (
	ErrUserContextInfoMissing = errors.New("user context info missing")
	ErrWrongUserContext = errors.New("user context is wrong")
)
