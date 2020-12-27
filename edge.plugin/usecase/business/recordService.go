package business

import (
	"github.com/timoth-y/scrapnote-api/records/core/model"
	"go.kicksware.com/api/service-common/api/events"
	"go.kicksware.com/api/service-common/core"

	"github.com/timoth-y/scrapnote-api/edge.plugin/config"
	"github.com/timoth-y/scrapnote-api/edge.plugin/core/service"
)

type recordService struct {
	events *events.Broker
	config config.ServiceConfig
}

func NewRecordService(config config.ServiceConfig, serializer core.Serializer) service.RecordService {
	return &recordService {
		events.NewEventsBroker(config.Events, "amq.topic", serializer),
		config,
	}
}

func (s *recordService) GetOne(id string) (*model.Record, error) {
	panic("implement me")
}

func (s *recordService) Get(topic string) ([]*model.Record, error) {
	panic("implement me")
}

func (s *recordService) Add(record *model.Record) error {
	return s.events.Emmit("records.add", record)
}

func (s *recordService) Update(record *model.Record) error {
	return s.events.Emmit("records.update", record)
}

func (s *recordService) Delete(id string) error {
	return s.events.Emmit("records.delete", id)
}