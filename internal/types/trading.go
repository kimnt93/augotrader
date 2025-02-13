package types

type CurrentPostion struct {
	AccountId string  `json:"account_id"`
	Symbol    string  `json:"symbol"`
	Position  float64 `json:"position"`
	AvgPrice  float64 `json:"avg_price"`
}

type AccountPortfolio struct {
	AccountId string `json:"account_id"`
}

type AccountBalance struct {
	AccountId string  `json:"account_id"`
	Balance   float64 `json:"balance"`
}
