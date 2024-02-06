package gin

import (
	"net/http"

	"github.com/GeovaneCavalcante/tree-genealogical/internal/http/presenter"
	"github.com/GeovaneCavalcante/tree-genealogical/pkg/logger"
	"github.com/GeovaneCavalcante/tree-genealogical/relationship"
	"github.com/gin-gonic/gin"
)

// @Summary Create a relationship
// @Description Create a relationship
// @Tags relationship
// @Accept json,xml
// @Produce json,xml
// @Param relationship body presenter.PaternityRelationshipRequest true "Relationship"
// @Success 201 {object} presenter.PaternityRelationshipResponse
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 500 {object} errorResponse
// @Router /relationship [post]
func createRelationshipHandler(s relationship.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] Create relationship started")
		var r presenter.PaternityRelationshipRequest
		if err := bindData(c, &r); err != nil {
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

		rp := presenter.NewPaternityRelationshipResponse(rs)

		logger.Info("[Handler] Create relationship finished")

		respondAccept(c, http.StatusCreated, rp)
	}
}

// @Summary List relationships
// @Description List relationships
// @Tags relationship
// @Accept json,xml
// @Produce json,xml
// @Success 200 {array} presenter.PaternityRelationshipResponse
// @Failure 500 {object} errorResponse
// @Router /relationship [get]
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
			respondAccept(c, http.StatusOK, []presenter.PaternityRelationshipResponse{})
			return
		}

		logger.Info("[Handler] List relationship finished")

		rP := presenter.NewPaternityRelationshipsResponse(relationships)

		respondAccept(c, http.StatusOK, rP)
	}
}

// @Summary Get a relationship
// @Description Get a relationship
// @Tags relationship
// @Accept json,xml
// @Produce json,xml
// @Param id path string true "Relationship ID"
// @Success 200 {object} presenter.PaternityRelationshipResponse
// @Failure 404 {object} errorResponse "Relationship not found"
// @Failure 500 {object} errorResponse
// @Router /relationship/{id} [get]
func getRelationshipHandler(s relationship.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] Get relationship started")
		relationshipID := c.Param("id")

		if IsEmpty(relationshipID) {
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

		rp := presenter.NewPaternityRelationshipResponse(r)

		respondAccept(c, http.StatusOK, rp)
	}
}

// @Summary Update a relationship
// @Description Update a relationship
// @Tags relationship
// @Accept json,xml
// @Produce json,xml
// @Param id path string true "Relationship ID"
// @Param relationship body presenter.PaternityRelationshipRequest true "Relationship"
// @Success 200 {object} presenter.PaternityRelationshipResponse
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 404 {object} errorResponse "Relationship not found"
// @Failure 500 {object} errorResponse
// @Router /relationship/{id} [put]
func updateRelationshipHandler(s relationship.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] Update relationship started")
		relationshipID := c.Param("id")

		if IsEmpty(relationshipID) {
			logger.Info("[Handler] Update relationship not found")
			respondAccept(c, http.StatusNotFound, gin.H{"error": "relationship not found"})
			return
		}

		var r presenter.PaternityRelationshipRequest
		if err := bindData(c, &r); err != nil {
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

		rp := presenter.NewPaternityRelationshipResponse(rs)

		respondAccept(c, http.StatusOK, rp)
	}
}

// @Summary Delete a relationship
// @Description Delete a relationship
// @Tags relationship
// @Accept json,xml
// @Produce json,xml
// @Param id path string true "Relationship ID"
// @Success 204
// @Failure 404 {object} errorResponse "Relationship not found"
// @Failure 500 {object} errorResponse
// @Router /relationship/{id} [delete]
func deleteRelationshipHandler(s relationship.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] Delete relationship started")
		relationshipID := c.Param("id")

		if IsEmpty(relationshipID) {
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
