package api

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	"github.com/timoth-y/scrapnote-api/edge.webapp/api/graph"
	"github.com/timoth-y/scrapnote-api/edge.webapp/api/graph/generated"
	"github.com/timoth-y/scrapnote-api/edge.webapp/config"
	"github.com/timoth-y/scrapnote-api/edge.webapp/core/service"
)

//go:generate gqlgen generate

func NewHandler(records service.RecordService, config config.ServiceConfig) chi.Router {
	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5000"},
		AllowCredentials: true,
	}).Handler)
	router.Handle("/graph", handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(records)})))
	return router
}