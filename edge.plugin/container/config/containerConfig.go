package config

import (
	"github.com/timoth-y/scrapnote-api/lib.common/api/rest"
	"github.com/timoth-y/scrapnote-api/lib.common/container"
	"github.com/timoth-y/scrapnote-api/lib.common/core"

	"github.com/timoth-y/scrapnote-api/edge.plugin/config"
	"github.com/timoth-y/scrapnote-api/edge.plugin/container/factory"
	"github.com/timoth-y/scrapnote-api/edge.plugin/usecase/business"
	"github.com/timoth-y/scrapnote-api/edge.plugin/usecase/serializer/json"
)

func ConfigureContainer(container container.ServiceContainer, config config.ServiceConfig) {
	container.BindInstance(config).
		BindSingleton(func() core.Serializer { return json.NewSerializer()}).
		BindSingleton(func() core.AuthService { return rest.NewAuthService(config.Auth)}).
		BindSingleton(business.NewRecordService).

		BindSingleton(factory.ProvideEdgeHandler).
		BindSingleton(factory.ProvideEndpointRouter).

		BindTransient(factory.ProvideServer)
}
