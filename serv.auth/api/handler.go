package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"github.com/timoth-y/scrapnote-api/data.users/core/model"
	"go.kicksware.com/api/service-common/config"

	users "github.com/timoth-y/scrapnote-api/data.users/core/service"
	usersJson "github.com/timoth-y/scrapnote-api/data.users/usecase/serializer/json"

	"github.com/timoth-y/scrapnote-api/serv.auth/core/meta"
	"github.com/timoth-y/scrapnote-api/serv.auth/core/service"
)

type Handler struct {
	service     service.AuthService
	contentType string
}

func NewHandler(service service.AuthService, config config.CommonConfig) *Handler {
	return &Handler{
		service,
		config.ContentType,
	}
}

func (h *Handler) SingUp(w http.ResponseWriter, r *http.Request) {
	user, err := h.getRequestBody(r); if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := h.service.SingUp(user); if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupAuthCookie(w, token)
	h.setupResponse(w, token, http.StatusOK)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	user, err := h.getRequestBody(r); if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := h.service.Login(user); if err != nil {
		if errors.Cause(err) == service.ErrPasswordInvalid ||
			errors.Cause(err) == service.ErrNotConfirmed {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupAuthCookie(w, token)
	h.setupResponse(w, token, http.StatusOK)
}

func (h *Handler) Remote(w http.ResponseWriter, r *http.Request) {
	user, err := h.getRequestBody(r); if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := h.service.Remote(user); if err != nil {
		if errors.Cause(err) == service.ErrInvalidRemoteID {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		} else if errors.Cause(err) == service.ErrInvalidRemoteProvider {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupAuthCookie(w, token)
	h.setupResponse(w, token, http.StatusOK)
}

func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	token, err := h.service.Refresh(chi.URLParam(r,"token")); if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupAuthCookie(w, token)
	h.setupResponse(w, token, http.StatusOK)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r,"token")
	if err := h.service.Logout(token); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.recallAuthCookie(w)
	h.setupResponse(w, token, http.StatusOK)
}

func (h *Handler) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := h.getRequestToken(r); if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			fmt.Println()
			return
		}

		if token == nil || !token.Valid {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) setupResponse(w http.ResponseWriter, body interface{}, statusCode int) {
	w.Header().Set("Content-Type", h.contentType)
	w.WriteHeader(statusCode)
	if body != nil {
		raw, err := h.serializer(h.contentType).Encode(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(raw); err != nil {
			log.Println(err)
		}
	}
}

func (h *Handler) setupAuthCookie(w http.ResponseWriter, token *meta.AuthToken) {
	cookie := &http.Cookie{
		Name: "AuthToken",
		Value: token.Token,
		Expires: token.Expires,
	}
	http.SetCookie(w, cookie)
}

func (h *Handler) recallAuthCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name: "AuthToken",
		Expires: time.Now(),
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func (h *Handler) getRequestBody(r *http.Request) (*model.User, error) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	body, err := h.serializer(contentType).Decode(requestBody)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (h *Handler) serializer(contentType string) users.UserSerializer {
	return usersJson.NewSerializer()
}

func (h *Handler) getRequestToken(r *http.Request) (token *jwt.Token, err error) {
	token, err = request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); ok {
			return h.service.PublicKey(), nil
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
