package services

import (
	"augotrader/internal/cache"
	"augotrader/internal/static"
	"augotrader/internal/types"
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
)

func GetBooksizeByAccount(accountId string) ([]types.AccountBookSizeConfig, error) {
	var result []types.AccountBookSizeConfig

	// Define the pattern to match keys in Redis
	pattern := fmt.Sprintf("%s.%s.*", static.CH_ACCOUNT_BOOKSIZE_CONFIG, accountId)

	// Find all keys that match the pattern
	keys, err := cache.RedisClient.Keys(cache.Ctx, pattern).Result()
	if err != nil {
		return result, err
	}

	// Iterate through each key found
	for _, key := range keys {
		// Get the JSON string from Redis
		jsonStr, err := cache.GetKeyStr(key)
		if err != nil {
			// Log or handle error, but continue to the next key
			continue
		}

		// Unmarshal the JSON string into a map
		var booksizeConfig types.AccountBookSizeConfig
		err = json.Unmarshal([]byte(jsonStr), &booksizeConfig)
		if err != nil {
			// Log or handle error, but continue to the next key
			continue
		}

		// Append the successfully unmarshaled configuration to the result slice
		result = append(result, booksizeConfig)
	}

	// Return an empty slice if no configurations were found
	return result, nil
}

func GetCurrentBooksize(accountId string, symbol string) (types.AccountBookSizeConfig, error) {
	var result types.AccountBookSizeConfig

	key := fmt.Sprintf("%s.%s.%s", static.CH_ACCOUNT_BOOKSIZE_CONFIG, accountId, symbol)
	jsonStr, err := cache.GetKeyStr(key)
	if err != nil {
		return result, err
	}

	// Load into list
	err = json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		log.Error().Msgf("Failed to unmarshal JSON string (%s): %v", jsonStr, err)
		return result, err
	}

	return result, nil
}

func DeleteCurrentBookSize(accountId string, symbol string) (types.AccountBookSizeConfig, error) {
	accountBooksize, err := GetCurrentBooksize(accountId, symbol)
	if err != nil {
		return accountBooksize, err
	}
	key := fmt.Sprintf("%s.%s.%s", static.CH_ACCOUNT_BOOKSIZE_CONFIG, accountId, symbol)
	_, err = cache.DeleteKey(key)
	if err != nil {
		return accountBooksize, err
	}
	return accountBooksize, nil
}

func SetCurrentBookSize(accountId string, symbol string, target_position, target_offset float64, is_disabled bool) (types.AccountBookSizeConfig, error) {
	// Create new booksize config object
	booksizeConfig := types.AccountBookSizeConfig{
		AccountId:      accountId,
		Symbol:         symbol,
		TargetPosition: target_position,
		Offset:         target_offset,
		IsDisabled:     is_disabled,
	}
	key := fmt.Sprintf("%s.%s.%s", static.CH_ACCOUNT_BOOKSIZE_CONFIG, accountId, symbol)

	strBooksizeConfig, err := json.Marshal(booksizeConfig)
	if err != nil {
		return booksizeConfig, err
	}

	_, err = cache.SetKeyStr(key, strBooksizeConfig)
	if err != nil {
		return booksizeConfig, err
	}
	return booksizeConfig, nil
}
