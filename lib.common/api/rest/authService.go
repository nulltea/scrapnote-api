package rest

import (
	"crypto/rsa"

	"go.kicksware.com/api/service-common/config"
	"go.kicksware.com/api/service-common/core"
	"go.kicksware.com/api/service-common/util"
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
