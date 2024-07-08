package services

import (
	"augotrader/internal/cache"
	"augotrader/internal/static"
	"fmt"
)

func GetCurrentOffset(accountId string, symbol string) (float64, error) {
	key := fmt.Sprintf("%s.%s.%s", static.CH_TARGET_OFFSET, accountId, symbol)
	return cache.GetKeyFloat(key)
}

func GetTargetPosition(accountId string, symbol string) (float64, error) {
	key := fmt.Sprintf("%s.%s.%s", static.CH_TARGET_POSITION, accountId, symbol)
	return cache.GetKeyFloat(key)
}

func IsDisabledAccount(accountId string) (bool, error) {
	key := fmt.Sprintf("%s.%s", static.CH_ACCOUNT_DISABLED, accountId)
	// Return the boolean value of the key
	return cache.GetKeyBoolean(key)
}
