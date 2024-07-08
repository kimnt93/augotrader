package cache

import (
	"github.com/rs/zerolog/log"
)

func Ping() {
	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Error().Msgf("Failed to connect to Redis: %v", err)
		return
	}
}
