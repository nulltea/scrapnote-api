package repo

import (
	"context"

	"github.com/timoth-y/scrapnote-api/data.records/api/rpc/proto"
	"github.com/timoth-y/scrapnote-api/data.records/core/model"
)

type RecordRepository interface {
	Retrieve(ctx context.Context, ids []string) ([]*model.Record, error)
	RetrieveBy(ctx context.Context, filter *proto.RecordFilter) ([]*model.Record, error)
	RetrieveAll(ctx context.Context) ([]*model.Record, error)
	Store(ctx context.Context, record *model.Record) error
	Modify(ctx context.Context, record *model.Record) error
	Remove(ctx context.Context, id string)  error
}