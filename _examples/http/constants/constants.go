package constants

import "time"

const (
	Port           = 8080
	IdleTimeout    = time.Minute
	RequestTimeout = time.Second * 3
	ReadTimeout    = 10 * time.Second
	WriteTimeout   = 30 * time.Second
)
