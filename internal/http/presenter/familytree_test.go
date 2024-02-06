package presenter

import (
	"github.com/GeovaneCavalcante/tree-genealogical/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FamilyTreePresenerTestSuite struct {
	suite.Suite
}

func (suite *FamilyTreePresenerTestSuite) SetupTest() {
}

func (suite *FamilyTreePresenerTestSuite) TestNewKinshipDistanceResponse() {
	suite.Run("When distance is not empty", func() {
		distance := 1
		response := NewKinshipDistanceResponse(distance)
		assert.Equal(suite.T(), distance, response.Distance)
	})
}

func (suite *FamilyTreePresenerTestSuite) TestNewDetermineRelationResponse() {
	suite.Run("When relationship is not empty", func() {
		relationship := "relationship"
		response := NewDetermineRelationResponse(relationship)
		assert.Equal(suite.T(), relationship, response.Relationship)
	})
}

func (suite *FamilyTreePresenerTestSuite) TestNewFamilyTreeResponse() {

	suite.Run("When relatives is not empty", func() {
		relatives := []*entity.Relative{
			{
				Person: &entity.Person{
					Name: "name",
					Relationships: []*entity.Relationship{
						{
							SecundePerson: &entity.Person{
								Name: "name",
							},
						},
					},
				},
				Type: "type",
			},

			{
				Person: &entity.Person{
					Name: "name",
					Relationships: []*entity.Relationship{
						{
							SecundePerson: nil,
						},
					},
				},
				Type: "type",
			},
		}
		response := NewFamilyTreeResponse(relatives)
		assert.Equal(suite.T(), "name", response.Members[0].Name)
		assert.Equal(suite.T(), "type", response.Members[0].TypeRelationship)
		assert.Equal(suite.T(), "name", response.Members[0].Relationships[0].Name)
	})

	suite.Run("When relatives is empty", func() {
		relatives := []*entity.Relative{
			{
				Person: nil,
				Type:   "type",
			},
		}
		response := NewFamilyTreeResponse(relatives)
		suite.Len(response.Members, 0)
	})
}
