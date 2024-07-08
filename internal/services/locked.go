package services

import (
	"augotrader/internal/cache"
	"augotrader/internal/static"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

func IsLockedSymbol(symbol string) bool {
	key := fmt.Sprintf("%s.%s", static.CH_ACCOUNT_SYMBOL_LOCKED, symbol)
	exists, err := cache.RedisClient.Exists(cache.Ctx, key).Result()
	if err != nil {
		log.Error().Msgf("Failed to check if symbol is locked: %v", err)
		return true
	}
	return exists == 1
}

func IsLockedAccount(accountId string) bool {
	key := fmt.Sprintf("%s.%s", static.CH_ACCOUNT_SYMBOL_LOCKED, accountId)
	exists, err := cache.RedisClient.Exists(cache.Ctx, key).Result()
	if err != nil {
		log.Error().Msgf("Failed to check if account is locked: %v", err)
		return true
	}
	return exists == 1
}

func IsLockedAccountSymbol(accountId string, symbol string) (bool, error) {
	key := fmt.Sprintf("%s.%s.%s", static.CH_ACCOUNT_SYMBOL_LOCKED, accountId, symbol)
	exists, err := cache.RedisClient.Exists(cache.Ctx, key).Result()
	if err != nil {
		log.Error().Msgf("Failed to check if account symbol is locked: %v", err)
		return true, err
	}
	return exists == 1, nil
}

func LockSymbol(symbol string) int {
	key := fmt.Sprintf("%s.%s", static.CH_ACCOUNT_SYMBOL_LOCKED, symbol)
	err := cache.RedisClient.Set(cache.Ctx, key, 1, time.Duration(static.DEFAULT_LOCK_TTL)*time.Second).Err()
	if err != nil {
		log.Error().Msgf("Failed to lock symbol: %v", err)
		return 0
	}
	return 1
}

func LockAccount(accountId string) int {
	key := fmt.Sprintf("%s.%s", static.CH_ACCOUNT_SYMBOL_LOCKED, accountId)
	err := cache.RedisClient.Set(cache.Ctx, key, 1, time.Duration(static.DEFAULT_LOCK_TTL)*time.Second).Err()
	if err != nil {
		log.Error().Msgf("Failed to lock account: %v", err)
		return 0
	}
	return 1
}

func LockAccountSymbol(accountId string, symbol string) int {
	key := fmt.Sprintf("%s.%s.%s", static.CH_ACCOUNT_SYMBOL_LOCKED, accountId, symbol)
	err := cache.RedisClient.Set(cache.Ctx, key, 1, time.Duration(static.DEFAULT_LOCK_TTL)*time.Second).Err()
	if err != nil {
		log.Error().Msgf("Failed to lock account symbol: %v", err)
		return 0
	}
	return 1
}

func UnlockSymbol(symbol string) int {
	key := fmt.Sprintf("%s.%s", static.CH_ACCOUNT_SYMBOL_LOCKED, symbol)
	err := cache.RedisClient.Del(cache.Ctx, key).Err()
	if err != nil {
		log.Error().Msgf("Failed to unlock symbol: %v", err)
		return 0
	}
	return 1
}

func UnlockAccount(accountId string) int {
	key := fmt.Sprintf("%s.%s", static.CH_ACCOUNT_SYMBOL_LOCKED, accountId)
	err := cache.RedisClient.Del(cache.Ctx, key).Err()
	if err != nil {
		log.Error().Msgf("Failed to unlock account: %v", err)
		return 0
	}
	return 1
}

func UnlockAccountSymbol(accountId string, symbol string) int {
	key := fmt.Sprintf("%s.%s.%s", static.CH_ACCOUNT_SYMBOL_LOCKED, accountId, symbol)
	err := cache.RedisClient.Del(cache.Ctx, key).Err()
	if err != nil {
		log.Error().Msgf("Failed to unlock account symbol: %v", err)
		return 0
	}
	return 1
}

func GetAllAccountIds() ([]string, error) {
	key := static.CH_ACCOUNT_SYMBOL_LOCKED
	return cache.RedisClient.SMembers(cache.Ctx, key).Result()
}
