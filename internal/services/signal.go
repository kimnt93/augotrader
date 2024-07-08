package services

import (
	"augotrader/internal/cache"
	"augotrader/internal/static"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

type Signal struct {
	SinalType string    `json:"type"`
	Time      time.Time `json:"time"`
	Name      string    `json:"name"`
	Symbol    string    `json:"symbol"`
	Price     float64   `json:"price"`
	Position  float64   `json:"position"`
}

func GetCurrentSignalByName(name string) (Signal, error) {
	// Init default signal
	var signal Signal
	signal.Name = name
	signal.Position = 0.0

	// Get signal from hash
	key := fmt.Sprintf("%s.%s", static.CH_SIGNAL_CHANNEL, name)
	signalStr, err := cache.GetKeyStr(key)
	if err != nil {
		return signal, err
	}

	// Unmarshal the JSON string into the Signal struct
	err = json.Unmarshal([]byte(signalStr), &signal)
	if err != nil {
		log.Error().Msgf("Failed to unmarshal signal (%s): %v", signalStr, err)
	}

	return signal, nil
}
