package handlers

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/ssengalanto/midt"
	"github.com/ssengalanto/midt/_examples/http/constants"
	"github.com/ssengalanto/midt/_examples/http/dto"
	"github.com/ssengalanto/midt/_examples/http/http/response"
	"github.com/ssengalanto/midt/_examples/http/query"
)

type GetPersonHandler struct {
	m *midt.Midt
}

func NewGetPersonHandler(m *midt.Midt) *GetPersonHandler {
	return &GetPersonHandler{m: m}
}

func (g *GetPersonHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.RequestTimeout)
	defer cancel()

	if r.Method != http.MethodGet {
		log.Println("invalid http method")
		response.Error(w, http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/person/")

	req := dto.GetPersonRequest{ID: id}

	q := query.NewGetPersonQuery(req.ID)
	res, err := g.m.Send(ctx, q)
	if err != nil {
		response.Error(w, http.StatusInternalServerError)
		return
	}

	response.Success(w, http.StatusCreated, res)
}
