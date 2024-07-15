package handlers

import (
	"augotrader/internal/services"
	"augotrader/internal/types"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type SetTradingStrategyRequest struct {
	AccountId string  `json:"account_id"`
	Symbol    string  `json:"symbol"`
	Name      string  `json:"name"`
	Weight    float64 `json:"weight"`
}
type DeleteTradingStrategyRequest struct {
	AccountId string `json:"account_id"`
	Symbol    string `json:"symbol"`
	Name      string `json:"name"`
}

// GetTradingStrategyWeightHandler gets the trading strategy weight for the given account
//
//	@Summary		Get trading strategy weight by account
//	@Description	Get trading strategy weight for the given account
//	@Produce		json
//	@Param			accountId	path		string	true	"Account ID"
//	@Success		200			{object}	GetTradingStrategyWeightByAccountResponse
//	@Failure		500			{object}	types.DefaultErrorResponse
//	@Router			/trading/strategy/{accountId}/{symbol} [get]
func GetTradingStrategyWeightsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["accountId"]
	symbol := vars["symbol"]

	// Get account balance
	strategies, err := services.GetTradingStrategyWeights(accountId, symbol)
	if err != nil {
		response := NewErrorResponse[types.TradingStrategyWeight](err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Return success response
	response := NewSuccessResponse[types.TradingStrategyWeight](strategies)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// SetTradingStrategyWeightHandler sets the trading strategy weight for the given account
//
//	@Summary		Set trading strategy weight by account
//	@Description	Set trading strategy weight for the given account
//	@Produce		json
//	@Param			accountId	path		string						true	"Account ID"
//	@Param			strategies	body		SetTradingStrategyRequest	true	"Strategies"
//	@Success		200			{object}	SetTradingStrategyWeightByAccountResponse
//	@Failure		500			{object}	types.DefaultErrorResponse
//	@Router			/trading/strategy [delete]
func SetTradingStrategyWeightHandler(w http.ResponseWriter, r *http.Request) {

	var req SetTradingStrategyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response := NewErrorResponse[types.TradingStrategyWeight](err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	strategies, err := services.SetTradingStrategyWeight(req.AccountId, req.Symbol, req.Name, req.Weight)
	if err != nil {
		response := NewErrorResponse[types.TradingStrategyWeight](err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Return success response
	response := NewSuccessResponse[types.TradingStrategyWeight](strategies)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// DeleteTradingStrategyWeightHandler deletes the trading strategy weight for the given account
// and strategy name
//
// @Summary		Delete trading strategy weight by account
// @Description	Delete trading strategy weight for the given account
// @Produce		json
// @Param			accountId	path		string						true	"Account ID"
// @Param			symbol		path		string						true	"Symbol"
// @Param			name		body		DeleteTradingStrategyRequest	true	"Name"
// @Success		200			{object}	DeleteTradingStrategyWeightByAccountResponse
// @Failure		500			{object}	types.DefaultErrorResponse
// @Router			/trading/strategy [delete]
func DeleteTradingStrategyWeightHandler(w http.ResponseWriter, r *http.Request) {
	var req DeleteTradingStrategyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response := NewErrorResponse[types.TradingStrategyWeight](err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	strategy, err := services.DeleteTradingStrategyWeight(req.AccountId, req.Symbol, req.Name)
	if err != nil {
		response := NewErrorResponse[types.TradingStrategyWeight](err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Return success response
	response := NewSuccessResponse[types.TradingStrategyWeight](strategy)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
