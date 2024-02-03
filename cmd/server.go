package main

import (
	"log"

	"github.com/GeovaneCavalcante/tree-genealogical/config"
	"github.com/GeovaneCavalcante/tree-genealogical/internal/http/gin"
	"github.com/GeovaneCavalcante/tree-genealogical/internal/http/webserver"
)

func main() {
	envs := config.LoadEnvVars()

	h := gin.Handlers(envs)

	if err := webserver.Start(envs.APIPort, h); err != nil {
		log.Fatalf("Failed to start API: %v", err)
	}
}
