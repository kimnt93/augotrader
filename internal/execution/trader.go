package execution

import (
	"sync"

	"github.com/rs/zerolog/log"

	"augotrader/internal/services"
	"augotrader/internal/types"
)

var (
	activeRequests = make(map[string]bool)
	mutex          sync.Mutex
)

func ExecuteTrade(executeMeta types.ExecuteMeta) bool {
	symbol := executeMeta.CurrentSigal.Symbol
	accountId := executeMeta.AccountId

	if services.IsLockedAccountSymbol(accountId, symbol) {
		return false
	} else {
		// Lock account and symbol before execute trade
		services.LockAccountSymbol(accountId, symbol)

		targetPosition := executeMeta.AccountConfig.TartgetPosition
		targetOffset := executeMeta.AccountConfig.TargetOffset

		currentPosition, err := services.GetCurrentPosition(accountId, symbol)
		if err != nil {
			log.Error().Msgf("Error while getting position (%s, %s): %v", accountId, symbol, err)
			return false
		} else {

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
	}

	// // Lock to check and update activeRequests map
	// mutex.Lock()
	// if activeRequests[symbol] {
	// 	mutex.Unlock()
	// 	fmt.Printf("Task for %s is already running. Ignoring new task.\n", symbol)
	// 	return
	// }
	// activeRequests[symbol] = true
	// mutex.Unlock()

	// // Perform the trade execution
	// currentPosition := services.GetPosition(symbol)
	// success := utils.UpdatePosition(symbol, signal.TargetPosition, currentPosition)

	// if !success {
	// 	// Re-send signal to Redis for retry
	// 	data, _ := json.Marshal(signal)
	// 	config.RedisClient.Publish(config.Ctx, config.SignalChannel, data)
	// }

	// // Remove the task from activeRequests after completion
	// mutex.Lock()
	// delete(activeRequests, symbol)
	// mutex.Unlock()
}
