package services

import (
	"augotrader/internal/cache"
	"augotrader/internal/static"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

func IsLockedAccountSymbol(accountId string, symbol string) (bool, error) {
	key := fmt.Sprintf("%s.%s.%s", static.CH_ACCOUNT_SYMBOL_LOCKED, accountId, symbol)
	exists, err := cache.RedisClient.Exists(cache.Ctx, key).Result()
	if err != nil {
		log.Error().Msgf("Failed to check if account symbol is locked: %v", err)
		return true, err
	}
	return exists == 1, nil
}

func LockAccountSymbol(accountId string, symbol string) (bool, error) {
	key := fmt.Sprintf("%s.%s.%s", static.CH_ACCOUNT_SYMBOL_LOCKED, accountId, symbol)
	err := cache.RedisClient.Set(cache.Ctx, key, 1, time.Duration(static.DEFAULT_LOCK_TTL)*time.Second).Err()
	if err != nil {
		log.Error().Msgf("Failed to lock account symbol: %v", err)
		return false, err
	}
	return true, nil
}

func UnlockAccountSymbol(accountId string, symbol string) (bool, error) {
	key := fmt.Sprintf("%s.%s.%s", static.CH_ACCOUNT_SYMBOL_LOCKED, accountId, symbol)
	err := cache.RedisClient.Del(cache.Ctx, key).Err()
	if err != nil {
		log.Error().Msgf("Failed to unlock account symbol: %v", err)
		return false, err
	}
	return true, nil
}
