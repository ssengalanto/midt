package query

import (
	"context"
	"errors"
	"fmt"

	"github.com/ssengalanto/midt/_examples/http/dto"
	"github.com/ssengalanto/midt/_examples/http/memory"
)

type GetPersonQueryHandler struct {
	personRepository memory.Repository
}

func NewGetPersonQueryHandler(personRepository memory.Repository) *GetPersonQueryHandler {
	return &GetPersonQueryHandler{
		personRepository: personRepository,
	}
}

func (g *GetPersonQueryHandler) Name() string {
	return fmt.Sprintf("%T", &GetPersonQuery{})
}

func (g *GetPersonQueryHandler) Handle(ctx context.Context, request any) (any, error) {
	empty := dto.GetPersonResponse{}

	q, ok := request.(*GetPersonQuery)
	if !ok {
		return empty, errors.New("invalid command")
	}

	pers, err := g.personRepository.FindByID(q.ID)
	if err != nil {
		return empty, errors.New("pers creation failed")
	}

	return dto.GetPersonResponse{
		ID:        pers.ID,
		Email:     pers.Email,
		FirstName: pers.FirstName,
		LastName:  pers.LastName,
	}, nil
}
