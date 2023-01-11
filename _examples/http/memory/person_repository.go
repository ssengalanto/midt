package memory

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
)

type Repository interface {
	Save(request CreatePersonRequest) (string, error)
	FindByID(id string) (Person, error)
	UpdateByID(id string, request UpdatePersonRequest) (string, error)
	DeleteById(id string) (string, error)
}

type PersonRepository struct {
	person map[string]Person
}

func NewPersonRepository() *PersonRepository {
	return &PersonRepository{
		person: map[string]Person{},
	}
}

func (p *PersonRepository) Save(request CreatePersonRequest) (string, error) {
	id := generateID()

	if _, ok := p.person[id]; ok {
		return "", errors.New("person already exists")
	}

	p.person[id] = Person{
		ID:        id,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
	}
	return id, nil
}

func (p *PersonRepository) FindByID(id string) (Person, error) {
	empty := Person{}

	person, ok := p.person[id]
	if !ok {
		return empty, errors.New("person not found")
	}

	return person, nil
}

func (p *PersonRepository) UpdateByID(id string, request UpdatePersonRequest) (string, error) {
	_, ok := p.person[id]
	if !ok {
		return "", errors.New("person not found")
	}

	p.person[id] = Person{
		ID:        id,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
	}
	return id, nil
}

func (p *PersonRepository) DeleteById(id string) (string, error) {
	_, ok := p.person[id]
	if !ok {
		return "", errors.New("person not found")
	}

	delete(p.person, id)
	return id, nil
}

func generateID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal("generating unique id failed")
	}

	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
