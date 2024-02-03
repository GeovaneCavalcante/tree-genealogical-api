package gin

import (
	"github.com/GeovaneCavalcante/tree-genealogical/relationship"
	"github.com/gin-gonic/gin"
)

func createRelationshipHandler(s relationship.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r relationship.Relationship
		if err := c.BindJSON(&r); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := s.Create(c, &r); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, r)
	}
}

func listRelationshipHandler(s relationship.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		filters := map[string]interface{}{}

		relationships, err := s.List(c, filters)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if len(relationships) == 0 {
			c.JSON(200, []relationship.Relationship{})
			return
		}

		c.JSON(200, relationships)
	}
}

func getRelationshipHandler(s relationship.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		relationshipID := c.Param("id")

		r, err := s.Get(c, relationshipID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, r)
	}
}

func updateRelationshipHandler(s relationship.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		relationshipID := c.Param("id")
		var r relationship.Relationship
		if err := c.BindJSON(&r); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := s.Update(c, relationshipID, &r); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, r)
	}
}

func deleteRelationshipHandler(s relationship.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		relationshipID := c.Param("id")

		if err := s.Delete(c, relationshipID); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(204, nil)
	}
}

func MakeRelationshipHandlers(r *gin.RouterGroup, s relationship.UseCase) {
	r.POST("", createRelationshipHandler(s))
	r.GET("", listRelationshipHandler(s))
	r.GET("/:id", getRelationshipHandler(s))
	r.PUT("/:id", updateRelationshipHandler(s))
	r.DELETE("/:id", deleteRelationshipHandler(s))
}
