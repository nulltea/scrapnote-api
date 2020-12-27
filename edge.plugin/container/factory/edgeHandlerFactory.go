package factory

import (
	"github.com/go-chi/chi"

	"github.com/timoth-y/scrapnote-api/edge.plugin/api"
	"github.com/timoth-y/scrapnote-api/edge.plugin/config"
	"github.com/timoth-y/scrapnote-api/edge.plugin/core/service"
)

func ProvideEdgeHandler(service service.RecordService,  config config.ServiceConfig) *api.Handler {
	return api.NewHandler(service, config.Common)
}

func ProvideEndpointRouter(handler *api.Handler) chi.Router {
	return api.ProvideRoutes(handler)
}