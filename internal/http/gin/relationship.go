package gin

import (
	"net/http"

	"github.com/GeovaneCavalcante/tree-genealogical/pkg/logger"
	"github.com/GeovaneCavalcante/tree-genealogical/relationship"
	"github.com/gin-gonic/gin"
)

func createRelationshipHandler(s relationship.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] Create relationship started")
		var r relationship.Relationship
		if err := c.BindJSON(&r); err != nil {
			logger.Error("[Handler] Create relationship error: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := s.Create(c, &r); err != nil {
			logger.Error("[Handler] Create relationship error: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		logger.Info("[Handler] Create relationship finished")

		c.JSON(201, r)
	}
}

func listRelationshipHandler(s relationship.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] List relationship started")
		filters := map[string]interface{}{}

		relationships, err := s.List(c, filters)
		if err != nil {
			logger.Error("[Handler] List relationship error: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if len(relationships) == 0 {
			logger.Info("[Handler] List relationship not found")
			c.JSON(http.StatusOK, []relationship.Relationship{})
			return
		}

		logger.Info("[Handler] List relationship finished")
		c.JSON(http.StatusOK, relationships)
	}
}

func getRelationshipHandler(s relationship.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] Get relationship started")
		relationshipID := c.Param("id")

		if relationshipID == "" {
			logger.Info("[Handler] Get relationship not found")
			c.JSON(http.StatusNotFound, gin.H{"error": "relationship not found"})
			return
		}

		r, err := s.Get(c, relationshipID)
		if err != nil {
			logger.Error("[Handler] Get relationship error: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if r == nil {
			logger.Info("[Handler] Get relationship not found")
			c.JSON(http.StatusNotFound, gin.H{"error": "relationship not found"})
			return
		}

		logger.Info("[Handler] Get relationship finished")
		c.JSON(http.StatusOK, r)
	}
}

func updateRelationshipHandler(s relationship.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] Update relationship started")
		relationshipID := c.Param("id")

		if relationshipID == "" {
			logger.Info("[Handler] Update relationship not found")
			c.JSON(http.StatusNotFound, gin.H{"error": "relationship not found"})
			return
		}

		var r relationship.Relationship
		if err := c.BindJSON(&r); err != nil {
			logger.Error("[Handler] Update relationship error: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := s.Update(c, relationshipID, &r); err != nil {
			logger.Error("[Handler] Update relationship error: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		logger.Info("[Handler] Update relationship finished")
		c.JSON(http.StatusOK, r)
	}
}

func deleteRelationshipHandler(s relationship.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] Delete relationship started")
		relationshipID := c.Param("id")

		if relationshipID == "" {
			logger.Info("[Handler] Delete relationship not found")
			c.JSON(http.StatusNotFound, gin.H{"error": "relationship not found"})
			return
		}

		if err := s.Delete(c, relationshipID); err != nil {
			logger.Error("[Handler] Delete relationship error: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		logger.Info("[Handler] Delete relationship finished")
		c.JSON(http.StatusNoContent, nil)
	}
}

func MakeRelationshipHandlers(r *gin.RouterGroup, s relationship.UseCase) {
	r.POST("", createRelationshipHandler(s))
	r.GET("", listRelationshipHandler(s))
	r.GET("/:id", getRelationshipHandler(s))
	r.PUT("/:id", updateRelationshipHandler(s))
	r.DELETE("/:id", deleteRelationshipHandler(s))
}
