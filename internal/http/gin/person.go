package gin

import (
	"fmt"
	"net/http"

	"github.com/GeovaneCavalcante/tree-genealogical/internal/entity"
	"github.com/GeovaneCavalcante/tree-genealogical/internal/http/presenter"
	"github.com/GeovaneCavalcante/tree-genealogical/person"
	"github.com/GeovaneCavalcante/tree-genealogical/pkg/logger"
	"github.com/gin-gonic/gin"
)

// @Summary Create a person
// @Description Create a person
// @Tags person
// @Accept json,xml
// @Produce json,xml
// @Param person body presenter.PersonRequest true "Person"
// @Success 201 {object} presenter.PersonResponse
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 500 {object} errorResponse
// @Router /person [post]
func createPersonHandler(s person.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] Create person started")
		var p presenter.PersonRequest
		if err := bindData(c, &p); err != nil {
			logger.Error("[Handler] Create person error: ", err)
			respondAccept(c, http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println(p)

		if err := p.Validate(); err != nil {
			logger.Error("[Handler] Create person error: ", err)
			respondAccept(c, http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		pp := p.ToPerson()

		if err := s.Create(c, pp); err != nil {
			logger.Error("[Handler] Create person error: ", err)
			respondAccept(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		person := presenter.NewPersonResponse(pp)

		logger.Info("[Handler] Create person finished")
		respondAccept(c, http.StatusCreated, person)
	}
}

// @Summary List persons
// @Description List persons
// @Tags person
// @Accept json,xml
// @Produce json,xml
// @Param name query string false "Filter by person's lasted name (no implemeted)"
// @Success 200 {array} presenter.PersonResponse
// @Failure 500 {object} errorResponse
// @Router /person [get]
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
			respondAccept(c, http.StatusOK, []entity.Person{})
			return
		}

		pp := presenter.NewPersonsResponse(persons)

		logger.Info("[Handler] List person finished")
		respondAccept(c, http.StatusOK, pp)
	}
}

// @Summary Get a person
// @Description Get a person
// @Tags person
// @Accept json,xml
// @Produce json,xml
// @Param id path string true "Person ID"
// @Success 200 {object} presenter.PersonResponse
// @Failure 404 {object} errorResponse "Person not found"
// @Failure 500 {object} errorResponse
// @Router /person/{id} [get]
func getPersonHandler(s person.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] Get person started")

		personID := c.Param("id")

		if IsEmpty(personID) {
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

		pp := presenter.NewPersonResponse(p)

		logger.Info("[Handler] Get person finished")
		respondAccept(c, http.StatusOK, pp)
	}
}

// @Summary Update a person
// @Description Update a person
// @Tags person
// @Accept json,xml
// @Produce json,xml
// @Param id path string true "Person ID"
// @Param person body presenter.PersonRequest true "Person"
// @Success 200 {object} presenter.PersonResponse
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 404 {object} errorResponse "Person not found"
// @Failure 500 {object} errorResponse
// @Router /person/{id} [put]
func updatePersonHandler(s person.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("[Handler] Update person started")
		personID := c.Param("id")

		if IsEmpty(personID) {
			logger.Info("[Handler] Update person not found")
			respondAccept(c, http.StatusNotFound, gin.H{"error": "person not found"})
			return
		}

		var p presenter.PersonRequest
		if err := bindData(c, &p); err != nil {
			logger.Error("[Handler] Update person error: ", err)
			respondAccept(c, http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := p.Validate(); err != nil {
			logger.Error("[Handler] Update person error: ", err)
			respondAccept(c, http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		pp := p.ToPerson()

		if err := s.Update(c, personID, pp); err != nil {
			logger.Error("[Handler] Update person error: ", err)
			respondAccept(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		person := presenter.NewPersonResponse(pp)

		logger.Info("[Handler] Update person finished")
		respondAccept(c, http.StatusOK, person)
	}
}

// @Summary Delete a person
// @Description Delete a person
// @Tags person
// @Accept json,xml
// @Produce json,xml
// @Param id path string true "Person ID"
// @Success 204
// @Failure 404 {object} errorResponse "Person not found"
// @Failure 500 {object} errorResponse
// @Router /person/{id} [delete]
func deletePersonHandler(s person.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		personID := c.Param("id")

		if IsEmpty(personID) {
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
