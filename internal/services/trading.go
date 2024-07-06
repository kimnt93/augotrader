package services

import (
	"augotrader/internal/cache"
	"augotrader/internal/static"
	"augotrader/internal/types"
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
)

type CreateOrderRequest struct {
	AccountId  string  `json:"account_id"`
	Symbol     string  `json:"symbol"`
	OrderSide  string  `json:"order_side"`
	OrderType  string  `json:"order_type"`
	OrderPrice float64 `json:"order_price"`
	OrderQty   int     `json:"order_qty"`

	// Stop order for futures
	StopOrder  bool    `json:"stop_order,omitempty"`
	StopPrice  float64 `json:"stop_price,omitempty"`
	StopType   string  `json:"stop_type,omitempty"`
	StopStep   float64 `json:"stop_step,omitempty"`
	LossStep   float64 `json:"loss_step,omitempty"`
	ProfitStep float64 `json:"profit_step,omitempty"`
}

type CreateOrderResponse struct {
	AccountId  string  `json:"account_id"`
	Symbol     string  `json:"symbol"`
	OrderSide  string  `json:"order_side"`
	OrderType  string  `json:"order_type"`
	OrderPrice float64 `json:"order_price"`
	OrderQty   int     `json:"order_qty"`

	// Stop order for futures
	StopOrder  bool    `json:"stop_order,omitempty"`
	StopPrice  float64 `json:"stop_price,omitempty"`
	StopType   string  `json:"stop_type,omitempty"`
	StopStep   float64 `json:"stop_step,omitempty"`
	LossStep   float64 `json:"loss_step,omitempty"`
	ProfitStep float64 `json:"profit_step,omitempty"`
}

type CancelOrderRequest struct {
}

type CancelOrderResponse struct {
}

func GetCurrentSignal(name string) types.Signal {
	// Init default signal
	var signal types.Signal
	signal.Name = name
	signal.Position = 0.0

	// Get signal from hash
	key := fmt.Sprintf("%s.%s", static.CH_ACCOUNT_DISABLED_PREFIX, name)
	signalStr, err := cache.GetKeyStr(key)
	if err == nil {
		// Unmarshal the JSON string into the Signal struct
		err = json.Unmarshal([]byte(signalStr), &signal)
		if err != nil {
			log.Error().Msgf("Failed to unmarshal signal (%s): %v", signalStr, err)
		}
	}

	return signal
}

func CreateOrder(accountId string, orderType string) bool {
	log.Printf("Created")
}

func CancelOrder(accountId string, orderId string) bool {
	log.Printf("Created")
}

func ModifyOrder() int {
	return 1
}
