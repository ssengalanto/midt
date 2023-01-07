// Copyright 2023 Ssen Galanto. All rights reserved.

// Package mediatr provides mediator pattern that reduces coupling between components
// of a program by making them communicate indirectly, through a special mediator object.
package mediatr

type Mediatr struct{}

// New creates a new Mediatr instance.
func New() *Mediatr {
	return &Mediatr{}
}
