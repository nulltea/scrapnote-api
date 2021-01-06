package service

import (
	"context"

	"github.com/timoth-y/scrapnote-api/data.records/core/model"
)

type RecordService interface {
	GetOne(ctx context.Context, id string) (*model.Record, error)
	Get(ctx context.Context, topic string) ([]*model.Record, error)
	Add(ctx context.Context, record *model.Record) error
	Update(ctx context.Context, record *model.Record) error
	Delete(ctx context.Context, id string) error
}
