package business

import (
	"context"

	"github.com/timoth-y/scrapnote-api/data.records/api/rpc/proto"
	"github.com/timoth-y/scrapnote-api/data.records/core/model"
	"go.kicksware.com/api/service-common/api/events"
	"go.kicksware.com/api/service-common/api/gRPC"
	"go.kicksware.com/api/service-common/core"

	"github.com/timoth-y/scrapnote-api/edge.webapp/config"
	"github.com/timoth-y/scrapnote-api/edge.webapp/core/service"
)

type recordService struct {
	events *events.Broker
	remote proto.RecordServiceClient
	config config.ServiceConfig
}

func NewRecordService(config config.ServiceConfig, serializer core.Serializer) service.RecordService {
	return &recordService {
		events.NewEventsBroker(config.Events, "amq.topic", serializer),
		proto.NewRecordServiceClient(gRPC.NewRemoteConnection(config.RPC)),
		config,
	}
}

func (s *recordService) GetOne(id string) (*model.Record, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resp, err := s.remote.Get(ctx, &proto.RecordFilter {
		RecordID: []string{ id },
	}); if err != nil || len(resp.Records) == 0 {
		return nil, err
	}

	return resp.Records[0].ToNative(), nil
}

func (s *recordService) GetAll() ([]*model.Record, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resp, err := s.remote.Get(ctx, &proto.RecordFilter{})
	if err != nil || len(resp.Records) == 0 {
		return nil, err
	}

	return proto.RecordsToNative(resp.Records), nil
}

func (s *recordService) GetFrom(topic string) ([]*model.Record, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resp, err := s.remote.Get(ctx, &proto.RecordFilter {
		TopicID: topic,
	}); if err != nil || len(resp.Records) == 0 {
		return nil, err
	}

	return proto.RecordsToNative(resp.Records), nil
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