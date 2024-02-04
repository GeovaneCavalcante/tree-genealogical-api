package gin

import (
	"fmt"
	"io"
	"net/http"

	"github.com/GeovaneCavalcante/tree-genealogical/config"
	_ "github.com/GeovaneCavalcante/tree-genealogical/docs"
	"github.com/GeovaneCavalcante/tree-genealogical/familytree"
	"github.com/GeovaneCavalcante/tree-genealogical/person"
	"github.com/GeovaneCavalcante/tree-genealogical/relationship"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gopkg.in/yaml.v2"
)

type errorResponse struct {
	Error string `json:"error" xml:"error"`
}

func Handlers(envs *config.Environments, personService person.UseCase, relationshipServoce relationship.UseCase, familyTreeService familytree.UseCase) *gin.Engine {
	r := gin.Default()

	r.GET("/health", healthHandler)
	v1 := r.Group("/api/v1")

	url := ginSwagger.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	pG := v1.Group("/person")
	MakePersonHandlers(pG, personService)

	rG := v1.Group("/relationship")
	MakeRelationshipHandlers(rG, relationshipServoce)

	fG := v1.Group("/familytree")
	MakeFamilyTreeHandlers(fG, familyTreeService)

	return r
}

func healthHandler(c *gin.Context) {
	c.String(http.StatusOK, "App is healthy")
}

func respondAccept(c *gin.Context, status int, data interface{}) {
	fmt.Println(c.GetHeader("Accept"))
	switch c.GetHeader("Accept") {
	case "text/xml", "application/xml":
		c.XML(status, data)
	case "application/json":
		c.JSON(status, data)
	case "application/x-yaml", "text/yaml", "text/x-yaml", "application/yaml":
		yamlData, err := yaml.Marshal(data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		c.Data(status, "application/x-yaml", yamlData)
	default:
		c.JSON(status, data)
	}
}

func bindData(c *gin.Context, obj interface{}) error {
	switch c.GetHeader("Content-Type") {
	case "application/xml", "text/xml", "application/json":
		if err := c.ShouldBind(obj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return err
		}
	case "application/x-yaml", "text/yaml", "text/x-yaml", "application/yaml":
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return err
		}
		if err := yaml.Unmarshal(body, obj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid YAML request body"})
			return err
		}
	default:
		if err := c.BindJSON(obj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request body"})
			return err
		}
	}
	return nil
}
