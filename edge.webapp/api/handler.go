package api

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"go.kicksware.com/api/service-common/api/rest"
	"go.kicksware.com/api/service-common/core"

	"github.com/timoth-y/scrapnote-api/edge.webapp/api/graph"
	"github.com/timoth-y/scrapnote-api/edge.webapp/api/graph/generated"
	"github.com/timoth-y/scrapnote-api/edge.webapp/config"
	"github.com/timoth-y/scrapnote-api/edge.webapp/core/service"
)

//go:generate gqlgen generate

func NewHandler(records service.RecordService, auth core.AuthService, config config.ServiceConfig) chi.Router {
	router := chi.NewRouter()
	middleware := rest.NewAuthMiddleware(auth)
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5000"},
		AllowCredentials: true,
	}).Handler)
	router.With(middleware.Authenticator)
	router.With(middleware.Authorizer)
	router.Handle("/graph", handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(records)})))
	return router
}