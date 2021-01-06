package factory

import (
	"github.com/go-chi/chi"
	"github.com/timoth-y/scrapnote-api/lib.common/core"
	"github.com/timoth-y/scrapnote-api/lib.common/server"

	"github.com/timoth-y/scrapnote-api/serv.auth/config"
)

func ProvideServer(config config.ServiceConfig, router chi.Router) core.Server {
	srv := server.NewInstance(config.Common.Host)
	srv.SetupREST(router)
	return srv
}
