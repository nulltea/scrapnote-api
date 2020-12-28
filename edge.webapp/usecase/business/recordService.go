package business

import (
	"github.com/golang/glog"
	"github.com/timoth-y/scrapnote-api/data.records/api/rpc/proto"
	"github.com/timoth-y/scrapnote-api/data.records/core/model"
	"google.golang.org/grpc"

	"go.kicksware.com/api/service-common/api/events"
	"go.kicksware.com/api/service-common/core"

	"github.com/timoth-y/scrapnote-api/edge.webapp/config"
	"github.com/timoth-y/scrapnote-api/edge.webapp/core/service"
)

type recordService struct {
	events *events.Broker
	remote *proto.RecordServiceClient
	config config.ServiceConfig
}

func NewRecordService(config config.ServiceConfig, serializer core.Serializer) service.RecordService {
	conn, err := grpc.Dial(serviceEndpoint, opts...); if err != nil {
		glog.Fatalf("fail to dial: %v", err)
	}

	return &recordService {
		events.NewEventsBroker(config.Events, "amq.topic", serializer),
		proto.NewRecordServiceClient(),
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