package events

import (
	"context"

	"github.com/streadway/amqp"
)

func (b *Broker) Emmit(ctx context.Context, routingKey string, msg interface{}) error {
	event := EventMessage{
		Context: ctx,
		Payload: msg,
	}
	data, err := b.Serializer.Encode(event); if err != nil {
		return err
	}

	return b.Channel.Publish(
		b.Exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         data,
			DeliveryMode: amqp.Persistent,
		})
}