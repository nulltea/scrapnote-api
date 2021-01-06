package async

import (
	"context"
	"fmt"

	"github.com/golang/glog"
	"github.com/rs/xid"
	"github.com/timoth-y/scrapnote-api/lib.common/api/events"
	"github.com/timoth-y/scrapnote-api/lib.common/api/rest"
	"github.com/timoth-y/scrapnote-api/lib.common/core"
	"github.com/timoth-y/scrapnote-api/lib.common/core/errors"
	"github.com/timoth-y/scrapnote-api/lib.common/core/meta"

	"github.com/timoth-y/scrapnote-api/data.records/config"
	"github.com/timoth-y/scrapnote-api/data.records/core/model"
	"github.com/timoth-y/scrapnote-api/data.records/core/repo"
)

type handler struct {
	broker *events.Broker
	repo   repo.RecordRepository
	errors chan error
}

func NewHandler(repo repo.RecordRepository, serializer core.Serializer, config config.ServiceConfig) core.Handler {
	return &handler{
		broker: events.NewEventsBroker(config.Events, "amq.topic", serializer),
		repo:   repo,
		errors: make(chan error),
	}
}

func (h *handler) Setup() {
	if err := h.broker.Consume("records.add", "records.add", h.addHandler, 1); err != nil {
		glog.Fatalln(err)
	}
	if err := h.broker.Consume("records.update", "records.update", h.updateHandler, 1); err != nil {
		glog.Fatalln(err)
	}
	if err := h.broker.Consume("records.delete", "records.delete", h.deleteHandler, 1); err != nil {
		glog.Fatalln(err)
	}
}

func (h *handler) Serve() error {
	fmt.Println("Event listeners active...")
	err := <- h.errors
	return err
}

func (h *handler) addHandler(ctx context.Context, msg interface{}) bool {
	record, ok := msg.(*model.Record); if !ok {
		return false
	}
	user, ok := ctx.Value(rest.UserContextKey).(meta.UserContextInfo); if !ok {
		h.errors <- errors.ErrUserContextInfoMissing
		return false
	}
	record.UniqueID = xid.New().String()
	record.UserID = user.UniqueID
	if err := h.repo.Store(ctx, record); err != nil {
		h.errors <- err
		return false
	}
	fmt.Printf("add event handled for: %q\n", record.Content)
	return true
}

func (h *handler) updateHandler(ctx context.Context, msg interface{}) bool {
	record, ok := msg.(*model.Record); if !ok {
		return false
	}
	user, ok := ctx.Value(rest.UserContextKey).(meta.UserContextInfo); if !ok {
		h.errors <- errors.ErrUserContextInfoMissing
		return false
	}
	if record.UniqueID != user.UniqueID {
		h.errors <- errors.ErrWrongUserContext
		return false
	}
	if err := h.repo.Modify(ctx, record); err != nil {
		h.errors <- err
		return false
	}
	fmt.Printf("update event handled for: %q\n", record.Content)
	return true
}

func (h *handler) deleteHandler(ctx context.Context, msg interface{}) bool {
	record, ok := msg.(*model.Record); if !ok {
		return false
	}
	user, ok := ctx.Value(rest.UserContextKey).(meta.UserContextInfo); if !ok {
		h.errors <- errors.ErrUserContextInfoMissing
		return false
	}
	if record.UniqueID != user.UniqueID {
		h.errors <- errors.ErrWrongUserContext
		return false
	}
	if err := h.repo.Remove(ctx, record.UniqueID); err != nil {
		h.errors <- err
		return false
	}
	fmt.Printf("delete event handled for: %q\n", record.Content)
	return true
}