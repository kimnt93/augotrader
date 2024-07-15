package main

import (
	"augotrader/internal/routes"
	"net/http"

	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// Initialize the router and setup routes
	apiRouter := routes.SetupAPIRoutes()

	// Setup a file server for serving static HTML files
	fileServer := http.FileServer(http.Dir("./web"))
	http.Handle("/", fileServer)

	// Combine API routes and file server handler
	http.Handle("/api/", http.StripPrefix("/api", apiRouter))

	// Serve Swagger UI
	http.Handle("/docs/", httpSwagger.Handler(
		httpSwagger.URL("/docs/swagger.json"), // URL to swagger.json
	))

	// Start the server
	port := ":8080"
	log.Info().Msgf("Server starting on port %s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Server stopped")
	}
}
