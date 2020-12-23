package repo

import "github.com/timoth-y/scrapnote-api/records/core/model"

type RecordRepository interface {
	RetrieveOne(id string) (*model.Record, error)
	Retrieve(topic string) ([]*model.Record, error)
	Store(record *model.Record) error
	Modify(record *model.Record) error
	Remove(id string)  error
}