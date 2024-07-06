package services

import (
	"augotrader/internal/cache"
	"augotrader/internal/static"
	"augotrader/internal/types"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

func GetTradingStrategyWeight(accountId string) (map[string]types.TradingStrategyWeight, error) {
	key := fmt.Sprintf("%s.%s", static.CH_ACCOUNT_STRATEGY_WEIGHTS_PREFIX, accountId)
	// Use GET to retrieve the JSON string from Redis
	jsonStr, err := cache.RedisClient.Get(cache.Ctx, key).Result()
	if err != nil {
		log.Error().Msgf("Failed to get key (%s): %v", key, err)
		return nil, err
	}

	// Unmarshal the JSON string into a map
	var result map[string]types.TradingStrategyWeight
	err = json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		log.Error().Msgf("Failed to unmarshal JSON: %v", err)
		return nil, err
	}

	return result, nil
}

func GetCurrentOffset(accountId string, symbol string) (float64, error) {
	key := fmt.Sprintf("%s.%s.%s", static.CH_OFFSET_PREFIX, accountId, symbol)
	return cache.GetKeyFloat(key)
}

func IsDisabledAccount(accountId string) (bool, error) {
	key := fmt.Sprintf("%s.%s", static.CH_ACCOUNT_DISABLED_PREFIX, accountId)
	// Return the boolean value of the key
	return cache.GetKeyBoolean(key)
}

// Mock function to check the current position of the symbol
func GetCurrentPosition(accountId string, symbol string) (float64, error) {
	log.Printf("Get position ne")
	return 10, nil
}

// Mock function to update the position of the symbol
func UpdatePosition(accountId string, symbol string, targetPosition, currentPosition int) (bool, error) {
	// Mock implementation, replace with actual logic
	log.Printf("Updating position for %s: %d -> %d\n", symbol, currentPosition, targetPosition)
	return true, nil
}

// Mock function to check the portfolio
func CheckPortfolio(accountId string) string {
	return "Portfolio details"
}

func IsLockedSymbol(symbol string) bool {
	key := fmt.Sprintf("%s.%s", static.CH_LOCKED_PREFIX, symbol)
	exists, err := cache.RedisClient.Exists(cache.Ctx, key).Result()
	if err != nil {
		log.Error().Msgf("Failed to check if symbol is locked: %v", err)
		return true
	}
	return exists == 1
}

func IsLockedAccount(accountId string) bool {
	key := fmt.Sprintf("%s.%s", static.CH_LOCKED_PREFIX, accountId)
	exists, err := cache.RedisClient.Exists(cache.Ctx, key).Result()
	if err != nil {
		log.Error().Msgf("Failed to check if account is locked: %v", err)
		return true
	}
	return exists == 1
}

func IsLockedAccountSymbol(accountId string, symbol string) bool {
	key := fmt.Sprintf("%s.%s.%s", static.CH_LOCKED_PREFIX, accountId, symbol)
	exists, err := cache.RedisClient.Exists(cache.Ctx, key).Result()
	if err != nil {
		log.Error().Msgf("Failed to check if account symbol is locked: %v", err)
		return true
	}
	return exists == 1
}

func LockSymbol(symbol string) int {
	key := fmt.Sprintf("%s.%s", static.CH_LOCKED_PREFIX, symbol)
	err := cache.RedisClient.Set(cache.Ctx, key, 1, time.Duration(static.DEFAULT_LOCK_TTL)*time.Second).Err()
	if err != nil {
		log.Error().Msgf("Failed to lock symbol: %v", err)
		return 0
	}
	return 1
}

func LockAccount(accountId string) int {
	key := fmt.Sprintf("%s.%s", static.CH_LOCKED_PREFIX, accountId)
	err := cache.RedisClient.Set(cache.Ctx, key, 1, time.Duration(static.DEFAULT_LOCK_TTL)*time.Second).Err()
	if err != nil {
		log.Error().Msgf("Failed to lock account: %v", err)
		return 0
	}
	return 1
}

func LockAccountSymbol(accountId string, symbol string) int {
	key := fmt.Sprintf("%s.%s.%s", static.CH_LOCKED_PREFIX, accountId, symbol)
	err := cache.RedisClient.Set(cache.Ctx, key, 1, time.Duration(static.DEFAULT_LOCK_TTL)*time.Second).Err()
	if err != nil {
		log.Error().Msgf("Failed to lock account symbol: %v", err)
		return 0
	}
	return 1
}

func UnlockSymbol(symbol string) int {
	key := fmt.Sprintf("%s.%s", static.CH_LOCKED_PREFIX, symbol)
	err := cache.RedisClient.Del(cache.Ctx, key).Err()
	if err != nil {
		log.Error().Msgf("Failed to unlock symbol: %v", err)
		return 0
	}
	return 1
}

func UnlockAccount(accountId string) int {
	key := fmt.Sprintf("%s.%s", static.CH_LOCKED_PREFIX, accountId)
	err := cache.RedisClient.Del(cache.Ctx, key).Err()
	if err != nil {
		log.Error().Msgf("Failed to unlock account: %v", err)
		return 0
	}
	return 1
}

func UnlockAccountSymbol(accountId string, symbol string) int {
	key := fmt.Sprintf("%s.%s.%s", static.CH_LOCKED_PREFIX, accountId, symbol)
	err := cache.RedisClient.Del(cache.Ctx, key).Err()
	if err != nil {
		log.Error().Msgf("Failed to unlock account symbol: %v", err)
		return 0
	}
	return 1
}

func GetAllAccountIds() ([]string, error) {
	key := static.CH_ALL_ACCOUNT_PREFIX
	return cache.RedisClient.SMembers(cache.Ctx, key).Result()
}
