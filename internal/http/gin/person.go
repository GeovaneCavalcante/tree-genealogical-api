package gin

import (
	"net/http"

	"github.com/GeovaneCavalcante/tree-genealogical/person"
	"github.com/GeovaneCavalcante/tree-genealogical/pkg/logger"
	"github.com/gin-gonic/gin"
)

func createPersonHandler(s person.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] Create person started")
		var p person.Person
		if err := c.BindJSON(&p); err != nil {
			logger.Error("[Handler] Create person error: ", err)
			respondAccept(c, http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := s.Create(c, &p); err != nil {
			logger.Error("[Handler] Create person error: ", err)
			respondAccept(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		logger.Info("[Handler] Create person finished")
		respondAccept(c, http.StatusCreated, p)
	}
}

func listPersonHandler(s person.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] List person started")

		filters := map[string]interface{}{}

		persons, err := s.List(c, filters)
		if err != nil {
			logger.Error("[Handler] List person error: ", err)
			respondAccept(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if len(persons) == 0 {
			logger.Info("[Handler] List person not found")
			respondAccept(c, http.StatusOK, []person.Person{})
			return
		}

		logger.Info("[Handler] List person finished")
		respondAccept(c, http.StatusOK, persons)
	}
}

func getPersonHandler(s person.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] Get person started")

		personID := c.Param("id")

		if personID == "" {
			logger.Info("[Handler] Get person not found")
			respondAccept(c, http.StatusNotFound, gin.H{"error": "person not found"})
			return
		}

		p, err := s.Get(c, personID)
		if err != nil {
			logger.Error("[Handler] Get person error: ", err)
			respondAccept(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if p == nil {
			logger.Info("[Handler] Get person not found")
			respondAccept(c, http.StatusNotFound, gin.H{"error": "person not found"})
			return
		}

		logger.Info("[Handler] Get person finished")
		respondAccept(c, http.StatusOK, p)
	}
}

func updatePersonHandler(s person.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] Update person started")
		personID := c.Param("id")

		if personID == "" {
			logger.Info("[Handler] Update person not found")
			respondAccept(c, http.StatusNotFound, gin.H{"error": "person not found"})
			return
		}

		var p person.Person
		if err := c.BindJSON(&p); err != nil {
			logger.Error("[Handler] Update person error: ", err)
			respondAccept(c, http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := s.Update(c, personID, &p); err != nil {
			logger.Error("[Handler] Update person error: ", err)
			respondAccept(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		logger.Info("[Handler] Update person finished")
		respondAccept(c, http.StatusOK, p)
	}
}

func deletePersonHandler(s person.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		personID := c.Param("id")

		if personID == "" {
			logger.Info("[Handler] Delete person not found")
			respondAccept(c, http.StatusNotFound, gin.H{"error": "person not found"})
			return
		}

		logger.Info("[Handler] Delete person started")

		if err := s.Delete(c, personID); err != nil {
			logger.Error("[Handler] Delete person error: ", err)
			respondAccept(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		logger.Info("[Handler] Delete person finished")
		respondAccept(c, http.StatusNoContent, nil)
	}
}

func MakePersonHandlers(r *gin.RouterGroup, s person.UseCase) {
	r.Handle("POST", "/", createPersonHandler(s))
	r.Handle("GET", "/", listPersonHandler(s))
	r.Handle("GET", "/:id", getPersonHandler(s))
	r.Handle("PUT", "/:id", updatePersonHandler(s))
	r.Handle("DELETE", "/:id", deletePersonHandler(s))
}
