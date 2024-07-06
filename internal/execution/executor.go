package execution

import (
	"encoding/json"

	"augotrader/internal/cache"
	"augotrader/internal/notification"
	"augotrader/internal/services"
	"augotrader/internal/static"
	"augotrader/internal/types"

	"github.com/rs/zerolog/log"
)

func executor(executeMeta types.ExecuteMeta) bool {
	symbol := executeMeta.CurrentSigal.Symbol
	accountId := executeMeta.AccountId

	if services.IsLockedAccountSymbol(accountId, symbol) {
		return false
	}

	// Lock account and symbol before execute trade
	services.LockAccountSymbol(accountId, symbol)

	targetPosition := executeMeta.AccountConfig.TartgetPosition
	targetOffset := executeMeta.AccountConfig.TargetOffset

	currentPosition, err := services.GetCurrentPosition(accountId, symbol)
	if err != nil {
		log.Error().Msgf("Error while getting position (%s, %s): %v", accountId, symbol, err)
		return false
	} else {
		// Calculate expected position based on weights
		for strategyName, strategyValue := range executeMeta.Strategies {
			log.Debug().Msgf("Strategy Name: %s, Weight: %f", strategyName, strategyValue.Weight, strategyValue.Name)
			currentSignal := services.GetCurrentSignal(strategyValue.Name)

			// Add volume
			if currentSignal.Position == 0 {
				// Skip, do nothing
			} else {

			}
		}

	}

	executeQty := currentPosition - (targetPosition + targetOffset)
	if executeQty == 0 {
		// Do nothing
	} else {
		// Iterate over the Strategies map

	}

	// Release the lock
	services.UnlockAccountSymbol(accountId, symbol)
	return true
}

func executeTrade(accountId string, signal types.Signal) bool {

	disabled, err := services.IsDisabledAccount(accountId)
	if err != nil || disabled {
		log.Error().Msgf("Account %s is disabled (disabled=%b) or checking error: %v", accountId, disabled, err)
		return false
	}

	// current offset
	currentOffset, err := services.GetCurrentOffset(accountId, signal.Symbol)
	if err != nil {
		log.Error().Msgf("Error getting current offset: %v", err)
		return false
	}

	// current tradign strategy
	tradingStrategy, err := services.GetTradingStrategyWeight(accountId)
	if err != nil {
		log.Error().Msgf("Error getting trading strategy: %v", err)
		return false
	}

	var executeTradeMeta types.ExecuteMeta = types.ExecuteMeta{
		AccountId:    accountId,
		CurrentSigal: signal,
		AccountConfig: types.AccountConfig{
			TartgetPosition: signal.Position,
			TargetOffset:    currentOffset,
			AccountDisabled: disabled,
		},
		Strategies: tradingStrategy,
	}

	// execute trade
	return executor(executeTradeMeta)
}

// Run is the main function to execute the trade
func Run() {
	pubsub := cache.RedisClient.Subscribe(cache.Ctx, static.CH_SIGNAL_CHANNEL)
	defer pubsub.Close()

	ch := pubsub.Channel()

	for msg := range ch {
		allAccountids, err := services.GetAllAccountIds()
		if err != nil {
			log.Error().Msgf("Error getting all account ids: %v", err)
			continue
		}
		msgPayload := msg.Payload
		log.Info().Msgf("Received signal: %s", msgPayload)

		// Unmarshal the JSON string into the Signal struct
		var currentSignal types.Signal
		if err := json.Unmarshal([]byte(msgPayload), &currentSignal); err != nil {
			log.Error().Msgf("Error unmarshalling signal [%s]: %v", msgPayload, err)
			continue
		}

		// Loop all accounts and execute trade
		for _, accountId := range allAccountids {
			// Run in goroutine for each accountId
			go executeTrade(accountId, currentSignal)
		}

		// Start send Signal to Telegram
		notification.SendSignalToTelegram(currentSignal)
	}
}
