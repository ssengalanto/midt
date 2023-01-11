package event

import (
	"context"
	"fmt"
	"log"
)

type PersonCreatedEventHandler struct{}

func NewPersonCreatedEventHandler() *PersonCreatedEventHandler {
	return &PersonCreatedEventHandler{}
}

func (p *PersonCreatedEventHandler) Name() string {
	return fmt.Sprintf("%T", &PersonCreatedEvent{})
}

func (p *PersonCreatedEventHandler) Handle(ctx context.Context, notification any) error {
	evt := notification.(*PersonCreatedEvent)

	log.Printf("person created %+v", evt)

	return nil
}
