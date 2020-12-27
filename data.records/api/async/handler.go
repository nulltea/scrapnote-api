package async

import (
	"encoding/json"
	"fmt"

	"github.com/golang/glog"
	"github.com/streadway/amqp"
	"go.kicksware.com/api/service-common/api/events"
	"go.kicksware.com/api/service-common/core"

	"github.com/timoth-y/scrapnote-api/data.records/config"
	"github.com/timoth-y/scrapnote-api/data.records/core/model"
	"github.com/timoth-y/scrapnote-api/data.records/core/service"
)

type handler struct {
	broker *events.Broker
	repo   service.RecordService
	errors chan error
}

func NewHandler(service service.RecordService, serializer core.Serializer, config config.ServiceConfig) core.Handler {
	return &handler{
		broker: events.NewEventsBroker(config.Events, "amq.topic", serializer),
		repo:   service,
		errors: make(chan error),
	}
}

func (h *handler) Setup() {
	if err := h.broker.Consume("records.add", "records.add", h.addHandler, 1); err != nil {
		glog.Fatalln(err)
	}
	if err := h.broker.Consume("rating.update", "rating.update", h.updateHandler, 1); err != nil {
		glog.Fatalln(err)
	}
	if err := h.broker.Consume("rating.delete", "rating.delete", h.deleteHandler, 1); err != nil {
		glog.Fatalln(err)
	}
}

func (h *handler) Serve() error {
	fmt.Println("Event listeners active...")
	err := <- h.errors
	return err
}

func (h *handler) addHandler(msg amqp.Delivery) bool {
	record, ok := getRecord(msg.Body); if !ok {
		return false
	}
	if err := h.repo.Add(record); err != nil {
		h.errors <- err
		return false
	}
	fmt.Printf("add event handled for: %q\n", record.Content)
	return true
}

func (h *handler) updateHandler(msg amqp.Delivery) bool {
	record, ok := getRecord(msg.Body); if !ok {
		return false
	}
	if err := h.repo.Update(record); err != nil {
		h.errors <- err
		return false
	}
	fmt.Printf("update event handled for: %q\n", record.Content)
	return true
}

func (h *handler) deleteHandler(msg amqp.Delivery) bool {
	record, ok := getRecord(msg.Body); if !ok {
		return false
	}
	if err := h.repo.Update(record); err != nil {
		h.errors <- err
		return false
	}
	fmt.Printf("delete event handled for: %q\n", record.Content)
	return true
}

func getRecord(data []byte) (*model.Record, bool) {
	var rec *model.Record
	if err := json.Unmarshal(data, &rec); err != nil {
		glog.Errorln(err)
		return nil, false
	}
	return rec, true
}