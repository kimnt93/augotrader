package types

type AccountSymbolLockedStatus struct {
	AccountId string `json:"account_id"`
	Symbol    string `json:"symbol"`
	IsLocked  bool   `json:"is_locked"`
}
