package gRPC

import (
	"context"

	"go.kicksware.com/api/service-common/core"

	"github.com/timoth-y/scrapnote-api/records/api/gRPC/proto"
	"github.com/timoth-y/scrapnote-api/records/core/service"
)

//go:generate protoc --proto_path=../../protos --go_out=plugins=grpc,paths=source_relative:proto/. record.proto

type Handler struct {
	service     service.RecordService
	auth        core.AuthService
}

func NewHandler(service service.RecordService, auth core.AuthService) *Handler {
	return &Handler{
		service,
		auth,
	}
}

func (h Handler) Get(ctx context.Context, filter *proto.RecordFilter) (*proto.RecordResponse, error) {
	panic("implement me")
}

func (h Handler) Count(ctx context.Context, filter *proto.RecordFilter) (*proto.RecordResponse, error) {
	panic("implement me")
}

func (h Handler) Post(ctx context.Context, record *proto.Record) (*proto.RecordResponse, error) {
	panic("implement me")
}

func (h Handler) Edit(ctx context.Context, record *proto.Record) (*proto.RecordResponse, error) {
	panic("implement me")
}

func (h Handler) Delete(ctx context.Context, filter *proto.RecordFilter) (*proto.RecordResponse, error) {
	panic("implement me")
}