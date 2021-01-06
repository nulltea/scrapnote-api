package events

import (
	"context"

	"github.com/golang/glog"
	"github.com/streadway/amqp"
)

func (b *Broker) Consume(
	queueName,
	routingKey string,
	handler func(ctx context.Context, msg interface{}) bool,
	concurrency int) error {

	_, err := b.Channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil); if err != nil {
		return err
	}

	err = b.Channel.QueueBind(
		queueName,
		routingKey,
		b.Exchange,
		false,
		nil); if err != nil {
		return err
	}

	// prefetch 4x as many messages as we can handle at once
	prefetchCount := concurrency * 4
	err = b.Channel.Qos(prefetchCount, 0, false)
	if err != nil {
		return err
	}

	deliveries, err := b.Channel.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	); if err != nil {
		return err
	}

	// create a goroutine for the number of concurrent threads requested
	for i := 0; i < concurrency; i++ {
		go func() {
			for d := range deliveries {
				msg, err := b.decodeDelivery(d); if err != nil {
					d.Nack(false, false)
					continue
				}
				if handler(msg.Context, msg.Payload) {
					d.Ack(false)
				} else {
					d.Nack(false, true)
				}
			}
			glog.Fatalln("consume error: consumer closed ")
		}()
	}
	return nil
}

func (b *Broker) decodeDelivery(d amqp.Delivery) (EventMessage, error) {
	event := EventMessage{}
	err := b.Serializer.DecodeInto(d.Body, &event); if err != nil {
		return event, err
	}
	return event, nil
}