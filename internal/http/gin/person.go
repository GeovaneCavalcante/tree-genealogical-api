package gin

import (
	"github.com/GeovaneCavalcante/tree-genealogical/person"
	"github.com/GeovaneCavalcante/tree-genealogical/pkg/logger"
	"github.com/gin-gonic/gin"
)

func createPersonHandler(s person.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p person.Person
		if err := c.BindJSON(&p); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := s.Create(c, &p); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, p)
	}
}

func listPersonHandler(s person.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] List person started")

		filters := map[string]interface{}{}

		persons, err := s.List(c, filters)
		if err != nil {
			logger.Error("[Handler] List person error: ", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if len(persons) == 0 {
			logger.Info("[Handler] List person not found")
			c.JSON(200, []person.Person{})
			return
		}

		logger.Info("[Handler] List person finished")
		c.JSON(200, persons)
	}
}

func getPersonHandler(s person.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] Get person started")

		personID := c.Param("id")

		p, err := s.Get(c, personID)
		if err != nil {
			logger.Error("[Handler] Get person error: ", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if p == nil {
			logger.Info("[Handler] Get person not found")
			c.JSON(404, gin.H{"error": "person not found"})
			return
		}

		logger.Info("[Handler] Get person finished")
		c.JSON(200, p)
	}
}

func updatePersonHandler(s person.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p person.Person
		if err := c.BindJSON(&p); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		personID := c.Param("id")

		if err := s.Update(c, personID, &p); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, p)
	}
}

func deletePersonHandler(s person.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		personID := c.Param("id")

		if err := s.Delete(c, personID); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(204, nil)
	}
}

func MakePersonHandlers(r *gin.RouterGroup, s person.UseCase) {
	r.Handle("POST", "/", createPersonHandler(s))
	r.Handle("GET", "/", listPersonHandler(s))
	r.Handle("GET", "/:id", getPersonHandler(s))
	r.Handle("PUT", "/:id", updatePersonHandler(s))
	r.Handle("DELETE", "/:id", deletePersonHandler(s))
}
