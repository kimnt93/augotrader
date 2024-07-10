package types

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
