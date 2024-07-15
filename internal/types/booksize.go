package types

type AccountBookSizeConfig struct {
	AccountId      string  `json:"account_id"`
	Symbol         string  `json:"symbol"`
	TargetPosition float64 `json:"target_position"`
	Offset         float64 `json:"target_offset"`
	IsDisabled     bool    `json:"is_disabled"`
}
