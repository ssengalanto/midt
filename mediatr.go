// Copyright 2023 Ssen Galanto. All rights reserved.

// Package mediatr provides mediator pattern that reduces coupling between components
// of a program by making them communicate indirectly, through a special mediator object.
package mediatr

import (
	"context"
	"fmt"
	"reflect"
)

type Mediatr struct {
	requestHandlerRegistry      map[string]RequestHandler
	notificationHandlerRegistry map[string]NotificationHandler
	pipelineBehaviourRegistry   []PipelineBehaviour
}

type RequestHandler interface {
	Name() string
	Handle(ctx context.Context, request any) (any, error)
}

type NotificationHandler interface {
	Name() string
	Handle(ctx context.Context, notification any) error
}

type RequestHandlerFunc func() (any, error)

type PipelineBehaviour interface {
	Handle(ctx context.Context, request any, next RequestHandlerFunc) (any, error)
}

// New creates a new Mediatr instance.
func New() *Mediatr {
	return &Mediatr{
		requestHandlerRegistry:      map[string]RequestHandler{},
		notificationHandlerRegistry: map[string]NotificationHandler{},
		pipelineBehaviourRegistry:   []PipelineBehaviour{},
	}
}

// RegisterRequestHandler register a RequestHandler in the registry.
func (m *Mediatr) RegisterRequestHandler(handler RequestHandler) error {
	hn := handler.Name()

	_, ok := m.requestHandlerRegistry[hn]
	if ok {
		return fmt.Errorf("%w: %s", ErrRequestHandlerAlreadyExists, hn)
	}

	m.requestHandlerRegistry[hn] = handler
	return nil
}

// RegisterNotificationHandler register a NotificationHandler in the registry.
func (m *Mediatr) RegisterNotificationHandler(handler NotificationHandler) error {
	hn := handler.Name()

	_, ok := m.notificationHandlerRegistry[hn]
	if ok {
		return fmt.Errorf("%w: %s", ErrNotificationHandlerAlreadyExists, hn)
	}

	m.notificationHandlerRegistry[hn] = handler
	return nil
}

// RegisterPipelineBehaviour register a PipelineBehaviour in the registry.
func (m *Mediatr) RegisterPipelineBehaviour(behaviour PipelineBehaviour) error {
	bt := reflect.TypeOf(behaviour)

	exists := m.existsPipeType(bt)
	if exists {
		return fmt.Errorf("%w: %s", ErrPipelineBehaviourAlreadyExists, bt)
	}

	m.pipelineBehaviourRegistry = prepend(m.pipelineBehaviourRegistry, behaviour)
	return nil
}

// existsPipeType checks if a pipeline behaviour exists in the registry.
func (m *Mediatr) existsPipeType(p reflect.Type) bool {
	for _, pipe := range m.pipelineBehaviourRegistry {
		if reflect.TypeOf(pipe) == p {
			return true
		}
	}
	return false
}

// prepend an element in the slice.
func prepend[T any](x []T, y T) []T {
	x = append([]T{y}, x...)
	return x
}
