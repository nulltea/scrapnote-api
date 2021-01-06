package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func ProvideRoutes(rest *Handler) chi.Router {
	router := chi.NewRouter()
	router.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.RequestID,
		middleware.RealIP,
	)
	router.Mount("/auth", restRoutes(rest))
	return router
}

func restRoutes(rest *Handler) (r *chi.Mux) {
	r = chi.NewRouter()
	r.Post("/sign-up", rest.SingUp)
	r.Post("/login", rest.Login)
	r.Post("/remote", rest.Remote)
	r.Get("/token-refresh", rest.RefreshToken)
	r.Get("/logout", rest.Logout)
	return
}
