package service

import (
	"context"

	"github.com/timoth-y/scrapnote-api/data.users/core/model"
)

type UserService interface {
	Fetch(ctx context.Context, ids []string) ([]*model.User, error)
	FetchOne(ctx context.Context, id string) (*model.User, error)
	FetchByEmail(ctx context.Context, email string) (*model.User, error)
	FetchByUsername(ctx context.Context, username string) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	Modify(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, user *model.User) error
}