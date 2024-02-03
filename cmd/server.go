package main

import (
	"log"

	"github.com/GeovaneCavalcante/tree-genealogical/config"
	"github.com/GeovaneCavalcante/tree-genealogical/database"
	"github.com/GeovaneCavalcante/tree-genealogical/internal/http/gin"
	"github.com/GeovaneCavalcante/tree-genealogical/internal/http/webserver"
	"github.com/GeovaneCavalcante/tree-genealogical/person"
	personInmemRepo "github.com/GeovaneCavalcante/tree-genealogical/person/inmem"
	"github.com/GeovaneCavalcante/tree-genealogical/relationship"
	relationshipInmemRepo "github.com/GeovaneCavalcante/tree-genealogical/relationship/inmem"
)

func main() {

	inmenDB := database.New()

	envs := config.LoadEnvVars()

	personRepo := personInmemRepo.NewPersonRepository(inmenDB)
	personService := person.NewService(personRepo)

	relationshipRepo := relationshipInmemRepo.NewRelationshipRepository(inmenDB)
	relationshipService := relationship.NewService(relationshipRepo)

	h := gin.Handlers(envs, personService, relationshipService)

	if err := webserver.Start(envs.APIPort, h); err != nil {
		log.Fatalf("Failed to start API: %v", err)
	}
}
