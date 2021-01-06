package rest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/golang/glog"

	"go.kicksware.com/api/service-common/core"
	"go.kicksware.com/api/service-common/core/meta"
)

var (
	ErrInvalidTokenClaims = errors.New("invalid token claims")
	GuestRole = "gst"
	UserContextKey = ContextKey("userinfo")
)

type ContextKey string

type AuthMiddleware struct {
	service core.AuthService
}


func NewAuthMiddleware(service core.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		service,
	}
}

func (m *AuthMiddleware) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := m.getRequestToken(r); if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			glog.Errorln(err)
			glog.Errorln(token.Raw)
			glog.Errorln(token.Claims)
			return
		}

		if token == nil || !token.Valid {
			http.Error(w, ErrInvalidTokenClaims.Error(), http.StatusInternalServerError)
			glog.Errorln(err)
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}

func (m *AuthMiddleware) Authorizer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := m.getRequestToken(r); if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			glog.Errorln(err)
			return
		}

		claims, err := getClaims(token); if err != nil {
			http.Error(w, ErrInvalidTokenClaims.Error(), http.StatusInternalServerError)
			glog.Errorln(err)
			return
		}
		if claims != nil && claims.Role != GuestRole {
			r.URL.User = url.UserPassword(claims.UniqueID, token.Raw)
			ctx := context.WithValue(r.Context(), UserContextKey, r.URL.User)
			r = r.WithContext(ctx)
		} else {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			fmt.Println()
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m *AuthMiddleware) UserSetter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := m.getRequestToken(r); if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			glog.Errorln(err)
			return
		}

		if claims, err := getClaims(token); err == nil && claims != nil && claims.Role != GuestRole {
			r.URL.User = url.UserPassword(claims.UniqueID, token.Raw)
			ctx := context.WithValue(r.Context(), UserContextKey, meta.UserContextInfo{
				UniqueID: claims.UniqueID,
				Token: token.Raw,
			})
			r = r.WithContext(ctx)
		} else {
			r.URL.User = nil
		}

		next.ServeHTTP(w, r)
	})
}

func (m *AuthMiddleware) getRequestToken(r *http.Request) (token *jwt.Token, err error) {
	token, err = request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); ok {
			return m.service.PublicKey(), nil
		}
		return nil, fmt.Errorf("authenticator: unexpected signing method: %q", token.Header["alg"])
	})
	return
}

func getClaims(token *jwt.Token) (*meta.AuthClaims, error) {
	payload, err := json.Marshal(token.Claims); if err != nil {
		return nil, err
	}
	claims := &meta.AuthClaims{}

	if err = json.Unmarshal(payload, claims); err != nil {
		return nil, err
	}
	return claims, nil
}
