package business

import (
	"github.com/timoth-y/scrapnote-api/records/config"
	"github.com/timoth-y/scrapnote-api/records/core/model"
	"github.com/timoth-y/scrapnote-api/records/core/repo"
	"github.com/timoth-y/scrapnote-api/records/core/service"
)

type recordService struct {
	repo repo.RecordRepository
	config config.ServiceConfig
}

func NewRecordService(repo repo.RecordRepository, config config.ServiceConfig) service.RecordService {
	return &recordService {
		repo,
		config,
	}
}

func (s *recordService) GetOne(id string) (*model.Record, error) {
	return s.repo.RetrieveOne(id)
}

func (s *recordService) Get(topic string) ([]*model.Record, error) {
	return s.repo.Retrieve(topic)
}

func (s *recordService) Add(record *model.Record) error {
	return s.repo.Store(record)
}

func (s *recordService) Update(record *model.Record) error {
	return s.repo.Modify(record)
}

func (s *recordService) Delete(id string) error {
	return s.repo.Remove(id)
}