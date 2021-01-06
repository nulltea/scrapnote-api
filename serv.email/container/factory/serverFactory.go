package factory

import (
	"github.com/timoth-y/scrapnote-api/lib.common/core"
	"github.com/timoth-y/scrapnote-api/lib.common/server"

	"github.com/timoth-y/scrapnote-api/serv.email/config"
)

func ProvideServer(config config.ServiceConfig, handler core.Handler) core.Server {
	srv := server.NewInstance(config.Common.Host)
	srv.SetupAMQP(handler)
	return srv
}
