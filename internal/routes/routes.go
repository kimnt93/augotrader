package routes

import (
	"augotrader/internal/handlers"

	"github.com/gorilla/mux"
)

func SetupAPIRoutes() *mux.Router {
	r := mux.NewRouter()

	// Create a subrouter for v1 API
	v1 := r.PathPrefix("/v1").Subrouter()

	// Login and account services
	v1.HandleFunc("/login", handlers.CreateLoginHandler).Methods("POST")
	v1.HandleFunc("/login", handlers.CreateLoginHandler).Methods("PUT")
	v1.HandleFunc("/login/{accountId}", handlers.GetLoginByAccountIdHandler).Methods("GET")
	v1.HandleFunc("/login/{accountId}", handlers.DeleteLoginInfoHandler).Methods("DELETE")
	v1.HandleFunc("/logins", handlers.GetAllLoginsHandler).Methods("GET")

	// Config

	// Offset
	v1.HandleFunc("/booksize", handlers.SetBooksizeConfigHandler).Methods("POST")
	v1.HandleFunc("/booksize", handlers.SetBooksizeConfigHandler).Methods("PUT")
	v1.HandleFunc("/booksize", handlers.DeleteBooksizeConfigHandler).Methods("DELETE")
	v1.HandleFunc("/booksize/{accountId}", handlers.GetBooksizeByAccountHandler).Methods("GET")

	// View detail
	// Signal by name and symbol
	v1.HandleFunc("/signal", handlers.SetSignalHandler).Methods("POST")
	v1.HandleFunc("/signal/symbol/{symbol}", handlers.GetSignalBySymbolHandler).Methods("GET")
	v1.HandleFunc("/signal/name/{name}", handlers.GetSignalByNameHandler).Methods("GET")

	// Trading
	v1.HandleFunc("/trading/balance/{accountId}", handlers.GetAccountBalanceHandler).Methods("GET")
	v1.HandleFunc("/trading/portfolio/{accountId}", handlers.GetAccountPortfolioHandler).Methods("GET")

	// Strategies
	// We do not want to delete strategies
	v1.HandleFunc("/trading/strategy/{accountId}/{symbol}", handlers.GetTradingStrategyWeightsHandler).Methods("GET")
	v1.HandleFunc("/trading/strategy", handlers.SetTradingStrategyWeightHandler).Methods("POST")
	v1.HandleFunc("/trading/strategy", handlers.SetTradingStrategyWeightHandler).Methods("PUT")
	v1.HandleFunc("/trading/strategy", handlers.DeleteTradingStrategyWeightHandler).Methods("DELETE")

	return r
}
