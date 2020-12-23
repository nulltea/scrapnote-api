package factory

import (
	"github.com/golang/glog"

	"github.com/timoth-y/scrapnote-api/edge/config"
	"github.com/timoth-y/scrapnote-api/edge/core/repo"
	"github.com/timoth-y/scrapnote-api/edge/usecase/storage/mongo"
)

func ProvideRepository(config config.ServiceConfig) repo.RecordRepository {
	repo, err := mongo.NewRepository(config.Mongo); if err != nil {
		glog.Fatal(err)
	}
	return repo
}
