package api

import (
	"context"
	"fmt"

	"github.com/golang/glog"
	"go.kicksware.com/api/service-common/api/events"
	"go.kicksware.com/api/service-common/core"

	"github.com/timoth-y/scrapnote-api/serv.email/config"
	"github.com/timoth-y/scrapnote-api/serv.email/core/model"
	"github.com/timoth-y/scrapnote-api/serv.email/core/service"
)

type handler struct {
	broker *events.Broker
	service service.MailService
	errors chan error
}

func NewHandler(service service.MailService, serializer core.Serializer, config config.ServiceConfig) core.Handler {
	return &handler{
		broker: events.NewEventsBroker(config.Events, "amq.topic", serializer),
		service: service,
		errors: make(chan error),
	}
}

func (h *handler) Setup() {
	if err := h.broker.Consume("email.verify", "email.verify", h.verifyHandler, 1); err != nil {
		glog.Fatalln(err)
	}
	if err := h.broker.Consume("email.reset", "records.update", h.resetHandler, 1); err != nil {
		glog.Fatalln(err)
	}
	if err := h.broker.Consume("email.notify", "records.delete", h.notifyHandler, 1); err != nil {
		glog.Fatalln(err)
	}
}

func (h *handler) Serve() error {
	fmt.Println("Event listeners active...")
	err := <- h.errors
	return err
}

func (h *handler) verifyHandler(ctx context.Context, msg interface{}) bool {
	request, ok := msg.(*model.EmailRequest); if !ok {
		return false
	}
	if err := h.service.SendEmailConfirmation(ctx, request.Email, request.CallbackURL); err != nil {
		h.errors <- err
		return false
	}
	fmt.Printf("verify user email event handled for: %q\n", request.Email)
	return true
}

func (h *handler) resetHandler(ctx context.Context, msg interface{}) bool {
	request, ok := msg.(*model.EmailRequest); if !ok {
		return false
	}
	if err := h.service.SendResetPassword(ctx, request.Email, request.CallbackURL); err != nil {
		h.errors <- err
		return false
	}
	fmt.Printf("reset user password event handled for: %q\n", request.Email)
	return true
}

func (h *handler) notifyHandler(ctx context.Context, msg interface{}) bool {
	request, ok := msg.(*model.EmailRequest); if !ok {
		return false
	}
	if err := h.service.SendNotification(ctx, request.Email, request.Content); err != nil {
		h.errors <- err
		return false
	}
	fmt.Printf("notify user event handled for: %q\n", request.Email)
	return true
}
