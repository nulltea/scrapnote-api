package rpc

import (
	"context"

	"github.com/timoth-y/scrapnote-api/lib.common/core"
	"google.golang.org/grpc"

	"github.com/timoth-y/scrapnote-api/data.users/api/rpc/proto"
	"github.com/timoth-y/scrapnote-api/data.users/core/model"
	"github.com/timoth-y/scrapnote-api/data.users/core/repo"
)

//go:generate protoc --proto_path=proto/ --go_out=plugins=grpc,paths=source_relative:proto/. user.proto

type Handler struct {
	repo repo.UserRepository
	auth core.AuthService
}

func ProvideRemoteSetup(handler *Handler) func(server *grpc.Server) {
	return func(server *grpc.Server) {
		proto.RegisterUserServiceServer(server, handler)
	}
}

func NewHandler(repo repo.UserRepository) *Handler {
	return &Handler{
		repo,
		nil,
	}
}

func (h Handler) Get(ctx context.Context, filter *proto.UserFilter) (r *proto.UserResponse, err error) {
	var users []*model.User

	users, err = h.repo.RetrieveBy(filter); if err != nil {
		return nil, err
	}

	r = &proto.UserResponse{
		Users: proto.NativeToUsers(users),
		Count: int64(len(users)),
	}
	return
}

func (h Handler) Count(ctx context.Context, filter *proto.UserFilter) (*proto.UserResponse, error) {
	panic("implement me")
}
