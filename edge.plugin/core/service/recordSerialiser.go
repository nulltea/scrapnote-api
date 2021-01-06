package service

import "github.com/timoth-y/scrapnote-api/data.records/core/model"

type RecordSerializer interface {
	Decode(input []byte) (*model.Record, error)
	DecodeRange(input []byte) ([]*model.Record, error)
	DecodeMap(input []byte) (map[string]interface{}, error)
	DecodeInto(input []byte, target interface{}) error
	Encode(input interface{}) ([]byte, error)
}