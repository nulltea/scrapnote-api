package rest

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func ProvideRoutes(rest *Handler) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.RequestID,
		middleware.RealIP,
	)
	router.Mount("/records", restRoutes(rest))
	return router
}

func restRoutes(rest *Handler) (r *chi.Mux) {
	r = chi.NewRouter()
	// r.Use(rest.auth.Authenticator)
	// r.Use(rest.auth.Authorizer)
	r.Get("/{recordID}", rest.Get)
	r.Post("/", rest.Post)
	r.Patch("/", rest.Patch)
	return
}
