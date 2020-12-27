package repo

import "github.com/timoth-y/scrapnote-api/data.records/core/model"

type RecordRepository interface {
	Retrieve(ids []string) ([]*model.Record, error)
	RetrieveBy(topic string) ([]*model.Record, error)
	Store(record *model.Record) error
	Modify(record *model.Record) error
	Remove(id string)  error
}