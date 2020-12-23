package factory

import (

	"github.com/timoth-y/scrapnote-api/record/config"
	"github.com/timoth-y/scrapnote-api/record/core/repo"
	"github.com/timoth-y/scrapnote-api/record/core/service"
	"github.com/timoth-y/scrapnote-api/record/usecase/business"
)

func ProvideDataService(repository repo.RecordRepository, config config.ServiceConfig) service.RecordService {
	return business.NewRecordService(repository, config)
}
