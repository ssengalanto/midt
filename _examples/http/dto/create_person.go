package dto

type CreatePersonRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type CreatePersonResponse struct {
	ID string `json:"id"`
}
