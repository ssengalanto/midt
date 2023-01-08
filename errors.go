package mediatr

import "fmt"

// Errors used by the mediatr package.
var (
	// ErrRequestHandlerAlreadyExists is returned when a request handler already exists in the registry.
	ErrRequestHandlerAlreadyExists = fmt.Errorf("request handler already exists")
)
