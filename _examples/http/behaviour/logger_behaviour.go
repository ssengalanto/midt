package behaviour

import (
	"context"
	"log"

	"github.com/ssengalanto/midt"
)

type LoggerBehaviour struct{}

func NewLoggerBehaviour() *LoggerBehaviour {
	return &LoggerBehaviour{}
}

func (l *LoggerBehaviour) Handle(
	ctx context.Context,
	request any,
	next midt.RequestHandlerFunc,
) (any, error) {
	log.Printf("executing %T", request)

	res, err := next()
	if err != nil {
		return nil, err
	}

	return res, nil
}
