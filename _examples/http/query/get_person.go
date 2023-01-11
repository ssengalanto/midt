package query

type GetPersonQuery struct {
	ID string
}

func NewGetPersonQuery(id string) *GetPersonQuery {
	return &GetPersonQuery{
		ID: id,
	}
}
