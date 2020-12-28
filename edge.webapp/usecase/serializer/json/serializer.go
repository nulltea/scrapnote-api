package json

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/timoth-y/scrapnote-api/data.records/core/model"

	"github.com/timoth-y/scrapnote-api/edge.webapp/core/service"
)

type serializer struct{}

func NewSerializer() service.RecordSerializer {
	return &serializer{}
}

func (r *serializer) Decode(input []byte) (ref *model.Record, err error) {
	if err = json.Unmarshal(input, &ref); err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerReference.Decode")
	}
	return
}

func (r *serializer) DecodeRange(input []byte) (refs []*model.Record, err error) {
	if err = json.Unmarshal(input, &refs); err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerReference.DecodeRange")
	}
	return
}

func (r *serializer) DecodeMap(input []byte) (map[string]interface{}, error) {
	queryMap := make(map[string]interface{})
	if err := json.Unmarshal(input, &queryMap); err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerReference.DecodeRange")
	}
	return queryMap, nil
}

func (r *serializer) DecodeInto(input []byte, target interface{}) error  {
	if err := json.Unmarshal(input, target); err != nil {
		return errors.Wrap(err, "serializer.SneakerReference.DecodeInto")
	}
	return nil
}

func (r *serializer) Encode(input interface{}) ([]byte, error) {
	raw, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerReference.Encode")
	}
	return raw, nil
}