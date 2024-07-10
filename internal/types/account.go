package types

type AccountLoginSummary struct {
	AccountId  string `json:"account_id"`
	IsDisabled bool   `json:"is_disabled"`
}

type LoginInfo struct {
	AccountId      string `json:"account_id"`
	ConsumerId     string `json:"consumer_id"`
	ConsumerSecret string `json:"consumer_secret"`
	PrivateKey     string `json:"private_key"`
	AuthToken      string `json:"auth_token"`
	IsPaperTrading bool   `json:"is_paper_trading"`
	IsDisabled     bool   `json:"is_disabled"`
}
