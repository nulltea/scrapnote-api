package factory

import (
	"github.com/golang/glog"

	"github.com/timoth-y/scrapnote-api/data.users/config"
	"github.com/timoth-y/scrapnote-api/data.users/core/repo"
	"github.com/timoth-y/scrapnote-api/data.users/usecase/storage/mongo"
)

func ProvideRepository(config config.ServiceConfig) repo.UserRepository {
	repo, err := mongo.NewRepository(config.Mongo); if err != nil {
		glog.Fatal(err)
	}
	return repo
}