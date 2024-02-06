package gin

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	mock_familytree "github.com/GeovaneCavalcante/tree-genealogical/familytree/mock"
	"github.com/GeovaneCavalcante/tree-genealogical/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type FamilyTreeHandlersTestSuite struct {
	suite.Suite
	FamilyTreeService *mock_familytree.MockUseCase
	FamilyGroup       *gin.RouterGroup
	Router            *gin.Engine
	PersonRoot        *entity.Person
	FamilyTree        []*entity.Relative
	BaseUrl           string
}

func (suite *FamilyTreeHandlersTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	suite.FamilyTreeService = mock_familytree.NewMockUseCase(ctrl)
	suite.Router = gin.Default()
	suite.BaseUrl = "/api/v1/familytree"
	suite.FamilyGroup = suite.Router.Group(suite.BaseUrl)

	MakeFamilyTreeHandlers(suite.FamilyGroup, suite.FamilyTreeService)

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

func (suite *FamilyTreeHandlersTestSuite) TestGetFamilyTree() {

	suite.Run("should return success when getting family tree", func() {

		responseExpected := "{\"members\":[{\"name\":\"John\",\"typeRelationship\":\"Root\",\"relationships\":[]},{\"name\":\"Robert\",\"typeRelationship\":\"Father\",\"relationships\":[]},{\"name\":\"Maria\",\"typeRelationship\":\"Mother\",\"relationships\":[]}]}"
		suite.FamilyTreeService.EXPECT().GetAllFamilyMembers(gomock.Any(), suite.PersonRoot.Name).Return(suite.FamilyTree, nil)

		req, err := http.NewRequest("GET", fmt.Sprintf("%s/members/%s", suite.BaseUrl, suite.PersonRoot.Name), nil)

		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)

		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.Equal(suite.T(), responseExpected, w.Body.String())
	})

	suite.Run("should return error when getting family tree", func() {
		suite.FamilyTreeService.EXPECT().GetAllFamilyMembers(gomock.Any(), suite.PersonRoot.Name).Return(nil, fmt.Errorf("get family tree error"))

		req, err := http.NewRequest("GET", fmt.Sprintf("%s/members/%s", suite.BaseUrl, suite.PersonRoot.Name), nil)

		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)

		assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"get family tree error\"}", w.Body.String())
	})

	suite.Run("should return error when getting family tree with invalid person name", func() {
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/members/%s", suite.BaseUrl, " "), nil)

		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)

		assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"personName should not be empty\"}", w.Body.String())
	})
}

func (suite *FamilyTreeHandlersTestSuite) TestDetermineRelationship() {
	suite.Run("should return success when determining relationship", func() {
		responseExpected := "{\"relationship\":\"Father\"}"
		suite.FamilyTreeService.EXPECT().DetermineRelationship(gomock.Any(), suite.PersonRoot.Name, "Robert").Return("Father", nil)

		req, err := http.NewRequest("GET", fmt.Sprintf("%s/relationship/%s/%s", suite.BaseUrl, suite.PersonRoot.Name, "Robert"), nil)

		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)

		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.Equal(suite.T(), responseExpected, w.Body.String())
	})

	suite.Run("should return success when determining relationship with unrelated", func() {
		responseExpected := "{\"relationship\":\"unrelated\"}"
		suite.FamilyTreeService.EXPECT().DetermineRelationship(gomock.Any(), suite.PersonRoot.Name, "Robert").Return("unrelated", nil)

		req, err := http.NewRequest("GET", fmt.Sprintf("%s/relationship/%s/%s", suite.BaseUrl, suite.PersonRoot.Name, "Robert"), nil)

		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)

		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.Equal(suite.T(), responseExpected, w.Body.String())
	})

	suite.Run("should return error when determining relationship", func() {
		suite.FamilyTreeService.EXPECT().DetermineRelationship(gomock.Any(), suite.PersonRoot.Name, "Robert").Return("", fmt.Errorf("determine relationship error"))

		req, err := http.NewRequest("GET", fmt.Sprintf("%s/relationship/%s/%s", suite.BaseUrl, suite.PersonRoot.Name, "Robert"), nil)

		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)

		assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"determine relationship error\"}", w.Body.String())
	})

	suite.Run("should return error when determining relationship with invalid person name", func() {
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/relationship/%s/%s", suite.BaseUrl, suite.PersonRoot.Name, " "), nil)

		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)

		assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"firstPersonName and secondPersonName should not be empty\"}", w.Body.String())
	})

	suite.Run("should return error when determining relationship with names equals", func() {
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/relationship/%s/%s", suite.BaseUrl, suite.PersonRoot.Name, suite.PersonRoot.Name), nil)

		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)

		assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"firstPersonName and secondPersonName should be different\"}", w.Body.String())
	})
}

func (suite *FamilyTreeHandlersTestSuite) TestDetermineKinship() {
	suite.Run("should return success when determining kinship", func() {
		responseExpected := "{\"distance\":1}"
		suite.FamilyTreeService.EXPECT().CalculateKinshipDistance(gomock.Any(), suite.PersonRoot.Name, "Robert").Return(1, nil)

		req, err := http.NewRequest("GET", fmt.Sprintf("%s/kinship/distance/%s/%s", suite.BaseUrl, suite.PersonRoot.Name, "Robert"), nil)

		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)

		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.Equal(suite.T(), responseExpected, w.Body.String())
	})

	suite.Run("should return error when determining kinship", func() {
		suite.FamilyTreeService.EXPECT().CalculateKinshipDistance(gomock.Any(), suite.PersonRoot.Name, "Robert").Return(0, fmt.Errorf("determine kinship error"))

		req, err := http.NewRequest("GET", fmt.Sprintf("%s/kinship/distance/%s/%s", suite.BaseUrl, suite.PersonRoot.Name, "Robert"), nil)

		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)

		assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"determine kinship error\"}", w.Body.String())
	})

	suite.Run("should return error when determining kinship with invalid person name", func() {
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/kinship/distance/%s/%s", suite.BaseUrl, suite.PersonRoot.Name, " "), nil)

		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)

		assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"firstPersonName and secondPersonName should not be empty\"}", w.Body.String())
	})

	suite.Run("should return error when determining kinship with names equals", func() {
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/kinship/distance/%s/%s", suite.BaseUrl, suite.PersonRoot.Name, suite.PersonRoot.Name), nil)

		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)

		assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"firstPersonName and secondPersonName should be different\"}", w.Body.String())
	})

}
