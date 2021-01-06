package repo

import (
	"github.com/timoth-y/scrapnote-api/data.users/api/rpc/proto"
	"github.com/timoth-y/scrapnote-api/data.users/core/model"
)

type UserRepository interface {
	Retrieve(ids []string) ([]*model.User, error)
	RetrieveBy(filter *proto.UserFilter) ([]*model.User, error)
	Store(user *model.User) error
	Modify(user *model.User) error
	Remove(id string) error
}