package services

import (
	"augotrader/internal/cache"
	"augotrader/internal/static"
	"augotrader/internal/types"
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
)

func GetTradingStrategyWeights(accountId string, symbol string) ([]types.TradingStrategyWeight, error) {
	result := []types.TradingStrategyWeight{}
	key := fmt.Sprintf("%s.%s.%s.*", static.CH_ACCOUNT_STRATEGY_WEIGHTS, accountId, symbol)
	// Use GET to retrieve the JSON string from Redis
	keys, err := cache.GetKeys(key)
	if err != nil {
		return result, err
	}

	for _, key := range keys {
		jsonStr, err := cache.GetKeyStr(key)
		if err != nil {
			return result, err
		}
		accountSymbolStrategy := types.TradingStrategyWeight{}
		err = json.Unmarshal([]byte(jsonStr), &accountSymbolStrategy)
		if err != nil {
			log.Error().Msgf("Failed to unmarshal accountSymbolStrategy (%s): %v", jsonStr, err)
			return result, err
		}
		result = append(result, accountSymbolStrategy)
	}
	return result, nil
}

func SetTradingStrategyWeight(accountId string, symbol string, strategyName string, strategyWeight float64) (types.TradingStrategyWeight, error) {
	key := fmt.Sprintf("%s.%s.%s.%s", static.CH_ACCOUNT_STRATEGY_WEIGHTS, accountId, symbol, strategyName)

	accountSymbolStrategy := types.TradingStrategyWeight{
		AccountId: accountId,
		Symbol:    symbol,
		Name:      strategyName,
		Weight:    strategyWeight,
	}

	// Marshal the TradingStrategyWeight struct into a JSON string
	accountSymbolStrategyStr, err := json.Marshal(accountSymbolStrategy)
	if err != nil {
		log.Error().Msgf("Failed to marshal accountSymbolStrategy (%v): %v", accountSymbolStrategy, err)
		return types.TradingStrategyWeight{}, err
	}

	// Use SET to store the JSON string in Redis
	_, err = cache.SetKeyStr(key, accountSymbolStrategyStr)
	if err != nil {
		return types.TradingStrategyWeight{}, err
	}

	return accountSymbolStrategy, nil
}

func getTradingStrategyWeight(accountId string, symbol string, strategyName string) (types.TradingStrategyWeight, error) {
	key := fmt.Sprintf("%s.%s.%s.%s", static.CH_ACCOUNT_STRATEGY_WEIGHTS, accountId, symbol, strategyName)

	// Use GET to retrieve the JSON string from Redis
	strategyStr, err := cache.GetKeyStr(key)
	if err != nil {
		return types.TradingStrategyWeight{}, err
	}

	// unmarsal the JSON string into the TradingStrategyWeight struct
	accountSymbolStrategy := types.TradingStrategyWeight{}
	err = json.Unmarshal([]byte(strategyStr), &accountSymbolStrategy)
	if err != nil {
		log.Error().Msgf("Failed to unmarshal accountSymbolStrategy (%s): %v", strategyStr, err)
		return types.TradingStrategyWeight{}, err
	}
	return accountSymbolStrategy, nil
}

func DeleteTradingStrategyWeight(accountId string, symbol string, strategyName string) (types.TradingStrategyWeight, error) {
	accountSymbolStrategy, err := getTradingStrategyWeight(accountId, symbol, strategyName)
	if err == nil {
		key := fmt.Sprintf("%s.%s.%s.%s", static.CH_ACCOUNT_STRATEGY_WEIGHTS, accountId, symbol, strategyName)
		_, err := cache.DeleteKey(key)
		if err != nil {
			return accountSymbolStrategy, err
		}

		return accountSymbolStrategy, err
	} else {
		// Not found
		return accountSymbolStrategy, nil
	}
}
