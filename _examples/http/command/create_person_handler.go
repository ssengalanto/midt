package command

import (
	"context"
	"errors"
	"fmt"

	"github.com/ssengalanto/midt/_examples/http/dto"
	"github.com/ssengalanto/midt/_examples/http/memory"
)

type CreatePersonCommandHandler struct {
	personRepository memory.Repository
}

func NewCreatePersonCommandHandler(personRepository memory.Repository) *CreatePersonCommandHandler {
	return &CreatePersonCommandHandler{
		personRepository: personRepository,
	}
}

func (c *CreatePersonCommandHandler) Name() string {
	return fmt.Sprintf("%T", &CreatePersonCommand{})
}

func (c *CreatePersonCommandHandler) Handle(ctx context.Context, request any) (any, error) {
	cmd, ok := request.(*CreatePersonCommand)
	if !ok {
		return nil, errors.New("invalid command")
	}

	id, err := c.personRepository.Save(memory.CreatePersonRequest{
		Email:     cmd.Email,
		FirstName: cmd.FirstName,
		LastName:  cmd.LastName,
	})
	if err != nil {
		return nil, errors.New("person creation failed")
	}

	return dto.CreatePersonResponse{ID: id}, nil
}
