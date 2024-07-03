package redis

import (
	"encoding/json"
	"log"
	"trading-bot/internal/config"
	"trading-bot/internal/execution"
)

func HandleSignals() {
	pubsub := config.RedisClient.Subscribe(config.Ctx, config.SignalChannel)
	defer pubsub.Close()

	ch := pubsub.Channel()

	for msg := range ch {
		var signal execution.Signal
		if err := json.Unmarshal([]byte(msg.Payload), &signal); err != nil {
			log.Printf("Error unmarshalling signal: %v", err)
			continue
		}

		// Start a new goroutine to process the signal
		go execution.ExecuteTrade(signal)
	}
}
