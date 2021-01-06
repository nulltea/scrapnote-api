package events

import "context"

type EventMessage struct {
	Context context.Context
	Payload interface{}
}