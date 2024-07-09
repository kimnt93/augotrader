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

func SetTradingStrategyWeightByAccount(accountId string, weights []TradingStrategyWeight) (bool, error) {
	key := fmt.Sprintf("%s.%s", static.CH_ACCOUNT_STRATEGY_WEIGHTS, accountId)
	// Marshal the map into a JSON string
	jsonStr, err := json.Marshal(weights)
	if err != nil {
		log.Error().Msgf("Failed to marshal map (%v): %v", weights, err)
		return false, err
	}

	// Use SET to store the JSON string in Redis
	return cache.SetKeyStr(key, jsonStr)
}
