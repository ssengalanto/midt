package memory

type Person struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
}

type CreatePersonRequest struct {
	Email     string
	FirstName string
	LastName  string
}

type UpdatePersonRequest struct {
	Email     string
	FirstName string
	LastName  string
}
