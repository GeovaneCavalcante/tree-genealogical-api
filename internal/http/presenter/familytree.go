package presenter

import "github.com/GeovaneCavalcante/tree-genealogical/internal/entity"

type FamilyTreeResponse struct {
	Members []*Member `json:"members" xml:"members"`
}

type DetermineRelationResponse struct {
	Relationship string `json:"relationship" xml:"relationship"`
}

type KinshipDistanceResponse struct {
	Distance int `json:"distance" xml:"distance"`
}

type Member struct {
	Name             string          `json:"name" xml:"name"`
	TypeRelationship string          `json:"typeRelationship" xml:"typeRelationship"`
	Relationships    []*Relationship `json:"relationships" xml:"relationships"`
}

type Relationship struct {
	Name string `json:"parent" xml:"parent"`
}

func NewKinshipDistanceResponse(distance int) *KinshipDistanceResponse {
	return &KinshipDistanceResponse{
		Distance: distance,
	}
}

func NewDetermineRelationResponse(relationship string) *DetermineRelationResponse {
	return &DetermineRelationResponse{
		Relationship: relationship,
	}
}

func NewFamilyTreeResponse(relatives []*entity.Relative) *FamilyTreeResponse {
	members := make([]*Member, 0, len(relatives))

	for _, relative := range relatives {
		if relative.Person == nil {
			continue
		}

		member := &Member{
			Name:             relative.Person.Name,
			TypeRelationship: relative.Type,
			Relationships:    make([]*Relationship, 0, len(relative.Person.Relationships)),
		}

		for _, rel := range relative.Person.Relationships {
			if rel.SecundePerson == nil {
				continue
			}
			member.Relationships = append(member.Relationships, &Relationship{
				Name: rel.SecundePerson.Name,
			})
		}

		members = append(members, member)
	}

	return &FamilyTreeResponse{
		Members: members,
	}
}
