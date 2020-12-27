package config

import (
	"go.kicksware.com/api/service-common/container"
	"go.kicksware.com/api/service-common/core"

	"github.com/timoth-y/scrapnote-api/edge.plugin/config"
	"github.com/timoth-y/scrapnote-api/edge.plugin/container/factory"
	"github.com/timoth-y/scrapnote-api/edge.plugin/usecase/business"
	"github.com/timoth-y/scrapnote-api/edge.plugin/usecase/serializer/json"
)

func ConfigureContainer(container container.ServiceContainer, config config.ServiceConfig) {
	container.BindInstance(config).
		BindSingleton(func() core.Serializer { return json.NewSerializer()}).
		BindSingleton(business.NewRecordService).

		BindSingleton(factory.ProvideEdgeHandler).
		BindSingleton(factory.ProvideEndpointRouter).

		BindTransient(factory.ProvideServer)
}
