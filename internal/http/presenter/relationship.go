package presenter

import (
	"github.com/GeovaneCavalcante/tree-genealogical/relationship"
	"github.com/go-playground/validator/v10"
)

type PaternityRelationshipResponse struct {
	ID     string `json:"id" xml:"id"`
	Parent string `json:"parent" xml:"parent"`
	Child  string `json:"child" xml:"child"`
}

type PaternityRelationshipRequest struct {
	Parent string `json:"parent" xml:"parent" validate:"required"`
	Child  string `json:"child" xml:"child" validate:"required"`
}

func NewPaternityRelationshipResponse(relationship *relationship.Relationship) *PaternityRelationshipResponse {
	return &PaternityRelationshipResponse{
		ID:     relationship.ID,
		Parent: relationship.SecundePersonID,
		Child:  relationship.MainPersonID,
	}
}

func NewPaternityRelationshipsResponse(relationships []*relationship.Relationship) []*PaternityRelationshipResponse {
	var response []*PaternityRelationshipResponse
	for _, r := range relationships {
		response = append(response, NewPaternityRelationshipResponse(r))
	}
	return response
}

func (p *PaternityRelationshipRequest) NewPaternityRelationshipRequest() *relationship.Relationship {
	return &relationship.Relationship{
		MainPersonID:    p.Child,
		SecundePersonID: p.Parent,
	}
}

func (p *PaternityRelationshipRequest) ToRelationship() *relationship.Relationship {
	return &relationship.Relationship{
		MainPersonID:    p.Child,
		SecundePersonID: p.Parent,
	}
}

func (p *PaternityRelationshipRequest) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(p)
}
