package main

import (
	"log"

	"github.com/GeovaneCavalcante/tree-genealogical/config"
	"github.com/GeovaneCavalcante/tree-genealogical/internal/http/gin"
	"github.com/GeovaneCavalcante/tree-genealogical/internal/http/webserver"
	"github.com/GeovaneCavalcante/tree-genealogical/person"
	personInmemRepo "github.com/GeovaneCavalcante/tree-genealogical/person/inmem"
)

var persons []*person.Person

func main() {
	envs := config.LoadEnvVars()

	persons = []*person.Person{

		// {
		// 	ID:   "1",
		// 	Name: "Geovane",
		// },
	}

	personRepo := personInmemRepo.NewPersonRepository(persons)
	personService := person.NewService(personRepo)

	h := gin.Handlers(envs, personService)

	if err := webserver.Start(envs.APIPort, h); err != nil {
		log.Fatalf("Failed to start API: %v", err)
	}
}
