package gin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/GeovaneCavalcante/tree-genealogical/internal/entity"
	"github.com/GeovaneCavalcante/tree-genealogical/internal/http/presenter"
	mock_relationship "github.com/GeovaneCavalcante/tree-genealogical/relationship/mock"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type RelationshipHandlersTestSuite struct {
	suite.Suite
	RelationshipService *mock_relationship.MockUseCase
	FamilyGroup         *gin.RouterGroup
	Router              *gin.Engine
	RelationshipInput   *presenter.PaternityRelationshipRequest
	Relationship        *entity.Relationship
	BaseUrl             string
}

func (suite *RelationshipHandlersTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	suite.RelationshipService = mock_relationship.NewMockUseCase(ctrl)
	suite.Router = gin.Default()
	suite.BaseUrl = "/api/v1/relationship/"
	suite.FamilyGroup = suite.Router.Group(suite.BaseUrl)

	MakeRelationshipHandlers(suite.FamilyGroup, suite.RelationshipService)

	suite.RelationshipInput = &presenter.PaternityRelationshipRequest{
		Parent: uuid.New().String(),
		Child:  uuid.New().String(),
	}

	suite.Relationship = &entity.Relationship{
		MainPersonID:    uuid.New().String(),
		SecundePersonID: uuid.New().String(),
	}

}

func (suite *RelationshipHandlersTestSuite) TestCreate() {
	suite.Run("should return success when creating a relationship", func() {
		expectedResponse := fmt.Sprintf("{\"id\":\"\",\"parent\":\"%s\",\"child\":\"%s\"}", suite.RelationshipInput.Parent, suite.RelationshipInput.Child)
		suite.RelationshipService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
		body, _ := json.Marshal(suite.RelationshipInput)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", suite.BaseUrl, bytes.NewBuffer(body))
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusCreated, w.Code)
		assert.Equal(suite.T(), expectedResponse, w.Body.String())
	})

	suite.Run("should return error when creating a relationship", func() {
		suite.RelationshipService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(fmt.Errorf("create relationship error: error"))
		body, _ := json.Marshal(suite.RelationshipInput)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", suite.BaseUrl, bytes.NewBuffer(body))
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"create relationship error: error\"}", w.Body.String())
	})

	suite.Run("should return error when creating a relationship with invalid data", func() {
		invalidRelationship := &presenter.PaternityRelationshipRequest{}
		body, _ := json.Marshal(invalidRelationship)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", suite.BaseUrl, bytes.NewBuffer(body))
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"Key: 'PaternityRelationshipRequest.Parent' Error:Field validation for 'Parent' failed on the 'required' tag\\nKey: 'PaternityRelationshipRequest.Child' Error:Field validation for 'Child' failed on the 'required' tag\"}", w.Body.String())
	})

	suite.Run("should return error when creating a relationship with invalid json", func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", suite.BaseUrl, bytes.NewBuffer([]byte("invalid")))
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"invalid character 'i' looking for beginning of value\"}", w.Body.String())
	})
}

func (suite *RelationshipHandlersTestSuite) TestList() {
	suite.Run("should return success when listing relationships", func() {
		expectedResponse := fmt.Sprintf("[{\"id\":\"\",\"parent\":\"%s\",\"child\":\"%s\"}]", suite.Relationship.SecundePersonID, suite.Relationship.MainPersonID)
		suite.RelationshipService.EXPECT().List(gomock.Any(), gomock.Any()).Return([]*entity.Relationship{suite.Relationship}, nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", suite.BaseUrl, nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.Equal(suite.T(), expectedResponse, w.Body.String())
	})

	suite.Run("should return error when listing relationships", func() {
		suite.RelationshipService.EXPECT().List(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("list relationship error: error"))
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", suite.BaseUrl, nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"list relationship error: error\"}", w.Body.String())
	})

	suite.Run("should return success when listing relationships with no data", func() {
		suite.RelationshipService.EXPECT().List(gomock.Any(), gomock.Any()).Return([]*entity.Relationship{}, nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", suite.BaseUrl, nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.Equal(suite.T(), "[]", w.Body.String())
	})
}

func (suite *RelationshipHandlersTestSuite) TestGet() {
	suite.Run("should return success when getting a relationship", func() {
		expectedResponse := fmt.Sprintf("{\"id\":\"\",\"parent\":\"%s\",\"child\":\"%s\"}", suite.Relationship.SecundePersonID, suite.Relationship.MainPersonID)
		suite.RelationshipService.EXPECT().Get(gomock.Any(), gomock.Any()).Return(suite.Relationship, nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", suite.BaseUrl+uuid.New().String(), nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.Equal(suite.T(), expectedResponse, w.Body.String())
	})

	suite.Run("should return error when getting a relationship", func() {
		suite.RelationshipService.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("get relationship error: error"))
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", suite.BaseUrl+uuid.New().String(), nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"get relationship error: error\"}", w.Body.String())
	})

	suite.Run("should return error when getting a relationship empty", func() {
		suite.RelationshipService.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", suite.BaseUrl+uuid.New().String(), nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusNotFound, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"relationship not found\"}", w.Body.String())
	})

	suite.Run("should return error when getting a relationship with invalid id", func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", suite.BaseUrl+" ", nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusNotFound, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"relationship not found\"}", w.Body.String())
	})

}

func (suite *RelationshipHandlersTestSuite) TestUpdate() {
	suite.Run("should return success when updating a relationship", func() {
		expectedResponse := fmt.Sprintf("{\"id\":\"\",\"parent\":\"%s\",\"child\":\"%s\"}", suite.RelationshipInput.Parent, suite.RelationshipInput.Child)
		suite.RelationshipService.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		body, _ := json.Marshal(suite.RelationshipInput)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", suite.BaseUrl+uuid.New().String(), bytes.NewBuffer(body))
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.Equal(suite.T(), expectedResponse, w.Body.String())
	})

	suite.Run("should return error when updating a relationship", func() {
		suite.RelationshipService.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).Return(fmt.Errorf("update relationship error: error"))
		body, _ := json.Marshal(suite.RelationshipInput)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", suite.BaseUrl+uuid.New().String(), bytes.NewBuffer(body))
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"update relationship error: error\"}", w.Body.String())
	})

	suite.Run("should return error when updating a relationship with invalid data", func() {
		invalidRelationship := &presenter.PaternityRelationshipRequest{}
		body, _ := json.Marshal(invalidRelationship)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", suite.BaseUrl+uuid.New().String(), bytes.NewBuffer(body))
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"Key: 'PaternityRelationshipRequest.Parent' Error:Field validation for 'Parent' failed on the 'required' tag\\nKey: 'PaternityRelationshipRequest.Child' Error:Field validation for 'Child' failed on the 'required' tag\"}", w.Body.String())
	})

	suite.Run("should return error when updating a relationship with invalid json", func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", suite.BaseUrl+uuid.New().String(), bytes.NewBuffer([]byte("invalid")))
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"invalid character 'i' looking for beginning of value\"}", w.Body.String())
	})

	suite.Run("should return error when updating a relationship with invalid id", func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", suite.BaseUrl+" ", nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusNotFound, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"relationship not found\"}", w.Body.String())
	})
}

func (suite *RelationshipHandlersTestSuite) TestDelete() {
	suite.Run("should return success when deleting a relationship", func() {
		suite.RelationshipService.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", suite.BaseUrl+uuid.New().String(), nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusNoContent, w.Code)
		assert.Equal(suite.T(), "", w.Body.String())
	})

	suite.Run("should return error when deleting a relationship", func() {
		suite.RelationshipService.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(fmt.Errorf("delete relationship error: error"))
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", suite.BaseUrl+uuid.New().String(), nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"delete relationship error: error\"}", w.Body.String())
	})

	suite.Run("should return error when deleting a relationship with invalid id", func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", suite.BaseUrl+" ", nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusNotFound, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"relationship not found\"}", w.Body.String())
	})
}
