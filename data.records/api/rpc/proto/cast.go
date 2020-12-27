package proto

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/timoth-y/scrapnote-api/data.records/core/model"
)

func (m *Record) ToNative() *model.Record {
	return &model.Record{
		UniqueID: m.UniqueID,
		Content: m.Content,
		SourceURL: m.SourceURL,
		MarkerURL: m.MarkerURL,
		AddedAt: m.AddedAt.AsTime(),
	}
}

func (m *Record) FromNative(n *model.Record) *Record {
	m.UniqueID = n.UniqueID
	m.Content = n.Content
	m.SourceURL = n.SourceURL
	m.MarkerURL = n.MarkerURL
	m.AddedAt = timestamppb.New(n.AddedAt)
	return m
}

func NativeToRecords(native []*model.Record) []*Record {
	users := make([]*Record, 0)
	for _, user := range native {
		if user == nil {
			continue;
		}
		users = append(users, (&Record{}).FromNative(user))
	}
	return users
}

func RecordsToNative(in []*Record) []*model.Record {
	users := make([]*model.Record, 0)
	for _, user := range in {
		users = append(users, user.ToNative())
	}
	return users
}