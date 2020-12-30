package config

import (
	"go.kicksware.com/api/service-common/container"
	"go.kicksware.com/api/service-common/core"

	"github.com/timoth-y/scrapnote-api/data.users/api/async"
	"github.com/timoth-y/scrapnote-api/data.users/api/rpc"
	"github.com/timoth-y/scrapnote-api/data.users/config"
	"github.com/timoth-y/scrapnote-api/data.users/container/factory"
	"github.com/timoth-y/scrapnote-api/data.users/usecase/serializer/json"
)

func ConfigureContainer(container container.ServiceContainer, config config.ServiceConfig) {
	container.BindInstance(config).
		BindSingleton(factory.ProvideRepository).

		BindSingleton(func() core.Serializer { return json.NewSerializer()}).

		BindSingleton(async.NewHandler).
		BindSingleton(rpc.NewHandler).

		BindTransient(factory.ProvideServer)
}
