package routes

import (
	"augotrader/internal/handlers"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	// Serve Swagger UI
	r.PathPrefix("/#docs").Handler(httpSwagger.WrapHandler)

	// login services
	r.HandleFunc("/login", handlers.CreateLoginHandler).Methods("POST")
	r.HandleFunc("/logins", handlers.GetAllLoginsHandler).Methods("GET")
	r.HandleFunc("/login/{accountId}", handlers.GetLoginByAccountIdHandler).Methods("GET")

	// account services

	return r
}
