package types

type CreatedOrder struct {
	CreatedTime string `json:"created_time"`

	// Input parameters when creating order
	AccountId  string  `json:"account_id"`
	OrderId    string  `json:"order_id"`
	Symbol     string  `json:"symbol"`
	OrderSide  string  `json:"order_side"`
	OrderType  string  `json:"order_type"`
	OrderPrice float64 `json:"order_price"`
	OrderQty   float64 `json:"order_qty"`

	// After order is created, the following fields will be updated
	OrderStatus string  `json:"order_status"`
	OsQty       float64 `json:"os_qty"`
	FilledQty   float64 `json:"filled_qty"`
	RemainQty   float64 `json:"remain_qty"`
	OrderTime   string  `json:"order_time"`
	AvgPrice    float64 `json:"avg_price"`
}
