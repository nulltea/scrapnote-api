package factory

import (
	"github.com/go-chi/chi"
	"github.com/timoth-y/scrapnote-api/lib.common/core"

	"github.com/timoth-y/scrapnote-api/edge.plugin/api"
	"github.com/timoth-y/scrapnote-api/edge.plugin/config"
	"github.com/timoth-y/scrapnote-api/edge.plugin/core/service"
)

func ProvideEdgeHandler(service service.RecordService, auth core.AuthService, config config.ServiceConfig) *api.Handler {
	return api.NewHandler(service, auth, config.Common)
}

func ProvideEndpointRouter(handler *api.Handler) chi.Router {
	return api.ProvideRoutes(handler)
}