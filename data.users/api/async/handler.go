package async

import (
	"encoding/json"
	"fmt"

	"github.com/golang/glog"
	"github.com/streadway/amqp"
	"go.kicksware.com/api/service-common/api/events"
	"go.kicksware.com/api/service-common/core"

	"github.com/timoth-y/scrapnote-api/data.users/config"
	"github.com/timoth-y/scrapnote-api/data.users/core/model"
	"github.com/timoth-y/scrapnote-api/data.users/core/repo"
)

type handler struct {
	broker *events.Broker
	repo   repo.UserRepository
	errors chan error
}

func NewHandler(repo repo.UserRepository, serializer core.Serializer, config config.ServiceConfig) core.Handler {
	return &handler{
		broker: events.NewEventsBroker(config.Events, "amq.topic", serializer),
		repo:   repo,
		errors: make(chan error),
	}
}

func (h *handler) Setup() {
	if err := h.broker.Consume("users.add", "users.add", h.addHandler, 1); err != nil {
		glog.Fatalln(err)
	}
	if err := h.broker.Consume("users.update", "users.update", h.updateHandler, 1); err != nil {
		glog.Fatalln(err)
	}
	if err := h.broker.Consume("users.delete", "users.delete", h.deleteHandler, 1); err != nil {
		glog.Fatalln(err)
	}
}

func (h *handler) Serve() error {
	fmt.Println("Event listeners active...")
	err := <- h.errors
	return err
}

func (h *handler) addHandler(msg amqp.Delivery) bool {
	user, ok := getUser(msg.Body); if !ok {
		return false
	}
	if err := h.repo.Store(user); err != nil {
		h.errors <- err
		return false
	}
	fmt.Printf("add event handled for: %q\n", user.Username)
	return true
}

func (h *handler) updateHandler(msg amqp.Delivery) bool {
	user, ok := getUser(msg.Body); if !ok {
		return false
	}
	if err := h.repo.Modify(user); err != nil {
		h.errors <- err
		return false
	}
	fmt.Printf("update event handled for: %q\n", user.Username)
	return true
}

func (h *handler) deleteHandler(msg amqp.Delivery) bool {
	user, ok := getUser(msg.Body); if !ok {
		return false
	}
	if err := h.repo.Remove(user.UniqueID); err != nil {
		h.errors <- err
		return false
	}
	fmt.Printf("delete event handled for: %q\n", user.Username)
	return true
}

func getUser(data []byte) (*model.User, bool) {
	var rec *model.User
	if err := json.Unmarshal(data, &rec); err != nil {
		glog.Errorln(err)
		return nil, false
	}
	return rec, true
}