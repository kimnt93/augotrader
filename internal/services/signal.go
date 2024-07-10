package services

import (
	"augotrader/internal/cache"
	"augotrader/internal/static"
	"augotrader/internal/types"
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
)

func GetCurrentSignalByName(name string) (types.Signal, error) {
	// Init default signal
	var signal types.Signal
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

func GetSignalBySymbol(symbol string) ([]types.Signal, error) {
	pattern := fmt.Sprintf("%s.%s.*", static.CH_SIGNAL_CHANNEL, symbol)
	// Init default signal
	var signals []types.Signal

	// Get all signals from hash
	keys, err := cache.GetKeys(pattern)
	if err != nil {
		return signals, err
	}

	// Unmarshal the JSON string into the Signal struct
	for _, key := range keys {
		signalStr, err := cache.GetKeyStr(key)
		if err != nil {
			return signals, err
		}

		var signal types.Signal
		err = json.Unmarshal([]byte(signalStr), &signal)
		if err != nil {
			log.Error().Msgf("Failed to unmarshal signal (%s): %v", signalStr, err)
		}

		signals = append(signals, signal)
	}

	return signals, nil
}
