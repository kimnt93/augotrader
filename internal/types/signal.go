package types

import (
	"time"
)

// Signal name must have format [SYMBOL].[SIGNAL_NAME]
// Then set to key <PREFIX>.[SYMBOL].[SIGNAL_NAME]
type Signal struct {
	SinalType string    `json:"type"`
	Time      time.Time `json:"time"`
	Name      string    `json:"name"`
	Symbol    string    `json:"symbol"`
	Price     float64   `json:"price"`
	Position  float64   `json:"position"`
}
