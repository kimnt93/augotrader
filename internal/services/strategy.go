package services

import (
	"augotrader/internal/cache"
	"augotrader/internal/static"
	"augotrader/internal/types"
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
)

func GetTradingStrategyWeightByAccount(accountId string) ([]types.TradingStrategyWeight, error) {
	key := fmt.Sprintf("%s.%s", static.CH_ACCOUNT_STRATEGY_WEIGHTS, accountId)
	// Use GET to retrieve the JSON string from Redis
	jsonStr, err := cache.GetKeyStr(key)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON string into a map
	var result []types.TradingStrategyWeight
	// Load into list
	err = json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		log.Error().Msgf("Failed to unmarshal JSON string (%s): %v", jsonStr, err)
		return nil, err
	}

	return result, nil
}

func GetCurrentSignalByName(name string) (types.Signal, error) {
	// Init default signal
	var signal types.Signal
	signal.Name = name
	signal.Position = 0.0

	// Get signal from hash
	key := fmt.Sprintf("%s.%s", static.CH_SIGNAL_CHANNEL, name)
	signalStr, err := cache.GetKeyStr(key)
	if err != nil {
		return signal, err
	}

	// Unmarshal the JSON string into the Signal struct
	err = json.Unmarshal([]byte(signalStr), &signal)
	if err != nil {
		log.Error().Msgf("Failed to unmarshal signal (%s): %v", signalStr, err)
	}

	return signal, nil
}
