package gin

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/GeovaneCavalcante/tree-genealogical/internal/entity"
	"github.com/GeovaneCavalcante/tree-genealogical/internal/http/presenter"
	mock_person "github.com/GeovaneCavalcante/tree-genealogical/person/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type PersonHandlersTestSuite struct {
	suite.Suite
	PersonService *mock_person.MockUseCase
	FamilyGroup   *gin.RouterGroup
	Router        *gin.Engine
	PersonInput   *presenter.PersonRequest
	Person        *entity.Person
	BaseUrl       string
}

func (suite *PersonHandlersTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	suite.PersonService = mock_person.NewMockUseCase(ctrl)
	suite.Router = gin.Default()
	suite.BaseUrl = "/api/v1/person/"
	suite.FamilyGroup = suite.Router.Group(suite.BaseUrl)

	MakePersonHandlers(suite.FamilyGroup, suite.PersonService)

	suite.PersonInput = &presenter.PersonRequest{
		Name:   "John",
		Gender: "M",
	}

	suite.Person = &entity.Person{
		Name:   "John",
		Gender: "M",
	}

}

func (suite *PersonHandlersTestSuite) TestCreate() {
	suite.Run("should return success when creating a person", func() {
		expectedResponse := "{\"id\":\"\",\"name\":\"John\",\"gender\":\"M\"}"
		suite.PersonService.EXPECT().Create(gomock.Any(), suite.Person).Return(nil)
		body, _ := json.Marshal(suite.PersonInput)

		req, err := http.NewRequest("POST", suite.BaseUrl, bytes.NewBuffer(body))

		assert.Nil(suite.T(), err)
		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusCreated, w.Code)
		assert.Equal(suite.T(), expectedResponse, w.Body.String())
	})

	suite.Run("should return error when creating a person", func() {
		suite.PersonService.EXPECT().Create(gomock.Any(), suite.Person).Return(errors.New("error creating person"))
		body, _ := json.Marshal(suite.PersonInput)

		req, err := http.NewRequest("POST", suite.BaseUrl, bytes.NewBuffer(body))

		assert.Nil(suite.T(), err)
		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"error creating person\"}", w.Body.String())
	})

	suite.Run("should return error when creating a person with invalid data", func() {

		req, err := http.NewRequest("POST", suite.BaseUrl, bytes.NewBuffer([]byte("invalid data")))

		assert.Nil(suite.T(), err)
		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"invalid character 'i' looking for beginning of value\"}", w.Body.String())
	})

	suite.Run("should return error when creating a person with invalid person", func() {

		suite.PersonInput.Name = ""
		body, _ := json.Marshal(suite.PersonInput)

		req, err := http.NewRequest("POST", suite.BaseUrl, bytes.NewBuffer(body))

		assert.Nil(suite.T(), err)
		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"Key: 'PersonRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag\"}", w.Body.String())
	})

}

func (suite *PersonHandlersTestSuite) TestList() {
	suite.Run("should return success when listing persons", func() {
		expectedResponse := "[{\"id\":\"\",\"name\":\"John\",\"gender\":\"M\"}]"
		suite.PersonService.EXPECT().List(gomock.Any(), gomock.Any()).Return([]*entity.Person{suite.Person}, nil)

		req, err := http.NewRequest("GET", suite.BaseUrl, nil)

		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.Equal(suite.T(), expectedResponse, w.Body.String())
	})

	suite.Run("should return error when listing persons", func() {
		suite.PersonService.EXPECT().List(gomock.Any(), gomock.Any()).Return(nil, errors.New("error listing persons"))

		req, err := http.NewRequest("GET", suite.BaseUrl, nil)

		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"error listing persons\"}", w.Body.String())
	})

	suite.Run("should return when listing persons empity", func() {
		expectedResponse := "[]"
		suite.PersonService.EXPECT().List(gomock.Any(), gomock.Any()).Return([]*entity.Person{}, nil)

		req, err := http.NewRequest("GET", suite.BaseUrl, nil)

		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.Equal(suite.T(), expectedResponse, w.Body.String())
	})

}

func (suite *PersonHandlersTestSuite) TestGet() {
	suite.Run("should return success when getting a person", func() {
		expectedResponse := "{\"id\":\"\",\"name\":\"John\",\"gender\":\"M\"}"
		suite.PersonService.EXPECT().Get(gomock.Any(), "1").Return(suite.Person, nil)

		req, err := http.NewRequest("GET", suite.BaseUrl+"1", nil)

		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.Equal(suite.T(), expectedResponse, w.Body.String())
	})

	suite.Run("should return error when getting a person", func() {
		suite.PersonService.EXPECT().Get(gomock.Any(), "1").Return(nil, errors.New("error getting person"))

		req, err := http.NewRequest("GET", suite.BaseUrl+"1", nil)

		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"error getting person\"}", w.Body.String())
	})

	suite.Run("should return error when getting a person not found", func() {
		suite.PersonService.EXPECT().Get(gomock.Any(), "1").Return(nil, nil)

		req, err := http.NewRequest("GET", suite.BaseUrl+"1", nil)

		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusNotFound, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"person not found\"}", w.Body.String())
	})

	suite.Run("should return error when getting a person with invalid id", func() {
		req, err := http.NewRequest("GET", suite.BaseUrl+" ", nil)

		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusNotFound, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"person not found\"}", w.Body.String())
	})
}

func (suite *PersonHandlersTestSuite) TestUpdate() {
	suite.Run("should return success when updating a person", func() {
		expectedResponse := "{\"id\":\"\",\"name\":\"John\",\"gender\":\"M\"}"
		suite.PersonService.EXPECT().Update(gomock.Any(), "1", suite.Person).Return(nil)
		body, _ := json.Marshal(suite.PersonInput)

		req, err := http.NewRequest("PUT", suite.BaseUrl+"1", bytes.NewBuffer(body))

		assert.Nil(suite.T(), err)
		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.Equal(suite.T(), expectedResponse, w.Body.String())
	})

	suite.Run("should return error when updating a person", func() {
		suite.PersonService.EXPECT().Update(gomock.Any(), "1", suite.Person).Return(errors.New("error updating person"))
		body, _ := json.Marshal(suite.PersonInput)

		req, err := http.NewRequest("PUT", suite.BaseUrl+"1", bytes.NewBuffer(body))

		assert.Nil(suite.T(), err)
		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"error updating person\"}", w.Body.String())
	})

	suite.Run("should return error when updating a person with invalid data", func() {

		req, err := http.NewRequest("PUT", suite.BaseUrl+"1", bytes.NewBuffer([]byte("invalid data")))

		assert.Nil(suite.T(), err)
		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"invalid character 'i' looking for beginning of value\"}", w.Body.String())
	})

	suite.Run("should return error when updating a person with invalid person", func() {

		suite.PersonInput.Name = ""
		body, _ := json.Marshal(suite.PersonInput)

		req, err := http.NewRequest("PUT", suite.BaseUrl+"1", bytes.NewBuffer(body))

		assert.Nil(suite.T(), err)
		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"Key: 'PersonRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag\"}", w.Body.String())
	})

	suite.Run("should return error when updating a person with invalid id", func() {
		req, err := http.NewRequest("PUT", suite.BaseUrl+" ", nil)

		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusNotFound, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"person not found\"}", w.Body.String())
	})
}

func (suite *PersonHandlersTestSuite) TestDelete() {
	suite.Run("should return success when deleting a person", func() {
		suite.PersonService.EXPECT().Delete(gomock.Any(), "1").Return(nil)

		req, err := http.NewRequest("DELETE", suite.BaseUrl+"1", nil)

		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusNoContent, w.Code)
	})

	suite.Run("should return error when deleting a person", func() {
		suite.PersonService.EXPECT().Delete(gomock.Any(), "1").Return(errors.New("error deleting person"))

		req, err := http.NewRequest("DELETE", suite.BaseUrl+"1", nil)

		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"error deleting person\"}", w.Body.String())
	})

	suite.Run("should return error when deleting a person with invalid id", func() {
		req, err := http.NewRequest("DELETE", suite.BaseUrl+" ", nil)

		w := httptest.NewRecorder()
		assert.Nil(suite.T(), err)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusNotFound, w.Code)
		assert.Equal(suite.T(), "{\"error\":\"person not found\"}", w.Body.String())
	})
}
