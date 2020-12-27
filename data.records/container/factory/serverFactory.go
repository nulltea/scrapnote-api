package factory

import (
	"go.kicksware.com/api/service-common/core"
	"go.kicksware.com/api/service-common/server"

	"github.com/timoth-y/scrapnote-api/data.records/api/rpc"
	"github.com/timoth-y/scrapnote-api/data.records/config"
)

func ProvideServer(config config.ServiceConfig, asyncHandler core.Handler, rpcHandler *rpc.Handler) core.Server {
	srv := server.NewInstance(config.Common.Host)
	srv.SetupAuth()
	srv.SetupGRPC(rpc.ProvideRemoteSetup(rpcHandler))
	srv.SetupAMQP(asyncHandler)
	return srv
}
