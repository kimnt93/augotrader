package main

import (
	"augotrader/internal/routes"
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {
	r := routes.SetupRoutes()
	log.Info().Msg("Server starting on port 8080")
	log.Fatal().Err(http.ListenAndServe(":8080", r)).Msg("Server stopped")
}
