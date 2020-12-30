package config

import (
	"go.kicksware.com/api/service-common/container"

	"github.com/timoth-y/scrapnote-api/serv.auth/api"
	"github.com/timoth-y/scrapnote-api/serv.auth/config"
	"github.com/timoth-y/scrapnote-api/serv.auth/container/factory"
	"github.com/timoth-y/scrapnote-api/serv.auth/usecase/business"
)

func ConfigureContainer(container container.ServiceContainer, config config.ServiceConfig) {
	container.BindInstance(config).
		BindSingleton(business.NewUserService).
		BindSingleton(business.NewAuthServiceJWT).

		BindSingleton(api.NewHandler).
		BindSingleton(api.ProvideRoutes).

		BindTransient(factory.ProvideServer)
}
