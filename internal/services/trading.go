package services

import "augotrader/internal/types"

func GetCurrentPositions(accountId string) ([]types.CurrentPostion, error) {
	return []types.CurrentPostion{}, nil
}

func GetCurrentPosition(accountId string, symbol string) (types.CurrentPostion, error) {
	return types.CurrentPostion{}, nil
}

func GetAccountPortfolio(accountId string) (types.AccountPortfolio, error) {
	return types.AccountPortfolio{}, nil
}

func GetAccountBalance(accountId string) (types.AccountBalance, error) {
	return types.AccountBalance{}, nil
}
