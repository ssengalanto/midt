package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ssengalanto/midt"
	"github.com/ssengalanto/midt/_examples/http/command"
	"github.com/ssengalanto/midt/_examples/http/constants"
	"github.com/ssengalanto/midt/_examples/http/dto"
	"github.com/ssengalanto/midt/_examples/http/event"
	"github.com/ssengalanto/midt/_examples/http/http/response"
	"github.com/ssengalanto/midt/_examples/http/query"
)

type CreatePersonHandler struct {
	m *midt.Midt
}

func NewCreatePersonHandler(m *midt.Midt) *CreatePersonHandler {
	return &CreatePersonHandler{m: m}
}

func (c *CreatePersonHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.RequestTimeout)
	defer cancel()

	if r.Method != http.MethodPost {
		log.Println("invalid http method")
		response.Error(w, http.StatusMethodNotAllowed)
		return
	}

	var req dto.CreatePersonRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("invalid request body")
		response.Error(w, http.StatusBadRequest)
		return
	}

	// Executing CreatePersonCommand
	cmd := command.NewCreatePersonCommand(req.Email, req.FirstName, req.LastName)
	cmdres, err := c.m.Send(ctx, cmd)
	if err != nil {
		log.Printf("%T execution failed: %+v\n", cmd, err)
		response.Error(w, http.StatusInternalServerError)
		return
	}

	// Executing GetPersonQuery
	id := cmdres.(dto.CreatePersonResponse).ID
	q := query.NewGetPersonQuery(id)
	qres, err := c.m.Send(ctx, q)
	if err != nil {
		log.Printf("%T execution failed: %+v\n", q, err)
		response.Error(w, http.StatusInternalServerError)
		return
	}

	// Executing PersonCreatedEvent
	pers := qres.(dto.GetPersonResponse)
	evt := event.NewPersonCreatedEvent(pers.ID, pers.Email, pers.FirstName, pers.LastName)
	err = c.m.Publish(ctx, evt)
	if err != nil {
		log.Printf("%T execution failed: %+v\n", evt, err)
	}

	response.Success(w, http.StatusCreated, pers)
}
