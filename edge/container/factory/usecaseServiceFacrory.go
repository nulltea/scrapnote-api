package factory

import (

	"github.com/timoth-y/scrapnote-api/edge/config"
	"github.com/timoth-y/scrapnote-api/edge/core/repo"
	"github.com/timoth-y/scrapnote-api/edge/core/service"
	"github.com/timoth-y/scrapnote-api/edge/usecase/business"
)

func ProvideDataService(repository repo.RecordRepository, config config.ServiceConfig) service.RecordService {
	return business.NewRecordService(repository, config)
}
