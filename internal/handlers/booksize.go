package handlers

import (
	"augotrader/internal/services"
	"augotrader/internal/types"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type SetBooksizeConfigRequest = types.AccountBookSizeConfig
type DeleteBooksizeConfigRequest struct {
	AccountId string `json:"account_id"`
	Symbol    string `json:"symbol"`
}

// IsDisabledAccountHandler checks if an account is disabled
//
//	@Summary		Check if account is disabled
//	@Description	Check if an account is disabled
//	@Produce		json
//	@Param			accountId	path		string	true	"Account ID"
//	@Success		200			{object}	types.SetBooksizeConfigResponse
//	@Failure		500			{object}	types.DefaultErrorResponse
//	@Router			/booksize [post]
func SetBooksizeConfigHandler(w http.ResponseWriter, r *http.Request) {
	var req SetBooksizeConfigRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response := NewErrorResponse[types.AccountBookSizeConfig](err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	booksizeConfig, err := services.SetCurrentBookSize(req.AccountId, req.Symbol, req.TargetPosition, req.Offset, req.IsDisabled)
	if err != nil {
		response := NewErrorResponse[types.AccountBookSizeConfig](err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	response := NewSuccessResponse[types.AccountBookSizeConfig](booksizeConfig)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetBooksizeByAccountHandler gets the booksize configuration for an account
//
//	@Summary		Get booksize configuration by account
//	@Description	Get booksize configuration by account
//	@Produce		json
//	@Param			accountId	path		string	true	"Account ID"
//	@Success		200			{object}	GetBooksizeConfigByAccountResponse
//	@Failure		500			{object}	types.DefaultErrorResponse
//	@Router			/booksize/{accountId} [get]
func GetBooksizeByAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["accountId"]

	booksizeConfigs, err := services.GetBooksizeByAccount(accountId)

	if err != nil {
		response := NewErrorResponse[types.AccountBookSizeConfig](err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	response := NewSuccessResponse[types.AccountBookSizeConfig](booksizeConfigs)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteBooksizeConfigHandler(w http.ResponseWriter, r *http.Request) {
	var req DeleteBooksizeConfigRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response := NewErrorResponse[types.AccountBookSizeConfig](err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	booksizeConfig, err := services.DeleteCurrentBookSize(req.AccountId, req.Symbol)
	if err != nil {
		response := NewErrorResponse[types.AccountBookSizeConfig](err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	response := NewSuccessResponse[types.AccountBookSizeConfig](booksizeConfig)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
