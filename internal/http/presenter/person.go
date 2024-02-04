package presenter

import (
	"github.com/GeovaneCavalcante/tree-genealogical/person"
	"github.com/go-playground/validator/v10"
)

type PersonResponse struct {
	ID     string `json:"id" xml:"id"`
	Name   string `json:"name" xml:"name"`
	Gender string `json:"gender" xml:"gender"`
}

type PersonRequest struct {
	Name   string `json:"name" xml:"name" validate:"required"`
	Gender string `json:"gender" xml:"gender" validate:"required,oneof=F M"`
}

func NewPersonResponse(person *person.Person) *PersonResponse {
	return &PersonResponse{
		ID:     person.ID,
		Name:   person.Name,
		Gender: person.Gender,
	}
}

func NewPersonsResponse(persons []*person.Person) []*PersonResponse {
	var response []*PersonResponse
	for _, p := range persons {
		response = append(response, NewPersonResponse(p))
	}
	return response
}

func NewPersonRequest(person *person.Person) *PersonRequest {
	return &PersonRequest{
		Name:   person.Name,
		Gender: person.Gender,
	}
}

func (p *PersonRequest) ToPerson() *person.Person {
	return &person.Person{
		Name:   p.Name,
		Gender: p.Gender,
	}
}

func (p *PersonRequest) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(p)
}
