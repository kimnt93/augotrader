package types

type AccountLockedStatus struct {
	AccountId string `json:"account_id"`
	IsLocked  bool   `json:"is_locked"`
}
