package config

import (
	"go.kicksware.com/api/service-common/container"
	"go.kicksware.com/api/service-common/core"

	"github.com/timoth-y/scrapnote-api/data.records/api/async"
	"github.com/timoth-y/scrapnote-api/data.records/api/rpc"
	"github.com/timoth-y/scrapnote-api/data.records/config"
	"github.com/timoth-y/scrapnote-api/data.records/container/factory"
	"github.com/timoth-y/scrapnote-api/data.records/usecase/serializer/json"
)

func ConfigureContainer(container container.ServiceContainer, config config.ServiceConfig) {
	container.BindInstance(config).
		BindSingleton(factory.ProvideRepository).

		BindSingleton(func() core.Serializer { return json.NewSerializer()}).

		BindSingleton(async.NewHandler).
		BindSingleton(rpc.NewHandler).

		BindTransient(factory.ProvideServer)
}
