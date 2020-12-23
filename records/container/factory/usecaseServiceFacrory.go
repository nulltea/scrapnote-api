package factory

import (

	"github.com/timoth-y/scrapnote-api/records/config"
	"github.com/timoth-y/scrapnote-api/records/core/repo"
	"github.com/timoth-y/scrapnote-api/records/core/service"
	"github.com/timoth-y/scrapnote-api/records/usecase/business"
)

func ProvideDataService(repository repo.RecordRepository, config config.ServiceConfig) service.RecordService {
	return business.NewRecordService(repository, config)
}
