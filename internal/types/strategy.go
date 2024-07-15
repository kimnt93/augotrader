package types

type TradingStrategyWeight struct {
	AccountId string  `json:"account_id"`
	Symbol    string  `json:"symbol"`
	Name      string  `json:"name"`
	Weight    float64 `json:"weight"`
}
