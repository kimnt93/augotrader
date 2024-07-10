package handlers

import (
	"augotrader/internal/services"
	"augotrader/internal/types"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CreateLoginRequest = types.LoginInfo

// @Summary Create a login entry
// @Description Create a login entry using the provided login information
// @Accept json
// @Produce json
// @Param loginInfo body LoginInfo true "Login information"
// @Success 201 {object} CreateLoginResponse
// @Router /login [post] or [put]
func CreateLoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginInfo CreateLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginInfo); err != nil {
		response := NewErrorResponse[types.LoginInfo](err.Error())
		json.NewEncoder(w).Encode(response)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	loginInfo, err := services.SetLoginInfoByAccountId(
		loginInfo.AccountId,
		loginInfo.ConsumerId,
		loginInfo.ConsumerSecret,
		loginInfo.PrivateKey,
		loginInfo.AuthToken,
		loginInfo.IsPaperTrading,
		loginInfo.IsDisabled,
	)
	if err != nil {
		response := NewErrorResponse[types.LoginInfo](err.Error())
		json.NewEncoder(w).Encode(response)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := NewSuccessResponse[types.LoginInfo](loginInfo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusCreated)
}

// @Summary Get all login entries
// @Description Get all login entries
// @Produce json
// @Success 200 {object} GetAllLoginsResponse
// @Router /login [get]
func GetAllLoginsHandler(w http.ResponseWriter, r *http.Request) {
	loginInfos, err := services.GetAllLoginInfo()
	if err != nil {
		response := NewErrorResponse[types.LoginInfo](err.Error())
		json.NewEncoder(w).Encode(response)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := NewSuccessResponse[types.LoginInfo](loginInfos)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// @Summary Get login entry by account ID
// @Description Get login entry by account ID
// @Produce json
// @Param accountId path string true "Account ID"
// @Success 200 {object} GetLoginByAccountIdResponse
// @Router /login/{accountId} [get]
func GetLoginByAccountIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["accountId"]

	loginInfo, err := services.GetLoginInfoByAccountId(accountId)
	if err != nil {
		response := NewErrorResponse[types.LoginInfo](err.Error())
		json.NewEncoder(w).Encode(response)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := NewSuccessResponse[types.LoginInfo](loginInfo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// @Summary Delete login entry by account ID
// @Description Delete login entry by account ID
// @Produce json
// @Param accountId path string true "Account ID"
// @Success 200 {object} DeleteLoginByAccountIdResponse
// @Router /login/{accountId} [delete]
func DeleteLoginInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["accountId"]

	loginInfo, err := services.DeleteLoginInfoByAccountId(accountId)
	if err != nil {
		response := NewErrorResponse[types.LoginInfo](err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	response := NewSuccessResponse[types.LoginInfo](loginInfo)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
