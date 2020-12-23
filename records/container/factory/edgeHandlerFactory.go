package factory

import (
	"github.com/go-chi/chi"

	rest "github.com/timoth-y/scrapnote-api/records/api/REST"
	"github.com/timoth-y/scrapnote-api/records/config"
	"github.com/timoth-y/scrapnote-api/records/core/service"
)

func ProvideEdgeHandler(service service.RecordService,  config config.ServiceConfig) *rest.Handler {
	return rest.NewHandler(service, config.Common)
}

func ProvideEndpointRouter(handler *rest.Handler) chi.Router {
	return rest.ProvideRoutes(handler)
}