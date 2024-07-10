package handlers

import (
	"augotrader/internal/services"
	"augotrader/internal/types"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetSignalBySymbolHandler gets the signal by symbol, using in list all signal of given symbol
// @Summary Get signal by symbol
// @Description Get signal by symbol
// @Produce json
// @Param symbol path string true "Symbol"
// @Success 200 {object} GetSignalBySymbolResponse
// @Failure 500 {object} types.DefaultErrorResponse
// @Router /signal/symbol/{symbol} [get]
func GetSignalBySymbolHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	symbol := vars["symbol"]

	signals, err := services.GetSignalBySymbol(symbol)
	if err != nil {
		response := NewErrorResponse[types.Signal](err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	response := NewSuccessResponse[types.Signal](signals)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetSignalByName gets the signal by name
// @Summary Get signal by name
// @Description Get signal by name
// @Produce json
// @Param name path string true "Name"
// @Success 200 {object} GetSignalByNameResponse
// @Failure 500 {object} types.DefaultErrorResponse
// @Router /signal/name/{name} [get]
func GetSignalByNameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	signal, err := services.GetCurrentSignalByName(name)
	if err != nil {
		response := NewErrorResponse[types.Signal](err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	response := NewSuccessResponse[types.Signal](signal)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
