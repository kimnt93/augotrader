package types

import "time"

type ReceivedAction struct {
	Symbol string `json:"symbol"`
}

type Signal struct {
	SinalType string    `json:"type"`
	Time      time.Time `json:"time"`
	Name      string    `json:"name"`
	Symbol    string    `json:"symbol"`
	Price     float64   `json:"price"`
	Position  float64   `json:"position"`
}

type AccountConfig struct {
	TartgetPosition float64 `json:"target_position"`
	TargetOffset    float64 `json:"target_offset"`
	AccountDisabled bool    `json:"account_disabled"`
}

type TradingStrategyWeight struct {
	Weight float64 `json:"weight"`
	Name   string  `json:"name"`
}

type ExecuteMeta struct {
	AccountId string
	Symbol    string
	// AccountConfig AccountConfig
	// Strategies    map[string]TradingStrategyWeight // Get from <>.AccountId
}
