package business

import (
	"context"

	"github.com/timoth-y/scrapnote-api/data.users/api/rpc/proto"
	"github.com/timoth-y/scrapnote-api/data.users/core/model"
	"go.kicksware.com/api/service-common/api/events"
	"go.kicksware.com/api/service-common/api/gRPC"
	"go.kicksware.com/api/service-common/core"

	"github.com/timoth-y/scrapnote-api/serv.email/config"
	"github.com/timoth-y/scrapnote-api/serv.email/core/service"
)

type userService struct {
	events *events.Broker
	remote proto.UserServiceClient
	config     config.ServiceConfig
}

func NewUserService(config config.ServiceConfig, serializer core.Serializer) service.UserService {
	return &userService{
		events.NewEventsBroker(config.Events, "amq.topic", serializer),
		proto.NewUserServiceClient(gRPC.NewRemoteConnection(config.RPC)),
		config,
	}
}

func (s *userService) Fetch(ids []string) ([]*model.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	resp, err := s.remote.Get(ctx, &proto.UserFilter{ UserID: ids })
	if err != nil {
		return nil, err
	}
	return proto.UsersToNative(resp.Users), nil
}

func (s *userService) FetchOne(id string) (*model.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	resp, err := s.remote.Get(ctx, &proto.UserFilter{ UserID: []string { id } })
	if err != nil || len(resp.Users) == 0 {
		return nil, err
	}
	return resp.Users[0].ToNative(), nil
}

func (s *userService) FetchByEmail(email string) (*model.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	resp, err := s.remote.Get(ctx, &proto.UserFilter{ Email: []string { email } })
	if err != nil || len(resp.Users) == 0 {
		return nil, err
	}
	return resp.Users[0].ToNative(), nil
}

func (s *userService) FetchByUsername(username string) (*model.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	resp, err := s.remote.Get(ctx, &proto.UserFilter{ Username: []string { username } })
	if err != nil || len(resp.Users) == 0 {
		return nil, err
	}
	return resp.Users[0].ToNative(), nil
}

func (s *userService) Create(user *model.User) error {
	return s.events.Emmit("users.add", user)
}

func (s *userService) Modify(user *model.User) error {
	return s.events.Emmit("users.update", user)
}

func (s *userService) Delete(user *model.User) error {
	return s.events.Emmit("users.delete", user)
}