package services

import (
	"augotrader/internal/cache"
	"augotrader/internal/static"
	"augotrader/internal/types"
	"encoding/json"
	"fmt"
)

func GetCurrentPrice(symbol string) (types.MarketInfo, error) {
	var marketInfo types.MarketInfo

	// Read from redis
	key := fmt.Sprintf("%s.%s", static.CH_LASTEST_PRICE, symbol)
	jsonStr, err := cache.GetKeyStr(key)
	if err != nil {
		return types.MarketInfo{}, err
	}

	// Unmarshal
	err = json.Unmarshal([]byte(jsonStr), &marketInfo)
	if err != nil {
		return types.MarketInfo{}, err
	}
	return marketInfo, nil
}

func GetCurrentRoom(symbol string) (types.ForeignRoomInfo, error) {
	var currentRoom types.ForeignRoomInfo

	// Read from redis
	key := fmt.Sprintf("%s.%s", static.CH_LASTEST_ROOM, symbol)
	jsonStr, err := cache.GetKeyStr(key)
	if err != nil {
		return types.ForeignRoomInfo{}, err
	}

	// Unmarshal
	err = json.Unmarshal([]byte(jsonStr), &currentRoom)
	if err != nil {
		return types.ForeignRoomInfo{}, err
	}
	return currentRoom, nil
}

func GetCurrentOHLC(symbol string) (types.PriceVolumeInfo, error) {
	var currentBar types.PriceVolumeInfo

	// Read from redis
	key := fmt.Sprintf("%s.%s", static.CH_LASTEST_PRICE, symbol)
	jsonStr, err := cache.GetKeyStr(key)
	if err != nil {
		return currentBar, err
	}

	// Unmarshal
	err = json.Unmarshal([]byte(jsonStr), &currentBar)
	if err != nil {
		return currentBar, err
	}
	return currentBar, nil
}
