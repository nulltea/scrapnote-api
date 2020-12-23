package service

import "github.com/timoth-y/scrapnote-api/edge/core/model"

type RecordService interface {
	GetOne(id string) (*model.Record, error)
	Get(topic string) ([]*model.Record, error)
	Add(record *model.Record) error
	Update(record *model.Record) error
	Delete(id string) error
}
