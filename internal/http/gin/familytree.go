package gin

import (
	"net/http"

	"github.com/GeovaneCavalcante/tree-genealogical/familytree"
	"github.com/GeovaneCavalcante/tree-genealogical/internal/http/presenter"
	"github.com/GeovaneCavalcante/tree-genealogical/pkg/logger"
	"github.com/gin-gonic/gin"
)

// @Summary Find family members
// @Description Find family members
// @Tags familytree
// @Accept json,xml
// @Produce json,xml
// @Param personName path string true "Person Name"
// @Success 200 {object} presenter.FamilyTreeResponse
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 500 {object} errorResponse
// @Router /familytree/members/{personName} [get]
func findFamilyMembersHandler(s familytree.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] Find family members started")
		personName := c.Param("personName")

		if IsEmpty(personName) {
			logger.Error("[Handler] Find family members error: personName should not be empty", nil)
			respondAccept(c, http.StatusBadRequest, gin.H{"error": "personName should not be empty"})
			return
		}

		relatives, err := s.GetAllFamilyMembers(c, personName)
		if err != nil {
			logger.Error("[Handler] Find family members error: ", err)
			respondAccept(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		r := presenter.NewFamilyTreeResponse(relatives)

		logger.Info("[Handler] Find family members finished")

		respondAccept(c, http.StatusOK, r)
	}
}

// @Summary Determine relationship
// @Description Determine relationship
// @Tags familytree
// @Accept json,xml
// @Produce json,xml
// @Param firstPersonName path string true "First Person Name"
// @Param secondPersonName path string true "Second Person Name"
// @Success 200 {object} presenter.DetermineRelationResponse
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 500 {object} errorResponse
// @Router /familytree/relationship/{firstPersonName}/{secondPersonName} [get]
func determineRelationshipHandler(s familytree.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] Determine relationship started")
		firstPersonName := c.Param("firstPersonName")
		secondPersonName := c.Param("secondPersonName")

		if firstPersonName == secondPersonName {
			logger.Error("[Handler] Determine relationship error: firstPersonName and secondPersonName should be different", nil)
			respondAccept(c, http.StatusBadRequest, gin.H{"error": "firstPersonName and secondPersonName should be different"})
			return
		}

		if IsEmpty(firstPersonName) || IsEmpty(secondPersonName) {
			logger.Error("[Handler] Determine relationship error: firstPersonName and secondPersonName should not be empty", nil)
			respondAccept(c, http.StatusBadRequest, gin.H{"error": "firstPersonName and secondPersonName should not be empty"})
			return
		}

		relationship, err := s.DetermineRelationship(c, firstPersonName, secondPersonName)
		if err != nil {
			logger.Error("[Handler] Determine relationship error: ", err)
			respondAccept(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		determineRelationResponse := presenter.NewDetermineRelationResponse(relationship)

		logger.Info("[Handler] Determine relationship finished")

		respondAccept(c, http.StatusOK, determineRelationResponse)

	}
}

// @Summary Determine kinship distance
// @Description Determine kinship distance
// @Tags familytree
// @Accept json,xml
// @Produce json,xml
// @Param firstPersonName path string true "First Person Name"
// @Param secondPersonName path string true "Second Person Name"
// @Success 200 {object} presenter.KinshipDistanceResponse
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 500 {object} errorResponse
// @Router /familytree/kinship/distance/{firstPersonName}/{secondPersonName} [get]
func determineKinshipHandler(s familytree.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] Determine kinship started")
		firstPersonName := c.Param("firstPersonName")
		secondPersonName := c.Param("secondPersonName")

		if firstPersonName == secondPersonName {
			logger.Error("[Handler] Determine kinship error: firstPersonName and secondPersonName should be different", nil)
			respondAccept(c, http.StatusBadRequest, gin.H{"error": "firstPersonName and secondPersonName should be different"})
			return
		}

		if IsEmpty(firstPersonName) || IsEmpty(secondPersonName) {
			logger.Error("[Handler] Determine kinship error: firstPersonName and secondPersonName should not be empty", nil)
			respondAccept(c, http.StatusBadRequest, gin.H{"error": "firstPersonName and secondPersonName should not be empty"})
			return
		}

		distance, err := s.CalculateKinshipDistance(c, firstPersonName, secondPersonName)
		if err != nil {
			logger.Error("[Handler] Determine kinship error: ", err)
			respondAccept(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		kinshipDistanceResponse := presenter.NewKinshipDistanceResponse(distance)

		logger.Info("[Handler] Determine kinship finished")

		respondAccept(c, http.StatusOK, kinshipDistanceResponse)
	}
}

func MakeFamilyTreeHandlers(r *gin.RouterGroup, s familytree.UseCase) {
	r.Handle("GET", "/members/:personName", findFamilyMembersHandler(s))
	r.Handle("GET", "/relationship/:firstPersonName/:secondPersonName", determineRelationshipHandler(s))
	r.Handle("GET", "/kinship/distance/:firstPersonName/:secondPersonName", determineKinshipHandler(s))
}
