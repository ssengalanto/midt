// Copyright 2023 Ssen Galanto. All rights reserved.

// Package midt provides mediator pattern that reduces coupling between components
// of a program by making them communicate indirectly, through a special mediator struct.
package midt

import (
	"context"
	"fmt"
	"reflect"
)

type RequestHandler interface {
	Name() string
	Handle(ctx context.Context, request any) (any, error)
}

type NotificationHandler interface {
	Name() string
	Handle(ctx context.Context, notification any) error
}

type PipelineBehaviour interface {
	Handle(ctx context.Context, request any, next RequestHandlerFunc) (any, error)
}

type RequestHandlerFunc func() (any, error)

type Midt struct {
	requestHandlerRegistry      map[string]RequestHandler
	notificationHandlerRegistry map[string]NotificationHandler
	pipelineBehaviourRegistry   []PipelineBehaviour
}

// New creates a new Midt instance.
func New() *Midt {
	return &Midt{
		requestHandlerRegistry:      map[string]RequestHandler{},
		notificationHandlerRegistry: map[string]NotificationHandler{},
		pipelineBehaviourRegistry:   []PipelineBehaviour{},
	}
}

// RegisterRequestHandler register a RequestHandler in the registry.
func (m *Midt) RegisterRequestHandler(handler RequestHandler) error {
	hn := handler.Name()

	_, ok := m.requestHandlerRegistry[hn]
	if ok {
		return fmt.Errorf("%w: %s", ErrRequestHandlerAlreadyExists, hn)
	}

	m.requestHandlerRegistry[hn] = handler
	return nil
}

// RegisterNotificationHandler register a NotificationHandler in the registry.
func (m *Midt) RegisterNotificationHandler(handler NotificationHandler) error {
	hn := handler.Name()

	_, ok := m.notificationHandlerRegistry[hn]
	if ok {
		return fmt.Errorf("%w: %s", ErrNotificationHandlerAlreadyExists, hn)
	}

	m.notificationHandlerRegistry[hn] = handler
	return nil
}

// RegisterPipelineBehaviour register a PipelineBehaviour in the registry.
func (m *Midt) RegisterPipelineBehaviour(behaviour PipelineBehaviour) error {
	bt := reflect.TypeOf(behaviour)

	exists := m.existsPipeType(bt)
	if exists {
		return fmt.Errorf("%w: %s", ErrPipelineBehaviourAlreadyExists, bt)
	}

	m.pipelineBehaviourRegistry = append(m.pipelineBehaviourRegistry, behaviour)
	return nil
}

// Send sends the request to its corresponding RequestHandler.
func (m *Midt) Send(ctx context.Context, request any) (any, error) {
	rt := reflect.TypeOf(request).String()

	handler, ok := m.requestHandlerRegistry[rt]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrRequestHandlerNotFound, rt)
	}

	if len(m.pipelineBehaviourRegistry) > 0 {
		var lastHandler RequestHandlerFunc = func() (any, error) {
			return handler.Handle(ctx, request)
		}

		var aggregateResult = lastHandler
		for _, pipe := range m.pipelineBehaviourRegistry {
			pipeValue := pipe
			nextValue := aggregateResult

			aggregateResult = func() (any, error) {
				return pipeValue.Handle(ctx, request, nextValue)
			}
		}

		response, err := aggregateResult()
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	response, err := handler.Handle(ctx, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Publish publishes the notification event to its corresponding NotificationHandler.
func (m *Midt) Publish(ctx context.Context, request any) error {
	rt := reflect.TypeOf(request).String()

	handler, ok := m.notificationHandlerRegistry[rt]
	if !ok {
		return fmt.Errorf("%w: %s", ErrRequestHandlerNotFound, rt)
	}

	err := handler.Handle(ctx, request)
	if err != nil {
		return err
	}

	return nil
}

// existsPipeType checks if a pipeline behaviour exists in the registry.
func (m *Midt) existsPipeType(p reflect.Type) bool {
	for _, pipe := range m.pipelineBehaviourRegistry {
		if reflect.TypeOf(pipe) == p {
			return true
		}
	}
	return false
}
