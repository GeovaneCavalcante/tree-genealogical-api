package gin

import (
	"net/http"

	"github.com/GeovaneCavalcante/tree-genealogical/config"
	"github.com/GeovaneCavalcante/tree-genealogical/person"
	"github.com/gin-gonic/gin"
)

func Handlers(envs *config.Environments, personService person.UseCase) *gin.Engine {
	r := gin.Default()

	r.GET("/health", healthHandler)
	v1 := r.Group("/api/v1")

	pG := v1.Group("/person")

	MakePersonHandler(pG, personService)

	return r
}

func healthHandler(c *gin.Context) {
	c.String(http.StatusOK, "App is healthy")
}
