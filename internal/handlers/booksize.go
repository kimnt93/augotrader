package handlers

import (
	"augotrader/internal/services"
	"augotrader/internal/types"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type SetBooksizeConfigRequest = services.AccountBookSizeConfig

type SetBooksizeConfigResponse struct {
	Success bool                           `json:"success"`
	Data    services.AccountBookSizeConfig `json:"data"`
}

type GetBooksizeConfigByAccountResponse struct {
	Success bool                             `json:"success"`
	Data    []services.AccountBookSizeConfig `json:"data"`
}

// IsDisabledAccountHandler checks if an account is disabled
// @Summary Check if account is disabled
// @Description Check if an account is disabled
// @Produce json
// @Param accountId path string true "Account ID"
// @Success 200 {object} types.SetBooksizeConfigResponse
// @Failure 500 {object} types.DefaultErrorResponse
// @Router /booksize [post]
func SetBooksizeConfigHandler(w http.ResponseWriter, r *http.Request) {
	var req SetBooksizeConfigRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response := types.DefaultErrorResponse{Success: false, Error: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	booksizeConfig, err := services.SetCurrentBookSize(req.AccountId, req.Symbol, req.TargetPosition, req.Offset, req.IsDisabled)
	if err != nil {
		response := types.DefaultErrorResponse{Success: false, Error: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	response := SetBooksizeConfigResponse{Success: true, Data: booksizeConfig}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetBooksizeByAccountHandler gets the booksize configuration for an account
// @Summary Get booksize configuration by account
// @Description Get booksize configuration by account
// @Produce json
// @Param accountId path string true "Account ID"
// @Success 200 {object} GetBooksizeConfigByAccountResponse
// @Failure 500 {object} types.DefaultErrorResponse
// @Router /booksize/{accountId} [get]
func GetBooksizeByAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["accountId"]

	booksizeConfigs, err := services.GetBooksizeByAccount(accountId)
	if err != nil {
		response := types.DefaultErrorResponse{Success: false, Error: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	response := GetBooksizeConfigByAccountResponse{Success: true, Data: booksizeConfigs}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
