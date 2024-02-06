package gin

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_familytree "github.com/GeovaneCavalcante/tree-genealogical/familytree/mock"
	mock_person "github.com/GeovaneCavalcante/tree-genealogical/person/mock"
	mock_relationship "github.com/GeovaneCavalcante/tree-genealogical/relationship/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type badYAML struct{}

func (b badYAML) MarshalYAML() (interface{}, error) {
	return nil, errors.New("expected error")
}

type TestStruct struct {
	Name string `json:"name" xml:"name" yaml:"name"`
}

type errorReader struct{}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("simulated read error")
}

type HandlersTestSuite struct {
	suite.Suite
	FamilyTreeService   *mock_familytree.MockUseCase
	PersonService       *mock_person.MockUseCase
	RelationshipService *mock_relationship.MockUseCase
}

func (suite *HandlersTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	suite.FamilyTreeService = mock_familytree.NewMockUseCase(ctrl)
	suite.PersonService = mock_person.NewMockUseCase(ctrl)
	suite.RelationshipService = mock_relationship.NewMockUseCase(ctrl)
}

func (suite *HandlersTestSuite) TestHandlers() {
	suite.T().Run("Should return a gin.Engine", func(t *testing.T) {
		r := Handlers(nil, suite.PersonService, suite.RelationshipService, suite.FamilyTreeService)
		assert.NotNil(t, r)
		assert.IsType(t, &gin.Engine{}, r)
	})
}

func (suite *HandlersTestSuite) TestHealthHandler() {
	suite.T().Run("Should return a string", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		healthHandler(c)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "App is healthy", w.Body.String())
	})
}

func (suite *HandlersTestSuite) TestRespondAccept() {
	suite.T().Run("Should return a JSON response", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		testData := gin.H{"message": "ok"}

		tests := []struct {
			acceptHeader    string
			expectedStatus  int
			expectedContent string
		}{
			{"application/json", http.StatusOK, `{"message":"ok"}`},
			{"application/xml", http.StatusOK, `<map><message>ok</message></map>`},
			{"application/x-yaml", http.StatusOK, "message: ok\n"},
			{"text/yaml", http.StatusOK, "message: ok\n"},
			{"", http.StatusOK, `{"message":"ok"}`},
		}

		for _, tt := range tests {
			suite.T().Run(tt.acceptHeader, func(t *testing.T) {
				req := httptest.NewRequest("GET", "/test", nil)
				req.Header.Set("Accept", tt.acceptHeader)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = req

				respondAccept(c, tt.expectedStatus, testData)

				assert.Equal(t, tt.expectedStatus, w.Code, "Status code should match")
				assert.Contains(t, w.Body.String(), tt.expectedContent, "Response body should contain the expected content")
			})
		}
	})

	suite.T().Run("Should return an error response", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		testData := badYAML{}

		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("Accept", "application/x-yaml")
		w := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(w)
		c.Request = req

		respondAccept(c, http.StatusOK, testData)

		assert.Equal(t, http.StatusInternalServerError, w.Code, "Status code should be 500")

		assert.Contains(t, w.Body.String(), "Internal Server Error", "Response body should contain the expected error message")
	})
}

func (suite *HandlersTestSuite) TestBindData() {
	suite.T().Run("Should bind JSON data", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		tests := []struct {
			contentType string
			requestBody string
			expected    TestStruct
			expectError bool
		}{
			{
				contentType: "application/json",
				requestBody: `{"name":"John"}`,
				expected:    TestStruct{Name: "John"},
				expectError: false,
			},
			{
				contentType: "application/xml",
				requestBody: `<TestStruct><name>John</name></TestStruct>`,
				expected:    TestStruct{Name: "John"},
				expectError: false,
			},
			{
				contentType: "application/x-yaml",
				requestBody: `name: John`,
				expected:    TestStruct{Name: "John"},
				expectError: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.contentType, func(t *testing.T) {
				req := httptest.NewRequest("POST", "/test", bytes.NewBufferString(tt.requestBody))
				req.Header.Set("Content-Type", tt.contentType)
				w := httptest.NewRecorder()

				c, _ := gin.CreateTestContext(w)
				c.Request = req

				var obj TestStruct
				err := bindData(c, &obj)

				if tt.expectError {
					assert.Error(t, err, "Expected an error")
				} else {
					assert.NoError(t, err, "Did not expect an error")
					assert.Equal(t, tt.expected, obj, "Expected object to be correctly deserialized")
				}
			})
		}
	})

	suite.T().Run("Should return an error when binding invalid YAML data", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		errorTests := []struct {
			contentType string
			requestBody string
			description string
		}{
			{
				contentType: "application/json",
				requestBody: `{"name": John}`,
				description: "Malformed JSON",
			},
			{
				contentType: "application/xml",
				requestBody: `<TestStruct><name>John</name></wrongTag>`,
				description: "Malformed XML",
			},
			{
				contentType: "application/x-yaml",
				requestBody: `name: "John:`,
				description: "Malformed YAML",
			},
		}

		for _, tt := range errorTests {
			t.Run(tt.description, func(t *testing.T) {

				req := httptest.NewRequest("POST", "/test", bytes.NewBufferString(tt.requestBody))
				req.Header.Set("Content-Type", tt.contentType)
				w := httptest.NewRecorder()

				c, _ := gin.CreateTestContext(w)
				c.Request = req

				var obj interface{}
				err := bindData(c, &obj)

				assert.Error(t, err, "Expected an error for %s", tt.description)
			})
		}
	})

	suite.T().Run("Should return an error when binding invalid content type", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		req := httptest.NewRequest("POST", "/test", &errorReader{})
		req.Header.Set("Content-Type", "application/x-yaml")
		w := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(w)
		c.Request = req

		var obj interface{}
		err := bindData(c, &obj)

		assert.Error(t, err, "Expected a read error")
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(HandlersTestSuite))
	suite.Run(t, new(FamilyTreeHandlersTestSuite))
	suite.Run(t, new(PersonHandlersTestSuite))
	suite.Run(t, new(RelationshipHandlersTestSuite))
}
