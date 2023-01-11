package dto

type GetPersonRequest struct {
	ID string `json:"id"`
}

type GetPersonResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
