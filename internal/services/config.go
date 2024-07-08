package services

import (
	"augotrader/internal/cache"
	"augotrader/internal/static"
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
)

type TradingStrategyWeight struct {
	Weight float64 `json:"weight"`
	Name   string  `json:"name"`
}

func GetAllAccountIds() ([]string, error) {
	key := static.CFG_ALL_ACCOUNT
	return cache.RedisClient.SMembers(cache.Ctx, key).Result()
}

func AddAccountById(accountId string) (bool, error) {
	key := static.CFG_ALL_ACCOUNT
	err := cache.RedisClient.SAdd(cache.Ctx, key, accountId).Err()
	if err != nil {
		log.Error().Msgf("Failed to add account id: %v", err)
		return false, err
	}
	return true, nil
}

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

func GetTradingStrategyWeightByAccount(accountId string) ([]TradingStrategyWeight, error) {
	key := fmt.Sprintf("%s.%s", static.CH_ACCOUNT_STRATEGY_WEIGHTS, accountId)
	// Use GET to retrieve the JSON string from Redis
	jsonStr, err := cache.GetKeyStr(key)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON string into a map
	var result []TradingStrategyWeight
	// Load into list
	err = json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		log.Error().Msgf("Failed to unmarshal JSON string (%s): %v", jsonStr, err)
		return nil, err
	}

	return result, nil
}
