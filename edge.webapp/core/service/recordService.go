package service

import "github.com/timoth-y/scrapnote-api/data.records/core/model"

type RecordService interface {
	GetOne(id string) (*model.Record, error)
	GetAll() ([]*model.Record, error)
	GetFrom(topic string) ([]*model.Record, error)
	Add(record *model.Record) error
	Update(record *model.Record) error
	Delete(id string) error
}
