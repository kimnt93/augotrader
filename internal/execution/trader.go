package execution

import (
	"encoding/json"
	"fmt"
	"sync"

	"../config"
)

var (
	activeRequests = make(map[string]bool)
	mutex          sync.Mutex
)

func ExecuteTrade(signal Signal) {
	symbol := signal.Symbol

	// Lock to check and update activeRequests map
	mutex.Lock()
	if activeRequests[symbol] {
		mutex.Unlock()
		fmt.Printf("Task for %s is already running. Ignoring new task.\n", symbol)
		return
	}
	activeRequests[symbol] = true
	mutex.Unlock()

	// Perform the trade execution
	currentPosition := api.GetPosition(symbol)
	success := utils.UpdatePosition(symbol, signal.TargetPosition, currentPosition)

	if !success {
		// Re-send signal to Redis for retry
		data, _ := json.Marshal(signal)
		config.RedisClient.Publish(config.Ctx, config.SignalChannel, data)
	}

	// Remove the task from activeRequests after completion
	mutex.Lock()
	delete(activeRequests, symbol)
	mutex.Unlock()
}
