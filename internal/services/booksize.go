package services

import (
	"augotrader/internal/cache"
	"augotrader/internal/static"
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
)

type AccountBookSizeConfig struct {
	AccountId      string  `json:"account_id"`
	Symbol         string  `json:"symbol"`
	TargetPosition float64 `json:"target_position"`
	Offset         float64 `json:"offset"`
	IsDisabled     bool    `json:"is_disabled"`
}

func GetBooksizeByAccount(accountId string) ([]AccountBookSizeConfig, error) {
	var result []AccountBookSizeConfig
	pattern := fmt.Sprintf("%s.%s.*", static.CH_ACCOUNT_BOOKSIZE_CONFIG, accountId)
	// Find all keys that match the pattern
	keys, err := cache.RedisClient.Keys(cache.Ctx, pattern).Result()

	if err != nil {
		return result, err
	}
	for _, key := range keys {
		// Get the JSON string from Redis
		jsonStr, err := cache.GetKeyStr(key)
		// unmashal the JSON string into a map
		var booksizeConfig AccountBookSizeConfig
		err = json.Unmarshal([]byte(jsonStr), &booksizeConfig)
		if err != nil {
			continue
		}
		result = append(result, booksizeConfig)
	}
	return result, nil

}

func GetCurrentBooksize(accountId string, symbol string) (AccountBookSizeConfig, error) {
	var result AccountBookSizeConfig

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

func SetCurrentBookSize(accountId string, symbol string, target_position, target_offset float64, is_disabled bool) (AccountBookSizeConfig, error) {
	// Create new booksize config object
	booksizeConfig := AccountBookSizeConfig{
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
