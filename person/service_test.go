package person

import (
	"context"
	"errors"
	"testing"

	"github.com/GeovaneCavalcante/tree-genealogical/internal/entity"
	mock_person "github.com/GeovaneCavalcante/tree-genealogical/person/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type PersonServiceTestSuite struct {
	suite.Suite
	PersonRepoMock *mock_person.MockRepository
	Person         *entity.Person
}

func (suite *PersonServiceTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	suite.PersonRepoMock = mock_person.NewMockRepository(ctrl)
	suite.Person = &entity.Person{
		ID:     "1",
		Name:   "John",
		Gender: "M",
	}
}

func (suite *PersonServiceTestSuite) TestCreate() {
	ctx := context.Background()
	suite.Run("should return success when creating a person", func() {
		suite.PersonRepoMock.EXPECT().Create(gomock.Any(), suite.Person).Return(nil)
		service := NewService(suite.PersonRepoMock)
		err := service.Create(ctx, suite.Person)
		assert.Nil(suite.T(), err)
	})

	suite.Run("should return error when creating a person", func() {
		suite.PersonRepoMock.EXPECT().Create(gomock.Any(), suite.Person).Return(errors.New("database error"))
		service := NewService(suite.PersonRepoMock)
		err := service.Create(ctx, suite.Person)
		assert.NotNil(suite.T(), err)
		assert.Equal(suite.T(), "create person error: database error", err.Error())
	})
}

func (suite *PersonServiceTestSuite) TestGet() {
	ctx := context.Background()
	suite.Run("should return success when getting a person", func() {
		suite.PersonRepoMock.EXPECT().Get(gomock.Any(), suite.Person.ID).Return(suite.Person, nil)
		service := NewService(suite.PersonRepoMock)
		person, err := service.Get(ctx, suite.Person.ID)
		assert.Nil(suite.T(), err)
		assert.Equal(suite.T(), suite.Person, person)
	})

	suite.Run("should return error when getting a person", func() {
		suite.PersonRepoMock.EXPECT().Get(gomock.Any(), suite.Person.ID).Return(nil, errors.New("database error"))
		service := NewService(suite.PersonRepoMock)
		person, err := service.Get(ctx, suite.Person.ID)
		assert.NotNil(suite.T(), err)
		assert.Nil(suite.T(), person)
		assert.Equal(suite.T(), "get person error: database error", err.Error())
	})
}

func (suite *PersonServiceTestSuite) TestList() {
	ctx := context.Background()
	suite.Run("should return success when listing persons", func() {
		suite.PersonRepoMock.EXPECT().List(gomock.Any(), nil).Return([]*entity.Person{suite.Person}, nil)
		service := NewService(suite.PersonRepoMock)
		persons, err := service.List(ctx, nil)
		assert.Nil(suite.T(), err)
		assert.Equal(suite.T(), []*entity.Person{suite.Person}, persons)
	})

	suite.Run("should return error when listing persons", func() {
		suite.PersonRepoMock.EXPECT().List(gomock.Any(), nil).Return(nil, errors.New("database error"))
		service := NewService(suite.PersonRepoMock)
		persons, err := service.List(ctx, nil)
		assert.NotNil(suite.T(), err)
		assert.Nil(suite.T(), persons)
		assert.Equal(suite.T(), "list person error: database error", err.Error())
	})

	suite.Run("should return not found when listing persons", func() {
		suite.PersonRepoMock.EXPECT().List(gomock.Any(), nil).Return(nil, nil)
		service := NewService(suite.PersonRepoMock)
		persons, err := service.List(ctx, nil)
		assert.Nil(suite.T(), err)
		assert.Nil(suite.T(), persons)
	})
}

func (suite *PersonServiceTestSuite) TestUpdate() {
	ctx := context.Background()
	suite.Run("should return success when updating a person", func() {
		suite.PersonRepoMock.EXPECT().Get(gomock.Any(), suite.Person.ID).Return(suite.Person, nil)
		suite.PersonRepoMock.EXPECT().Update(gomock.Any(), suite.Person.ID, suite.Person).Return(nil)
		service := NewService(suite.PersonRepoMock)
		err := service.Update(ctx, suite.Person.ID, suite.Person)
		assert.Nil(suite.T(), err)
	})

	suite.Run("should return error when updating a person", func() {
		suite.PersonRepoMock.EXPECT().Get(gomock.Any(), suite.Person.ID).Return(suite.Person, nil)
		suite.PersonRepoMock.EXPECT().Update(gomock.Any(), suite.Person.ID, suite.Person).Return(errors.New("database error"))
		service := NewService(suite.PersonRepoMock)
		err := service.Update(ctx, suite.Person.ID, suite.Person)
		assert.NotNil(suite.T(), err)
		assert.Equal(suite.T(), "update person error: database error", err.Error())
	})

	suite.Run("should return error when getting a person to update", func() {
		suite.PersonRepoMock.EXPECT().Get(gomock.Any(), suite.Person.ID).Return(nil, errors.New("database error"))
		service := NewService(suite.PersonRepoMock)
		err := service.Update(ctx, suite.Person.ID, suite.Person)
		assert.NotNil(suite.T(), err)
		assert.Equal(suite.T(), "update person error: get person error: database error", err.Error())
	})

	suite.Run("should return not found when getting a person to update", func() {
		suite.PersonRepoMock.EXPECT().Get(gomock.Any(), suite.Person.ID).Return(nil, nil)
		service := NewService(suite.PersonRepoMock)
		err := service.Update(ctx, suite.Person.ID, suite.Person)
		assert.NotNil(suite.T(), err)
		assert.Equal(suite.T(), "update person error: not found", err.Error())
	})
}

func (suite *PersonServiceTestSuite) TestDelete() {
	ctx := context.Background()
	suite.Run("should return success when deleting a person", func() {
		suite.PersonRepoMock.EXPECT().Get(gomock.Any(), suite.Person.ID).Return(suite.Person, nil)
		suite.PersonRepoMock.EXPECT().Delete(gomock.Any(), suite.Person.ID).Return(nil)
		service := NewService(suite.PersonRepoMock)
		err := service.Delete(ctx, suite.Person.ID)
		assert.Nil(suite.T(), err)
	})

	suite.Run("should return error when deleting a person", func() {
		suite.PersonRepoMock.EXPECT().Get(gomock.Any(), suite.Person.ID).Return(suite.Person, nil)
		suite.PersonRepoMock.EXPECT().Delete(gomock.Any(), suite.Person.ID).Return(errors.New("database error"))
		service := NewService(suite.PersonRepoMock)
		err := service.Delete(ctx, suite.Person.ID)
		assert.NotNil(suite.T(), err)
		assert.Equal(suite.T(), "delete person error: database error", err.Error())
	})

	suite.Run("should return error when getting a person to delete", func() {
		suite.PersonRepoMock.EXPECT().Get(gomock.Any(), suite.Person.ID).Return(nil, errors.New("database error"))
		service := NewService(suite.PersonRepoMock)
		err := service.Delete(ctx, suite.Person.ID)
		assert.NotNil(suite.T(), err)
		assert.Equal(suite.T(), "delete person error: get person error: database error", err.Error())
	})

	suite.Run("should return not found when getting a person to delete", func() {
		suite.PersonRepoMock.EXPECT().Get(gomock.Any(), suite.Person.ID).Return(nil, nil)
		service := NewService(suite.PersonRepoMock)
		err := service.Delete(ctx, suite.Person.ID)
		assert.NotNil(suite.T(), err)
		assert.Equal(suite.T(), "delete person error: not found", err.Error())
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(PersonServiceTestSuite))
}
