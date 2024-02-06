package presenter

import (
	"github.com/GeovaneCavalcante/tree-genealogical/internal/entity"
	"github.com/stretchr/testify/suite"
)

type RelationshipPresenerTestSuite struct {
	suite.Suite
	Relationship *entity.Relationship
	Pr           *PaternityRelationshipRequest
}

func (suite *RelationshipPresenerTestSuite) SetupTest() {
	suite.Relationship = &entity.Relationship{
		MainPersonID:    "123",
		SecundePersonID: "456",
		MainPerson: &entity.Person{
			Name:   "Ruff",
			Gender: "M",
		},
		SecundePerson: &entity.Person{
			Name:   "Mary",
			Gender: "F",
		},
	}

	suite.Pr = &PaternityRelationshipRequest{
		Parent: suite.Relationship.SecundePersonID,
		Child:  suite.Relationship.MainPersonID,
	}
}

func (suite *RelationshipPresenerTestSuite) TestNewRelationshipResponse() {
	suite.Run("When relationship is not empty", func() {
		response := NewPaternityRelationshipResponse(suite.Relationship)
		suite.Equal(suite.Relationship.MainPersonID, response.Child)
		suite.Equal(suite.Relationship.SecundePersonID, response.Parent)
	})
}

func (suite *RelationshipPresenerTestSuite) TestNewRelationshipsResponse() {
	suite.Run("When relationships is not empty", func() {
		response := NewPaternityRelationshipsResponse([]*entity.Relationship{suite.Relationship})
		suite.Equal(suite.Relationship.MainPersonID, response[0].Child)
		suite.Equal(suite.Relationship.SecundePersonID, response[0].Parent)
	})
}

func (suite *RelationshipPresenerTestSuite) TestNewRelationshipRequest() {
	suite.Run("When relationship is not empty", func() {
		response := suite.Pr.NewPaternityRelationshipRequest()
		suite.Equal(suite.Relationship.MainPersonID, response.MainPersonID)
		suite.Equal(suite.Relationship.SecundePersonID, response.SecundePersonID)
	})
}

func (suite *RelationshipPresenerTestSuite) TestToRelationship() {
	suite.Run("When relationship is not empty", func() {
		response := suite.Pr.ToRelationship()
		suite.Equal(suite.Relationship.MainPersonID, response.MainPersonID)
		suite.Equal(suite.Relationship.SecundePersonID, response.SecundePersonID)
	})
}

func (suite *RelationshipPresenerTestSuite) TestValidate() {
	suite.Run("When relationship is not empty", func() {
		err := suite.Pr.Validate()
		suite.Nil(err)
	})
}
