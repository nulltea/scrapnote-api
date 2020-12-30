package config

import (
	"go.kicksware.com/api/service-common/container"
	"go.kicksware.com/api/service-common/core"

	"github.com/timoth-y/scrapnote-api/edge.webapp/api"
	"github.com/timoth-y/scrapnote-api/edge.webapp/config"
	"github.com/timoth-y/scrapnote-api/edge.webapp/container/factory"
	"github.com/timoth-y/scrapnote-api/edge.webapp/usecase/business"
	"github.com/timoth-y/scrapnote-api/edge.webapp/usecase/serializer/json"
)

func ConfigureContainer(container container.ServiceContainer, config config.ServiceConfig) {
	container.BindInstance(config).
		BindSingleton(func() core.Serializer { return json.NewSerializer()}).
		BindSingleton(business.NewRecordService).

		BindSingleton(api.NewHandler).

		BindTransient(factory.ProvideServer)
}
