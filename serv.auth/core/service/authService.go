package service

import (
	"context"
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
	SingUp(ctx context.Context, user *model.User) error
	Login(ctx context.Context, user *model.User) (*meta.AuthToken, error)
	Remote(ctx context.Context, user *model.User) (*meta.AuthToken, error)
	GenerateToken(ctx context.Context, user *model.User) (*meta.AuthToken, error)
	Refresh(ctx context.Context, token string) (*meta.AuthToken, error)
	PublicKey() *rsa.PublicKey
	Logout(ctx context.Context, token string) error
}
