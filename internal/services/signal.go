package services

import (
	"augotrader/internal/cache"
	"augotrader/internal/static"
	"augotrader/internal/types"
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
)

func SetCurrentSignal(signalType string, signalTime string, name string, symbol string, price float64, position float64) (types.Signal, error) {
	var signal = types.Signal{
		SinalType: signalType,
		Time:      signalTime,
		Name:      name,
		Symbol:    symbol,
		Price:     price,
		Position:  position,
	}
	// Marshal the Signal struct into a JSON string
	signalStr, err := json.Marshal(signal)
	if err != nil {
		log.Error().Msgf("Failed to marshal signal (%v): %v", signal, err)
		return signal, err
	}

	// Set signal to hash
	key := fmt.Sprintf("%s.%s", static.CH_LASTEST_SIGNAL, signal.Name)
	_, err = cache.SetKeyStr(key, string(signalStr))
	if err != nil {
		return signal, err
	}

	return signal, nil
}

func GetCurrentSignalByName(name string) (types.Signal, error) {
	// Init default signal
	var signal types.Signal
	signal.Name = name
	signal.Position = 0.0

	// Get signal from hash
	key := fmt.Sprintf("%s.%s", static.CH_LASTEST_SIGNAL, name)
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
	pattern := fmt.Sprintf("%s.%s*", static.CH_LASTEST_SIGNAL, symbol)
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
