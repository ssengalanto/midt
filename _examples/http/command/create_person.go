package command

type CreatePersonCommand struct {
	Email     string
	FirstName string
	LastName  string
}

func NewCreatePersonCommand(email, fn, ln string) *CreatePersonCommand {
	return &CreatePersonCommand{
		Email:     email,
		FirstName: fn,
		LastName:  ln,
	}
}
