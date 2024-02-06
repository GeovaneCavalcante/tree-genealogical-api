package relationship

import (
	"context"
	"errors"
	"testing"

	"github.com/GeovaneCavalcante/tree-genealogical/internal/entity"
	mock_relationship "github.com/GeovaneCavalcante/tree-genealogical/relationship/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type RelationshipServiceTestSuite struct {
	suite.Suite
	RelationshipRepoMock *mock_relationship.MockRepository
	Relationship         *entity.Relationship
}

func (suite *RelationshipServiceTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	suite.RelationshipRepoMock = mock_relationship.NewMockRepository(ctrl)
	suite.Relationship = &entity.Relationship{
		ID:              "1",
		SecundePersonID: "2",
		MainPersonID:    "3",
	}
}

func (suite *RelationshipServiceTestSuite) TestCreate() {
	ctx := context.Background()
	suite.Run("should return success when creating a relationship", func() {
		suite.RelationshipRepoMock.EXPECT().Create(gomock.Any(), suite.Relationship).Return(nil)
		service := NewService(suite.RelationshipRepoMock)
		err := service.Create(ctx, suite.Relationship)
		suite.Nil(err)
	})

	suite.Run("should return error when creating a relationship", func() {
		suite.RelationshipRepoMock.EXPECT().Create(gomock.Any(), suite.Relationship).Return(errors.New("database error"))
		service := NewService(suite.RelationshipRepoMock)
		err := service.Create(ctx, suite.Relationship)
		suite.NotNil(err)
		suite.Equal("create relationship error: database error", err.Error())
	})
}

func (suite *RelationshipServiceTestSuite) TestGet() {
	ctx := context.Background()
	suite.Run("should return success when getting a relationship", func() {
		suite.RelationshipRepoMock.EXPECT().Get(gomock.Any(), suite.Relationship.ID).Return(suite.Relationship, nil)
		service := NewService(suite.RelationshipRepoMock)
		relationship, err := service.Get(ctx, suite.Relationship.ID)
		suite.Nil(err)
		suite.Equal(suite.Relationship, relationship)
	})

	suite.Run("should return error when getting a relationship", func() {
		suite.RelationshipRepoMock.EXPECT().Get(gomock.Any(), suite.Relationship.ID).Return(nil, errors.New("database error"))
		service := NewService(suite.RelationshipRepoMock)
		relationship, err := service.Get(ctx, suite.Relationship.ID)
		suite.NotNil(err)
		suite.Nil(relationship)
		suite.Equal("get relationship error: database error", err.Error())
	})
}

func (suite *RelationshipServiceTestSuite) TestList() {
	ctx := context.Background()
	suite.Run("should return success when listing relationships", func() {
		suite.RelationshipRepoMock.EXPECT().List(gomock.Any(), gomock.Any()).Return([]*entity.Relationship{suite.Relationship}, nil)
		service := NewService(suite.RelationshipRepoMock)
		relationships, err := service.List(ctx, map[string]interface{}{})
		suite.Nil(err)
		suite.NotNil(relationships)
		suite.Equal([]*entity.Relationship{suite.Relationship}, relationships)
	})

	suite.Run("should return error when listing relationships", func() {
		suite.RelationshipRepoMock.EXPECT().List(gomock.Any(), gomock.Any()).Return(nil, errors.New("database error"))
		service := NewService(suite.RelationshipRepoMock)
		relationships, err := service.List(ctx, map[string]interface{}{})
		suite.NotNil(err)
		suite.Nil(relationships)
		suite.Equal("list relationship error: database error", err.Error())
	})
}

func (suite *RelationshipServiceTestSuite) TestUpdate() {
	ctx := context.Background()
	suite.Run("should return success when updating a relationship", func() {
		suite.RelationshipRepoMock.EXPECT().Get(gomock.Any(), suite.Relationship.ID).Return(suite.Relationship, nil)
		suite.RelationshipRepoMock.EXPECT().Update(gomock.Any(), suite.Relationship.ID, suite.Relationship).Return(nil)
		service := NewService(suite.RelationshipRepoMock)
		err := service.Update(ctx, suite.Relationship.ID, suite.Relationship)
		suite.Nil(err)
	})

	suite.Run("should return error when updating a relationship", func() {
		suite.RelationshipRepoMock.EXPECT().Get(gomock.Any(), suite.Relationship.ID).Return(suite.Relationship, nil)
		suite.RelationshipRepoMock.EXPECT().Update(gomock.Any(), suite.Relationship.ID, suite.Relationship).Return(errors.New("database error"))
		service := NewService(suite.RelationshipRepoMock)
		err := service.Update(ctx, suite.Relationship.ID, suite.Relationship)
		suite.NotNil(err)
		suite.Equal("update relationship error: database error", err.Error())
	})

	suite.Run("should return error when getting a relationship", func() {
		suite.RelationshipRepoMock.EXPECT().Get(gomock.Any(), suite.Relationship.ID).Return(nil, errors.New("database error"))
		service := NewService(suite.RelationshipRepoMock)
		err := service.Update(ctx, suite.Relationship.ID, suite.Relationship)
		suite.NotNil(err)
		suite.Equal("update relationship error: get relationship error: database error", err.Error())
	})

	suite.Run("should return error when getting a relationship nil", func() {
		suite.RelationshipRepoMock.EXPECT().Get(gomock.Any(), suite.Relationship.ID).Return(nil, nil)
		service := NewService(suite.RelationshipRepoMock)
		err := service.Update(ctx, suite.Relationship.ID, suite.Relationship)
		suite.NotNil(err)
		suite.Equal("relationship not found", err.Error())
	})
}

func (suite *RelationshipServiceTestSuite) TestDelete() {
	ctx := context.Background()
	suite.Run("should return success when deleting a relationship", func() {
		suite.RelationshipRepoMock.EXPECT().Get(gomock.Any(), suite.Relationship.ID).Return(suite.Relationship, nil)
		suite.RelationshipRepoMock.EXPECT().Delete(gomock.Any(), suite.Relationship.ID).Return(nil)
		service := NewService(suite.RelationshipRepoMock)
		err := service.Delete(ctx, suite.Relationship.ID)
		suite.Nil(err)
	})

	suite.Run("should return error when deleting a relationship", func() {
		suite.RelationshipRepoMock.EXPECT().Get(gomock.Any(), suite.Relationship.ID).Return(suite.Relationship, nil)
		suite.RelationshipRepoMock.EXPECT().Delete(gomock.Any(), suite.Relationship.ID).Return(errors.New("database error"))
		service := NewService(suite.RelationshipRepoMock)
		err := service.Delete(ctx, suite.Relationship.ID)
		suite.NotNil(err)
		suite.Equal("delete relationship error: database error", err.Error())
	})

	suite.Run("should return error when getting a relationship", func() {
		suite.RelationshipRepoMock.EXPECT().Get(gomock.Any(), suite.Relationship.ID).Return(nil, errors.New("database error"))
		service := NewService(suite.RelationshipRepoMock)
		err := service.Delete(ctx, suite.Relationship.ID)
		suite.NotNil(err)
		suite.Equal("delete relationship error: get relationship error: database error", err.Error())
	})

	suite.Run("should return error when getting a relationship nil", func() {
		suite.RelationshipRepoMock.EXPECT().Get(gomock.Any(), suite.Relationship.ID).Return(nil, nil)
		service := NewService(suite.RelationshipRepoMock)
		err := service.Delete(ctx, suite.Relationship.ID)
		suite.NotNil(err)
		suite.Equal("relationship not found", err.Error())
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(RelationshipServiceTestSuite))
}
