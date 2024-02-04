package gin

import (
	"github.com/GeovaneCavalcante/tree-genealogical/familytree"
	"github.com/gin-gonic/gin"
)

func findFamilyMembersHandler(s familytree.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		personName := c.Param("personName")
		relatives, err := s.GetAllFamilyMembers(c, personName)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, relatives)
	}
}

func MakeFamilyTreeHandlers(r *gin.RouterGroup, s familytree.UseCase) {
	r.Handle("GET", "/:personName", findFamilyMembersHandler(s))
}
