package business

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/rs/xid"
	"github.com/timoth-y/scrapnote-api/data.users/api/rpc/proto"
	"github.com/timoth-y/scrapnote-api/data.users/core/model"
	"github.com/timoth-y/scrapnote-api/lib.common/api/events"
	"github.com/timoth-y/scrapnote-api/lib.common/api/gRPC"
	"github.com/timoth-y/scrapnote-api/lib.common/core"

	"github.com/timoth-y/scrapnote-api/serv.auth/config"
	"github.com/timoth-y/scrapnote-api/serv.auth/core/errors"
	"github.com/timoth-y/scrapnote-api/serv.auth/core/service"
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

func (s *userService) Fetch(ctx context.Context, ids []string) ([]*model.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	resp, err := s.remote.Get(ctx, &proto.UserFilter{ UserID: ids })
	if err != nil {
		return nil, err
	}
	return proto.UsersToNative(resp.Users), nil
}

func (s *userService) FetchOne(ctx context.Context, id string) (*model.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	resp, err := s.remote.Get(ctx, &proto.UserFilter{ UserID: []string { id } })
	if err != nil || len(resp.Users) == 0 {
		return nil, err
	}
	return resp.Users[0].ToNative(), nil
}

func (s *userService) FetchByEmail(ctx context.Context, email string) (*model.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	resp, err := s.remote.Get(ctx, &proto.UserFilter{ Email: []string { email } })
	if err != nil || len(resp.Users) == 0 {
		return nil, err
	}
	return resp.Users[0].ToNative(), nil
}

func (s *userService) FetchByUsername(ctx context.Context, username string) (*model.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	resp, err := s.remote.Get(ctx, &proto.UserFilter{ Username: []string { username } })
	if err != nil || len(resp.Users) == 0 {
		return nil, err
	}
	return resp.Users[0].ToNative(), nil
}

func (s *userService) Create(ctx context.Context, user *model.User) error {
	user.RegisterDate = time.Now()
	if len(user.UniqueID) < 8 {
		user.UniqueID = xid.NewWithTime(user.RegisterDate).String()
	}
	if len(user.Username) == 0 {
		s.GenerateUsername(ctx, user, false)
	}
	return s.events.Emmit(ctx, "users.add", user)
}

func (s *userService) Modify(ctx context.Context, user *model.User) error {
	return s.events.Emmit(ctx, "users.update", user)
}

func (s *userService) Delete(ctx context.Context, user *model.User) error {
	return s.events.Emmit(ctx, "users.delete", user)
}

func (s *userService) Verify(ctx context.Context, user *model.User) error {
	token := generateToken(user.UniqueID)
	requestParams := &struct {
		Email       string
		CallbackURL string
	}{
		Email: user.Email,
		CallbackURL: fmt.Sprintf("/%v", url.PathEscape(token)),
	}
	return s.events.Emmit(ctx, "email.verify", requestParams)
}

func (s *userService) Confirm(ctx context.Context, userID, token string) error {
	user, err := s.FetchOne(ctx, userID); if err != nil {
		return err
	}

	if !verifyToken(userID, token) {
		return errors.ErrTokenInvalid
	}

	user.Confirmed = true
	if err = s.Modify(ctx, user); err != nil {
		return err
	}
	return nil
}

func (s *userService) GenerateUsername(ctx context.Context, user *model.User, save bool) (string, error) {
	if len(user.Email) == 0 {
		return "", errors.ErrEmailInvalid
	}
	username := strings.Split(user.Email, "@")[0]
	if another, _ := s.FetchByUsername(ctx, username); another != nil {
		baseUsername := username
		for another != nil {
			rand.Seed(user.RegisterDate.Unix())
			username = fmt.Sprintf("%v_%v", baseUsername, strconv.Itoa(rand.Int())[:3])
			another, _ = s.FetchByUsername(ctx, username)
		}
	}
	user.Username = username
	if save {
		if err := s.Modify(ctx, user); err != nil {
			return "", err
		}
	}
	return username, nil;
}


func generateToken(userID string) string {
	h := md5.New()
	h.Write([]byte(strings.ToLower(userID)))
	return hex.EncodeToString(h.Sum(nil))
}

func verifyToken(userID, token string) bool {
	return generateToken(userID) == token
}