package service

import (
	"crypto/rsa"
	"errors"

	"github.com/timoth-y/scrapnote-api/data.users/core/model"

	"github.com/timoth-y/scrapnote-api/serv.auth/core/meta"
)

var (
	ErrPasswordInvalid       = errors.New("serv.auth/authService: invalid user password")
	ErrTokenInvalid          = errors.New("serv.auth/authService: invalid auth token")
	ErrNotConfirmed          = errors.New("serv.auth/authService: user email not confirmed")
	ErrInvalidRemoteID       = errors.New("serv.auth/authService: invalid remote OAuth identifier")
	ErrInvalidRemoteProvider = errors.New("serv.auth/authService: invalid remote OAuth provider")
)

type AuthService interface {
	SingUp(user *model.User) error
	Login(user *model.User) (*meta.AuthToken, error)
	Remote(user *model.User) (*meta.AuthToken, error)
	GenerateToken(user *model.User) (*meta.AuthToken, error)
	Refresh(token string) (*meta.AuthToken, error)
	PublicKey() *rsa.PublicKey
	Logout(token string) error
}
