package familytree

import (
	"context"
	"errors"
	"testing"

	mock_genealogy "github.com/GeovaneCavalcante/tree-genealogical/familytree/mock"
	"github.com/GeovaneCavalcante/tree-genealogical/internal/entity"
	mock_person "github.com/GeovaneCavalcante/tree-genealogical/person/mock"
	mock_relationship "github.com/GeovaneCavalcante/tree-genealogical/relationship/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type FamilytreeTestSuite struct {
	suite.Suite
	GenealogyMock        *mock_genealogy.MockGenealogyInterface
	PersonRepoMock       *mock_person.MockRepository
	RelationshipRepoMock *mock_relationship.MockRepository
	PersonRoot           *entity.Person
	FamilyTree           []*entity.Relative
}

func (suite *FamilytreeTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	suite.GenealogyMock = mock_genealogy.NewMockGenealogyInterface(ctrl)
	suite.PersonRepoMock = mock_person.NewMockRepository(ctrl)
	suite.RelationshipRepoMock = mock_relationship.NewMockRepository(ctrl)
	suite.PersonRoot = &entity.Person{
		ID:     "1",
		Name:   "John",
		Gender: "M",
	}

	suite.FamilyTree = []*entity.Relative{
		{
			Type:   "Root",
			Level:  0,
			Person: suite.PersonRoot,
		},
		{
			Type:  "Father",
			Level: 1,
			Person: &entity.Person{
				ID:     "2",
				Name:   "Robert",
				Gender: "M",
			},
		},
		{
			Type:  "Mother",
			Level: 1,
			Person: &entity.Person{
				ID:     "3",
				Name:   "Maria",
				Gender: "F",
			},
		},
	}

}

func (suite *FamilytreeTestSuite) TestGetAllFamilyMembers() {
	ctx := context.Background()
	suite.Run("should return to the family tree successfully", func() {
		suite.PersonRepoMock.EXPECT().GetByName(gomock.Any(), gomock.Any()).Return(suite.PersonRoot, nil)
		suite.PersonRepoMock.EXPECT().ListWithRelationships(gomock.Any(), gomock.Any()).Return([]*entity.Person{}, nil)
		suite.GenealogyMock.EXPECT().BuildFamilyTree(gomock.Any(), gomock.Any(), gomock.Any(), 0).Return(suite.FamilyTree)

		service := NewService(suite.GenealogyMock, suite.PersonRepoMock, suite.RelationshipRepoMock)
		family, err := service.GetAllFamilyMembers(ctx, "John")
		assert.Nil(suite.T(), err)
		assert.Len(suite.T(), family, 3)
	})

	suite.Run("should return an error when trying to get the person", func() {
		suite.PersonRepoMock.EXPECT().GetByName(gomock.Any(), gomock.Any()).Return(nil, errors.New("error database"))
		service := NewService(suite.GenealogyMock, suite.PersonRepoMock, suite.RelationshipRepoMock)
		family, err := service.GetAllFamilyMembers(ctx, "John")
		assert.NotNil(suite.T(), err)
		assert.Error(suite.T(), err, "get person error: %w", "error database")
		assert.Nil(suite.T(), family)
	})

	suite.Run("should return an error when trying to get the list of people", func() {
		suite.PersonRepoMock.EXPECT().GetByName(gomock.Any(), gomock.Any()).Return(suite.PersonRoot, nil)
		suite.PersonRepoMock.EXPECT().ListWithRelationships(gomock.Any(), gomock.Any()).Return(nil, errors.New("error database"))
		service := NewService(suite.GenealogyMock, suite.PersonRepoMock, suite.RelationshipRepoMock)
		family, err := service.GetAllFamilyMembers(ctx, "John")
		assert.NotNil(suite.T(), err)
		assert.Error(suite.T(), err, "get person error: %w", "error database")
		assert.Nil(suite.T(), family)
	})
}

func (suite *FamilytreeTestSuite) TestDetermineRelationship() {
	ctx := context.Background()
	suite.Run("should return the relationship between two people successfully", func() {
		suite.PersonRepoMock.EXPECT().GetByName(gomock.Any(), gomock.Any()).Return(suite.PersonRoot, nil)
		suite.PersonRepoMock.EXPECT().ListWithRelationships(gomock.Any(), gomock.Any()).Return([]*entity.Person{}, nil)
		suite.GenealogyMock.EXPECT().BuildFamilyTree(gomock.Any(), gomock.Any(), gomock.Any(), 1).Return(suite.FamilyTree)

		service := NewService(suite.GenealogyMock, suite.PersonRepoMock, suite.RelationshipRepoMock)
		relationship, err := service.DetermineRelationship(ctx, "John", "Robert")
		assert.Nil(suite.T(), err)
		assert.Equal(suite.T(), "Father", relationship)
	})

	suite.Run("should return an error when trying to get the person", func() {
		suite.PersonRepoMock.EXPECT().GetByName(gomock.Any(), gomock.Any()).Return(nil, errors.New("error database"))
		service := NewService(suite.GenealogyMock, suite.PersonRepoMock, suite.RelationshipRepoMock)
		relationship, err := service.DetermineRelationship(ctx, "John", "Robert")
		assert.NotNil(suite.T(), err)
		assert.Error(suite.T(), err, "get person error: %w", "error database")
		assert.Empty(suite.T(), relationship)
	})

	suite.Run("should return an error when trying to get the list of people", func() {
		suite.PersonRepoMock.EXPECT().GetByName(gomock.Any(), gomock.Any()).Return(suite.PersonRoot, nil)
		suite.PersonRepoMock.EXPECT().ListWithRelationships(gomock.Any(), gomock.Any()).Return(nil, errors.New("error database"))
		service := NewService(suite.GenealogyMock, suite.PersonRepoMock, suite.RelationshipRepoMock)
		relationship, err := service.DetermineRelationship(ctx, "John", "Robert")
		assert.NotNil(suite.T(), err)
		assert.Error(suite.T(), err, "get person error: %w", "error database")
		assert.Empty(suite.T(), relationship)
	})

	suite.Run("should return unrelated when the people are not related", func() {
		suite.PersonRepoMock.EXPECT().GetByName(gomock.Any(), gomock.Any()).Return(suite.PersonRoot, nil)
		suite.PersonRepoMock.EXPECT().ListWithRelationships(gomock.Any(), gomock.Any()).Return([]*entity.Person{}, nil)
		suite.GenealogyMock.EXPECT().BuildFamilyTree(gomock.Any(), gomock.Any(), gomock.Any(), 1).Return([]*entity.Relative{})

		service := NewService(suite.GenealogyMock, suite.PersonRepoMock, suite.RelationshipRepoMock)
		relationship, err := service.DetermineRelationship(ctx, "John", "Robert")
		assert.Nil(suite.T(), err)
		assert.Equal(suite.T(), "unrelated", relationship)
	})

	suite.Run("should return empty for person not found in the family", func() {
		suite.PersonRepoMock.EXPECT().GetByName(gomock.Any(), gomock.Any()).Return(suite.PersonRoot, nil)
		suite.PersonRepoMock.EXPECT().ListWithRelationships(gomock.Any(), gomock.Any()).Return([]*entity.Person{}, nil)
		suite.GenealogyMock.EXPECT().BuildFamilyTree(gomock.Any(), gomock.Any(), gomock.Any(), 1).Return(suite.FamilyTree)
		service := NewService(suite.GenealogyMock, suite.PersonRepoMock, suite.RelationshipRepoMock)
		relationship, err := service.DetermineRelationship(ctx, "John", "Leon")
		assert.Nil(suite.T(), err)
		assert.Empty(suite.T(), relationship)
	})
}

func (suite *FamilytreeTestSuite) TestCalculateKinshipDistance() {
	ctx := context.Background()
	suite.Run("should return the kinship distance between two people successfully", func() {
		suite.PersonRepoMock.EXPECT().GetByName(gomock.Any(), gomock.Any()).Return(suite.PersonRoot, nil)
		suite.PersonRepoMock.EXPECT().ListWithRelationships(gomock.Any(), gomock.Any()).Return([]*entity.Person{}, nil)
		suite.GenealogyMock.EXPECT().BuildFamilyTree(gomock.Any(), gomock.Any(), gomock.Any(), 1).Return(suite.FamilyTree)
		service := NewService(suite.GenealogyMock, suite.PersonRepoMock, suite.RelationshipRepoMock)
		distance, err := service.CalculateKinshipDistance(ctx, "John", "Robert")
		assert.Nil(suite.T(), err)
		assert.Equal(suite.T(), 1, distance)
	})

	suite.Run("should return an error when trying to get the person", func() {
		suite.PersonRepoMock.EXPECT().GetByName(gomock.Any(), gomock.Any()).Return(nil, errors.New("error database"))
		service := NewService(suite.GenealogyMock, suite.PersonRepoMock, suite.RelationshipRepoMock)
		distance, err := service.CalculateKinshipDistance(ctx, "John", "Robert")
		assert.NotNil(suite.T(), err)
		assert.Error(suite.T(), err, "get person error: %w", "error database")
		assert.Zero(suite.T(), distance)
	})

	suite.Run("should return an error when trying to get the list of people", func() {
		suite.PersonRepoMock.EXPECT().GetByName(gomock.Any(), gomock.Any()).Return(suite.PersonRoot, nil)
		suite.PersonRepoMock.EXPECT().ListWithRelationships(gomock.Any(), gomock.Any()).Return(nil, errors.New("error database"))
		service := NewService(suite.GenealogyMock, suite.PersonRepoMock, suite.RelationshipRepoMock)
		distance, err := service.CalculateKinshipDistance(ctx, "John", "Robert")
		assert.NotNil(suite.T(), err)
		assert.Error(suite.T(), err, "get person error: %w", "error database")
		assert.Zero(suite.T(), distance)
	})

	suite.Run("should return unrelated when the people are not related", func() {
		suite.PersonRepoMock.EXPECT().GetByName(gomock.Any(), gomock.Any()).Return(suite.PersonRoot, nil)
		suite.PersonRepoMock.EXPECT().ListWithRelationships(gomock.Any(), gomock.Any()).Return([]*entity.Person{}, nil)
		suite.GenealogyMock.EXPECT().BuildFamilyTree(gomock.Any(), gomock.Any(), gomock.Any(), 1).Return([]*entity.Relative{})
		service := NewService(suite.GenealogyMock, suite.PersonRepoMock, suite.RelationshipRepoMock)
		distance, err := service.CalculateKinshipDistance(ctx, "John", "Robert")
		assert.Nil(suite.T(), err)
		assert.Zero(suite.T(), distance)
	})

	suite.Run("should return empty for person not found in the family", func() {
		suite.PersonRepoMock.EXPECT().GetByName(gomock.Any(), gomock.Any()).Return(suite.PersonRoot, nil)
		suite.PersonRepoMock.EXPECT().ListWithRelationships(gomock.Any(), gomock.Any()).Return([]*entity.Person{}, nil)
		suite.GenealogyMock.EXPECT().BuildFamilyTree(gomock.Any(), gomock.Any(), gomock.Any(), 1).Return(suite.FamilyTree)
		service := NewService(suite.GenealogyMock, suite.PersonRepoMock, suite.RelationshipRepoMock)
		distance, err := service.CalculateKinshipDistance(ctx, "John", "Leon")
		assert.Nil(suite.T(), err)
		assert.Zero(suite.T(), distance)
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(FamilytreeTestSuite))
}
