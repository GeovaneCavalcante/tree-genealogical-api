package main

import (
	"log"

	"github.com/GeovaneCavalcante/tree-genealogical/config"
	"github.com/GeovaneCavalcante/tree-genealogical/database"
	"github.com/GeovaneCavalcante/tree-genealogical/familytree"
	"github.com/GeovaneCavalcante/tree-genealogical/internal/http/gin"
	"github.com/GeovaneCavalcante/tree-genealogical/internal/http/webserver"
	"github.com/GeovaneCavalcante/tree-genealogical/person"
	personInmemRepo "github.com/GeovaneCavalcante/tree-genealogical/person/inmem"
	"github.com/GeovaneCavalcante/tree-genealogical/pkg/genealogy"
	"github.com/GeovaneCavalcante/tree-genealogical/relationship"
	relationshipInmemRepo "github.com/GeovaneCavalcante/tree-genealogical/relationship/inmem"
)

// @title Tree Genealogical API
// @version 1.0
// @description This is a simple API to manage genealogical trees
// @host localhost:8080
// @BasePath /api/v1
func main() {

	inmenDB := database.New()

	envs := config.LoadEnvVars()

	personRepo := personInmemRepo.NewPersonRepository(inmenDB)
	personService := person.NewService(personRepo)

	relationshipRepo := relationshipInmemRepo.NewRelationshipRepository(inmenDB)
	relationshipService := relationship.NewService(relationshipRepo)

	genealogy := genealogy.NewFamilyTree()
	familytreeService := familytree.NewService(genealogy, personRepo, relationshipRepo)

	h := gin.Handlers(envs, personService, relationshipService, familytreeService)

	if err := webserver.Start(envs.APIPort, h); err != nil {
		log.Fatalf("Failed to start API: %v", err)
	}
}
