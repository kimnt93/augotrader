package routes

import (
	"augotrader/internal/handlers"

	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	// Serve Swagger UI
	r.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/templates")))

	// login and account services
	r.HandleFunc("/login", handlers.CreateLoginHandler).Methods("POST")
	r.HandleFunc("/login", handlers.CreateLoginHandler).Methods("PUT")
	r.HandleFunc("/login/{accountId}", handlers.GetLoginByAccountIdHandler).Methods("GET")
	r.HandleFunc("/login/{accountId}", handlers.DeleteLoginInfoHandler).Methods("DELETE")
	r.HandleFunc("/logins", handlers.GetAllLoginsHandler).Methods("GET")

	// config

	// offset
	r.HandleFunc("/booksize", handlers.SetBooksizeConfigHandler).Methods("POST")
	r.HandleFunc("/booksize", handlers.SetBooksizeConfigHandler).Methods("PUT")
	r.HandleFunc("/booksize/{accountId}", handlers.GetBooksizeByAccountHandler).Methods("GET")

	return r
}
