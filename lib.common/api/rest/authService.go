package rest

import (
	"crypto/rsa"

	"github.com/timoth-y/scrapnote-api/lib.common/config"
	"github.com/timoth-y/scrapnote-api/lib.common/core"
	"github.com/timoth-y/scrapnote-api/lib.common/util"
)

type authService struct {
	publicKey *rsa.PublicKey
}

func NewAuthService(authConfig config.AuthConfig) core.AuthService {
	return &authService{
		util.GetPublicKey(authConfig.PublicKeyPath),
	}
}

func (s *authService) PublicKey() *rsa.PublicKey {
	return s.publicKey
}

func (s *authService) Authenticate() (string, error)  {
	panic("not implemented")
}
