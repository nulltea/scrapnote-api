package service

import "github.com/timoth-y/scrapnote-api/data.users/core/model"

type UserService interface {
	Fetch(ids []string) ([]*model.User, error)
	FetchOne(id string) (*model.User, error)
	FetchByEmail(email string) (*model.User, error)
	FetchByUsername(username string) (*model.User, error)
	Create(user *model.User) error
	Modify(user *model.User) error
	Delete(user *model.User) error
}