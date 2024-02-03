package gin

import (
	"net/http"

	"github.com/GeovaneCavalcante/tree-genealogical/config"
	"github.com/gin-gonic/gin"
)

func Handlers(envs *config.Environments) *gin.Engine {
	r := gin.Default()

	r.GET("/health", healthHandler)
	r.Group("/api/")

	return r
}

func healthHandler(c *gin.Context) {
	c.String(http.StatusOK, "App is healthy")
}
