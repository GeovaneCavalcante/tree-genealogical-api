package presenter

import (
	"github.com/GeovaneCavalcante/tree-genealogical/person"
	"github.com/go-playground/validator/v10"
)

type Person struct {
	ID     string `json:"id"`
	Name   string `json:"name" validate:"required"`
	Gender string `json:"gender" validate:"required,oneof=F M"`
}

func NewPerson(person *person.Person) *Person {
	return &Person{
		ID:     person.ID,
		Name:   person.Name,
		Gender: person.Gender,
	}
}

func NewPersons(persons []*person.Person) []*Person {
	var response []*Person
	for _, p := range persons {
		response = append(response, NewPerson(p))
	}
	return response
}

func (p *Person) ToPerson() *person.Person {
	return &person.Person{
		ID:     p.ID,
		Name:   p.Name,
		Gender: p.Gender,
	}
}

func (p *Person) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(p)
}
