package services

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

func GetCurrentPositions(accountId string) ([]CurrentPostion, error) {
	return []CurrentPostion{}, nil
}

func GetCurrentPosition(accountId string, symbol string) (CurrentPostion, error) {
	return CurrentPostion{}, nil
}

func GetAccountPortfolio(accountId string) (AccountPortfolio, error) {
	return AccountPortfolio{}, nil
}

func GetAccountBalance(accountId string) (AccountBalance, error) {
	return AccountBalance{}, nil
}
