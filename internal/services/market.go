package services

import (
	"augotrader/internal/cache"
	"augotrader/internal/static"
	"encoding/json"
	"fmt"
)

type MarketInfo struct {
	Symbol       string  `json:"symbol"`
	CurrentPrice float64 `json:"current_price"`
	// top Bid/Ask data
	BidPrice float64
	AskPrice float64
	BidQty   float64
	AskQty   float64
}

type ForeignRoomInfo struct {
	Symbol      string  `json:"symbol"`
	ForeignBuy  float64 `json:"foreign_buy"`
	ForeignSell float64 `json:"foreign_sell"`
	ForeignRoom float64 `json:"foreign_room"`
}

type PriceVolumeInfo struct {
	Symbol string  `json:"symbol"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume float64 `json:"volume"`
	Value  float64 `json:"value"`
}

func GetCurrentPrice(symbol string) (MarketInfo, error) {
	var marketInfo MarketInfo

	// Read from redis
	key := fmt.Sprintf("%s.%s", static.CH_LASTEST_PRICE, symbol)
	jsonStr, err := cache.GetKeyStr(key)
	if err != nil {
		return MarketInfo{}, err
	}

	// Unmarshal
	err = json.Unmarshal([]byte(jsonStr), &marketInfo)
	if err != nil {
		return MarketInfo{}, err
	}
	return marketInfo, nil
}

func GetCurrentRoom(symbol string) (ForeignRoomInfo, error) {
	var currentRoom ForeignRoomInfo

	// Read from redis
	key := fmt.Sprintf("%s.%s", static.CH_LASTEST_ROOM, symbol)
	jsonStr, err := cache.GetKeyStr(key)
	if err != nil {
		return ForeignRoomInfo{}, err
	}

	// Unmarshal
	err = json.Unmarshal([]byte(jsonStr), &currentRoom)
	if err != nil {
		return ForeignRoomInfo{}, err
	}
	return currentRoom, nil
}

func GetCurrentOHLC(symbol string) (PriceVolumeInfo, error) {
	var currentBar PriceVolumeInfo

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
