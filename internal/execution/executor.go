package execution

import (
	"encoding/json"

	"../config"
	"github.com/rs/zerolog/log"
)

func ReceiveSignals() {
	pubsub := client.RedisClient.Subscribe(config.Ctx, config.SignalChannel)
	defer pubsub.Close()

	ch := pubsub.Channel()

	for msg := range ch {
		var signal execution.Signal
		if err := json.Unmarshal([]byte(msg.Payload), &signal); err != nil {
			log.Error().Msg("Error unmarshalling signal: %v", err)
			continue
		}

		// Start a new goroutine to process the signal
		go execution.ExecuteTrade(signal)

		// Start send Signal to Telegram
	}
}
