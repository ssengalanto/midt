package event

type PersonCreatedEvent struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
}

func NewPersonCreatedEvent(id, email, fn, ln string) *PersonCreatedEvent {
	return &PersonCreatedEvent{
		ID:        id,
		Email:     email,
		FirstName: fn,
		LastName:  ln,
	}
}
