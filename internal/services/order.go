package services

import (
	"augotrader/internal/static/errors"

	"github.com/rs/zerolog/log"
)

type CreatedOrder struct {
	CreatedTime string `json:"created_time"`

	// Input parameters when creating order
	AccountId  string  `json:"account_id"`
	OrderId    string  `json:"order_id"`
	Symbol     string  `json:"symbol"`
	OrderSide  string  `json:"order_side"`
	OrderType  string  `json:"order_type"`
	OrderPrice float64 `json:"order_price"`
	OrderQty   float64 `json:"order_qty"`

	// After order is created, the following fields will be updated
	OrderStatus string  `json:"order_status"`
	OsQty       float64 `json:"os_qty"`
	FilledQty   float64 `json:"filled_qty"`
	RemainQty   float64 `json:"remain_qty"`
	OrderTime   string  `json:"order_time"`
	AvgPrice    float64 `json:"avg_price"`
}

func CreateOrder(accountId string, orderType string, orderPrice float64, orderSide string) int {
	log.Printf("Created")
	return errors.CREATE_ORDER_ERROR
}

func CancelOrder(accountId string, order CreatedOrder) int {
	log.Printf("Created")
	return errors.CREATE_ORDER_ERROR
}

func ModifyOrder(accountId string, order CreatedOrder, newPrice float64, newQty float64) int {
	return 1
}

func GetAccountSymbolPendingOrders(accountId string, symbol string) ([]CreatedOrder, error) {
	return []CreatedOrder{}, nil
}

func GetAccountPendingOrders(accountId string) ([]CreatedOrder, error) {
	return []CreatedOrder{}, nil
}

func CancelAccountSymbolPendingOrders(accountId string, symbol string) int {
	return errors.CANCEL_ORDER_ERROR
}

func CancelAccountPendingOrders(accountId string) int {
	return errors.CANCEL_ORDER_ERROR
}

func CancelAllPendingOrders() int {
	accountIds, err := GetAllAccountIds()
	if err != nil {
		return errors.CANCEL_ORDER_ERROR
	}
	//  Loop and cancel by account id
	for _, accountId := range accountIds {
		CancelAccountPendingOrders(accountId)
	}
	return errors.CANCEL_ORDER_SUCCESS
}
