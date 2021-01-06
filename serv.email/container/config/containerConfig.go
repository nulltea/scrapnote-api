package config

import (
	"github.com/timoth-y/scrapnote-api/data.users/usecase/serializer/json"
	"go.kicksware.com/api/service-common/container"
	"go.kicksware.com/api/service-common/core"

	"github.com/timoth-y/scrapnote-api/serv.email/api"
	"github.com/timoth-y/scrapnote-api/serv.email/config"
	"github.com/timoth-y/scrapnote-api/serv.email/container/factory"
	"github.com/timoth-y/scrapnote-api/serv.email/usecase/business"
)

func ConfigureContainer(container container.ServiceContainer, config config.ServiceConfig) {
	container.BindInstance(config).
		BindSingleton(func() core.Serializer { return json.NewSerializer()}).

		BindSingleton(business.NewUserService).
		BindSingleton(business.NewMailService).

		BindSingleton(api.NewHandler).

		BindTransient(factory.ProvideServer)
}
