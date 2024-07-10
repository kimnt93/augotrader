package handlers

import (
	"augotrader/internal/services"
	"augotrader/internal/types"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAccountBalanceHandler gets the current positions for the given account
// @Summary Get current acciunt balance
// @Description Get current balance for the given account
// @Produce json
// @Param accountId path string true "Account ID"
// @Success 200 {object} GetCurrentPositionsResponse
// @Failure 500 {object} types.DefaultErrorResponse
// @Router /trading/balance/{accountId} [get]
func GetAccountBalanceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["accountId"]

	// Get account balance
	balance, err := services.GetAccountBalance(accountId)
	if err != nil {
		response := NewErrorResponse[types.AccountBalance](err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Return success response
	response := NewSuccessResponse[types.AccountBalance](balance)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// GetAccountPortfolioHandler gets the current symbol, position for the given account
// @Summary Get current account portfolio
// @Description Get current portfolio for the given account
// @Produce json
// @Param accountId path string true "Account ID"
// @Success 200 {object} GetCurrentPortfolioResponse
// @Failure 500 {object} types.DefaultErrorResponse
// @Router /trading/portfolio/{accountId} [get]
func GetAccountPortfolioHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["accountId"]

	// Get account portfolio
	portfolio, err := services.GetAccountPortfolio(accountId)
	if err != nil {
		response := NewErrorResponse[types.AccountPortfolio](err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Return success response
	response := NewSuccessResponse[types.AccountPortfolio](portfolio)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
