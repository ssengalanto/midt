package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ssengalanto/midt"
	"github.com/ssengalanto/midt/_examples/http/behaviour"
	"github.com/ssengalanto/midt/_examples/http/command"
	"github.com/ssengalanto/midt/_examples/http/constants"
	"github.com/ssengalanto/midt/_examples/http/event"
	"github.com/ssengalanto/midt/_examples/http/http/handlers"
	"github.com/ssengalanto/midt/_examples/http/memory"
	"github.com/ssengalanto/midt/_examples/http/query"
)

func main() {
	m := midt.New()
	repo := memory.NewPersonRepository()
	mux := http.NewServeMux()
	registerHandlers(repo, m)
	registerHttpHandlers(mux, m)
	registerBehaviours(m)
	registerNotifications(m)

	svr := &http.Server{
		Addr:         ":" + strconv.Itoa(constants.Port),
		Handler:      mux,
		IdleTimeout:  constants.IdleTimeout,
		ReadTimeout:  constants.ReadTimeout,
		WriteTimeout: constants.WriteTimeout,
	}
	err := svr.ListenAndServe()
	if err != nil {
		log.Fatalln("server failed to start")
	}
}

func registerHttpHandlers(mux *http.ServeMux, m *midt.Midt) {
	createPersonHandler := handlers.NewCreatePersonHandler(m)
	mux.Handle("/person", http.HandlerFunc(createPersonHandler.Handle))

	getPersonHandler := handlers.NewGetPersonHandler(m)
	mux.Handle("/person/", http.HandlerFunc(getPersonHandler.Handle))
}

func registerNotifications(m *midt.Midt) {
	personCreatedNotificationHandler := event.NewPersonCreatedEventHandler()
	err := m.RegisterNotificationHandler(personCreatedNotificationHandler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerBehaviours(m *midt.Midt) {
	loggerBehaviour := behaviour.NewLoggerBehaviour()
	err := m.RegisterPipelineBehaviour(loggerBehaviour)
	if err != nil {
		log.Fatal(err)
	}
}

func registerHandlers(repo memory.Repository, m *midt.Midt) {
	createPersonCommandHandler := command.NewCreatePersonCommandHandler(repo)
	err := m.RegisterRequestHandler(createPersonCommandHandler)
	if err != nil {
		log.Fatal(err)
	}

	getPersonQueryHandler := query.NewGetPersonQueryHandler(repo)
	err = m.RegisterRequestHandler(getPersonQueryHandler)
	if err != nil {
		log.Fatal(err)
	}
}
