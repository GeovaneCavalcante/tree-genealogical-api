package presenter

import (
	"github.com/GeovaneCavalcante/tree-genealogical/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PersonPresenerTestSuite struct {
	suite.Suite
	Person *entity.Person
}

func (suite *PersonPresenerTestSuite) SetupTest() {
	suite.Person = &entity.Person{
		ID:     "id",
		Name:   "Ruff",
		Gender: "M",
	}
}

func (suite *PersonPresenerTestSuite) TestNewPersonResponse() {
	suite.Run("When person is not empty", func() {
		response := NewPersonResponse(suite.Person)
		assert.Equal(suite.T(), suite.Person.ID, response.ID)
		assert.Equal(suite.T(), suite.Person.Name, response.Name)
	})
}

func (suite *PersonPresenerTestSuite) TestNewPersonsResponse() {
	suite.Run("When persons is not empty", func() {
		persons := []*entity.Person{
			{
				ID:     "id",
				Name:   "Ruff",
				Gender: "M",
			},
		}
		response := NewPersonsResponse(persons)
		assert.Equal(suite.T(), persons[0].ID, response[0].ID)
		assert.Equal(suite.T(), persons[0].Name, response[0].Name)
	})
}

func (suite *PersonPresenerTestSuite) TestNewPersonRequest() {
	suite.Run("When person is not empty", func() {
		response := NewPersonRequest(suite.Person)
		assert.Equal(suite.T(), suite.Person.Name, response.Name)
		assert.Equal(suite.T(), suite.Person.Gender, response.Gender)
	})
}

func (suite *PersonPresenerTestSuite) TestToPerson() {
	suite.Run("When person is not empty", func() {
		request := NewPersonRequest(suite.Person)
		response := request.ToPerson()
		assert.Equal(suite.T(), suite.Person.Name, response.Name)
		assert.Equal(suite.T(), suite.Person.Gender, response.Gender)
	})
}

func (suite *PersonPresenerTestSuite) TestValidate() {
	suite.Run("When person is not empty", func() {
		request := NewPersonRequest(suite.Person)
		err := request.Validate()
		assert.Nil(suite.T(), err)
	})
}
