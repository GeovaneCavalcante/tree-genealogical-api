package gin

import (
	"net/http"

	"github.com/GeovaneCavalcante/tree-genealogical/internal/http/presenter"
	"github.com/GeovaneCavalcante/tree-genealogical/pkg/logger"
	"github.com/GeovaneCavalcante/tree-genealogical/relationship"
	"github.com/gin-gonic/gin"
)

func createRelationshipHandler(s relationship.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] Create relationship started")
		var r presenter.PaternityRelationship
		if err := c.BindJSON(&r); err != nil {
			logger.Error("[Handler] Create relationship error: ", err)
			respondAccept(c, http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := r.Validate(); err != nil {
			logger.Error("[Handler] Create relationship error: ", err)
			respondAccept(c, http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rs := r.ToRelationship()

		if err := s.Create(c, rs); err != nil {
			logger.Error("[Handler] Create relationship error: ", err)
			respondAccept(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		rp := presenter.NewPaternityRelationship(rs)

		logger.Info("[Handler] Create relationship finished")

		respondAccept(c, http.StatusCreated, rp)
	}
}

func listRelationshipHandler(s relationship.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] List relationship started")
		filters := map[string]interface{}{}

		relationships, err := s.List(c, filters)
		if err != nil {
			logger.Error("[Handler] List relationship error: ", err)
			respondAccept(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if len(relationships) == 0 {
			logger.Info("[Handler] List relationship not found")
			respondAccept(c, http.StatusOK, []presenter.PaternityRelationship{})
			return
		}

		logger.Info("[Handler] List relationship finished")

		rP := presenter.NewPaternityRelationships(relationships)

		respondAccept(c, http.StatusOK, rP)
	}
}

func getRelationshipHandler(s relationship.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] Get relationship started")
		relationshipID := c.Param("id")

		if relationshipID == "" {
			logger.Info("[Handler] Get relationship not found")
			respondAccept(c, http.StatusNotFound, gin.H{"error": "relationship not found"})
			return
		}

		r, err := s.Get(c, relationshipID)
		if err != nil {
			logger.Error("[Handler] Get relationship error: ", err)
			respondAccept(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if r == nil {
			logger.Info("[Handler] Get relationship not found")
			respondAccept(c, http.StatusNotFound, gin.H{"error": "relationship not found"})
			return
		}

		logger.Info("[Handler] Get relationship finished")

		rp := presenter.NewPaternityRelationship(r)

		respondAccept(c, http.StatusOK, rp)
	}
}

func updateRelationshipHandler(s relationship.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] Update relationship started")
		relationshipID := c.Param("id")

		if relationshipID == "" {
			logger.Info("[Handler] Update relationship not found")
			c.JSON(http.StatusNotFound, gin.H{"error": "relationship not found"})
			respondAccept(c, http.StatusNotFound, gin.H{"error": "relationship not found"})
			return
		}

		var r presenter.PaternityRelationship
		if err := c.BindJSON(&r); err != nil {
			logger.Error("[Handler] Update relationship error: ", err)
			respondAccept(c, http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := r.Validate(); err != nil {
			logger.Error("[Handler] Update relationship error: ", err)
			respondAccept(c, http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rs := r.ToRelationship()

		if err := s.Update(c, relationshipID, rs); err != nil {
			logger.Error("[Handler] Update relationship error: ", err)
			respondAccept(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		logger.Info("[Handler] Update relationship finished")

		rp := presenter.NewPaternityRelationship(rs)

		respondAccept(c, http.StatusOK, rp)
	}
}

func deleteRelationshipHandler(s relationship.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] Delete relationship started")
		relationshipID := c.Param("id")

		if relationshipID == "" {
			logger.Info("[Handler] Delete relationship not found")
			respondAccept(c, http.StatusNotFound, gin.H{"error": "relationship not found"})
			return
		}

		if err := s.Delete(c, relationshipID); err != nil {
			logger.Error("[Handler] Delete relationship error: ", err)
			respondAccept(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		logger.Info("[Handler] Delete relationship finished")
		respondAccept(c, http.StatusNoContent, nil)
	}
}

func MakeRelationshipHandlers(r *gin.RouterGroup, s relationship.UseCase) {
	r.POST("", createRelationshipHandler(s))
	r.GET("", listRelationshipHandler(s))
	r.GET("/:id", getRelationshipHandler(s))
	r.PUT("/:id", updateRelationshipHandler(s))
	r.DELETE("/:id", deleteRelationshipHandler(s))
}
