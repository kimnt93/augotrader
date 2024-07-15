package services

import (
	"augotrader/internal/static/errors"
	"augotrader/internal/types"

	"github.com/rs/zerolog/log"
)

func CreateOrder(accountId string, orderType string, orderPrice float64, orderSide string, orderQty float64) int {
	log.Printf("Created")
	return errors.CREATE_ORDER_ERROR
}

func CancelOrder(accountId string, order types.CreatedOrder) int {
	log.Printf("Created")
	return errors.CREATE_ORDER_ERROR
}

func ModifyOrder(accountId string, order types.CreatedOrder, newPrice float64, newQty float64) int {
	return 1
}

func GetAccountSymbolPendingOrders(accountId string, symbol string) ([]types.CreatedOrder, error) {
	return []types.CreatedOrder{}, nil
}

func GetAccountPendingOrders(accountId string) ([]types.CreatedOrder, error) {
	return []types.CreatedOrder{}, nil
}

func CancelAccountSymbolPendingOrders(accountId string, symbol string) int {
	return errors.CANCEL_ORDER_ERROR
}

func CancelAccountPendingOrders(accountId string) int {
	return errors.CANCEL_ORDER_ERROR
}

func CancelAllPendingOrders() int {
	allAccounts, err := GetAllAccounts()
	if err != nil {
		return errors.CANCEL_ORDER_ERROR
	}
	//  Loop and cancel by account id
	for _, account := range allAccounts {
		CancelAccountPendingOrders(account.AccountId)
	}
	return errors.CANCEL_ORDER_SUCCESS
}
