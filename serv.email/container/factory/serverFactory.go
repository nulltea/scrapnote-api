package factory

import (
	"go.kicksware.com/api/service-common/core"
	"go.kicksware.com/api/service-common/server"

	"github.com/timoth-y/scrapnote-api/serv.email/config"
)

func ProvideServer(config config.ServiceConfig, handler core.Handler) core.Server {
	srv := server.NewInstance(config.Common.Host)
	srv.SetupAMQP(handler)
	return srv
}
