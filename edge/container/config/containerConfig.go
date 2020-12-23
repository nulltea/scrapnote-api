package config

import (
	"go.kicksware.com/api/service-common/container"

	"github.com/timoth-y/scrapnote-api/record/config"
	"github.com/timoth-y/scrapnote-api/record/container/factory"
)

func ConfigureContainer(container container.ServiceContainer, config config.ServiceConfig) {
	container.BindInstance(config).
		BindSingleton(factory.ProvideRepository).

		BindSingleton(factory.ProvideDataService).

		BindSingleton(factory.ProvideEdgeHandler).
		BindSingleton(factory.ProvideEndpointRouter).

		BindTransient(factory.ProvideServer)
}
