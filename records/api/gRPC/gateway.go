package gRPC

import (
	"google.golang.org/grpc"

	"github.com/timoth-y/scrapnote-api/records/api/proto"
)

func ProvideRemoteSetup(handler *Handler) func(server *grpc.Server) {
	return func(server *grpc.Server) {
		proto.RegisterRecordServiceServer(server, handler)
	}
}