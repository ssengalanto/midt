// Copyright 2023 Ssen Galanto. All rights reserved.

// Package mediatr provides mediator pattern that reduces coupling between components
// of a program by making them communicate indirectly, through a special mediator object.
package mediatr

import (
	"context"
	"fmt"
)

type Mediatr struct {
	requestHandlerRegistry      map[string]RequestHandler
	notificationHandlerRegistry map[string]NotificationHandler
}

type RequestHandler interface {
	Name() string
	Handle(ctx context.Context, request any) (any, error)
}

type NotificationHandler interface {
	Name() string
	Handle(ctx context.Context, notification any) error
}

// New creates a new Mediatr instance.
func New() *Mediatr {
	return &Mediatr{
		requestHandlerRegistry:      map[string]RequestHandler{},
		notificationHandlerRegistry: map[string]NotificationHandler{},
	}
}

// RegisterRequestHandler register a request handler in the registry.
func (m *Mediatr) RegisterRequestHandler(handler RequestHandler) error {
	hn := handler.Name()

	_, ok := m.requestHandlerRegistry[hn]
	if ok {
		return fmt.Errorf("%w: %s", ErrRequestHandlerAlreadyExists, hn)
	}

	m.requestHandlerRegistry[hn] = handler
	return nil
}

// RegisterNotificationHandler register a notification handler in the registry.
func (m *Mediatr) RegisterNotificationHandler(handler NotificationHandler) error {
	hn := handler.Name()

	_, ok := m.notificationHandlerRegistry[hn]
	if ok {
		return fmt.Errorf("%w: %s", ErrNotificationHandlerAlreadyExists, hn)
	}

	m.notificationHandlerRegistry[hn] = handler
	return nil
}
