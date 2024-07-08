package services

import (
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

type PendingOrder struct {
	AccountId  string  `json:"account_id"`
	OrderId    string  `json:"order_id"`
	Symbol     string  `json:"symbol"`
	OrderSide  string  `json:"order_side"`
	OrderType  string  `json:"order_type"`
	OrderPrice float64 `json:"order_price"`
	OrderQty   int     `json:"order_qty"`
}

func CreateOrder(accountId string, orderType string, orderPrice float64, orderSide string) bool {
	log.Printf("Created")
	return true
}

func CancelOrder(accountId string, orderId string) bool {
	log.Printf("Created")
	return true
}

func ModifyOrder() int {
	return 1
}

func GetSymbolPendingOrders(accountId string, symbol string) ([]PendingOrder, error) {
	return []PendingOrder{}, nil
}

func GetAllPendingOrders(accountId string) ([]PendingOrder, error) {
	return []PendingOrder{}, nil
}
