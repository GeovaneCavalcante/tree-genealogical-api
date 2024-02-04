package presenter

import (
	"github.com/GeovaneCavalcante/tree-genealogical/relationship"
	"github.com/go-playground/validator/v10"
)

type PaternityRelationship struct {
	ID     string `json:"id"`
	Parent string `json:"parent" validate:"required"`
	Child  string `json:"child" validate:"required"`
}

func NewPaternityRelationship(relationship *relationship.Relationship) *PaternityRelationship {
	return &PaternityRelationship{
		ID:     relationship.ID,
		Parent: relationship.SecundePersonID,
		Child:  relationship.MainPersonID,
	}
}

func NewPaternityRelationships(relationships []*relationship.Relationship) []*PaternityRelationship {
	var response []*PaternityRelationship
	for _, r := range relationships {
		response = append(response, NewPaternityRelationship(r))
	}
	return response
}

func (p *PaternityRelationship) ToRelationship() *relationship.Relationship {
	return &relationship.Relationship{
		ID:              p.ID,
		MainPersonID:    p.Child,
		SecundePersonID: p.Parent,
	}
}

func (p *PaternityRelationship) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(p)
}
