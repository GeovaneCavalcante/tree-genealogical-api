package gin

import (
	"net/http"

	"github.com/GeovaneCavalcante/tree-genealogical/config"
	"github.com/GeovaneCavalcante/tree-genealogical/person"
	"github.com/GeovaneCavalcante/tree-genealogical/relationship"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

func Handlers(envs *config.Environments, personService person.UseCase, relationshipServoce relationship.UseCase) *gin.Engine {
	r := gin.Default()

	r.GET("/health", healthHandler)
	v1 := r.Group("/api/v1")

	pG := v1.Group("/person")
	rG := v1.Group("/relationship")

	MakePersonHandlers(pG, personService)
	MakeRelationshipHandlers(rG, relationshipServoce)

	return r
}

func healthHandler(c *gin.Context) {
	c.String(http.StatusOK, "App is healthy")
}

func respondAccept(c *gin.Context, status int, data interface{}) {
	switch c.GetHeader("Accept") {
	case "application/xml":
		c.XML(status, data)
	case "application/json":
		c.JSON(status, data)
	case "application/x-yaml":
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
