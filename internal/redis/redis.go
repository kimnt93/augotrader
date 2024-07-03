package redis

import (
	"context"
	"fmt"

	"../config"
	"github.com/go-redis/redis/v9"
	"github.com/rs/zerolog/log"
)

var RedisClient *redis.Client
var ctx = context.Background()

func InitRedis() {
	opt := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.ExecConfig.RedisHost, config.ExecConfig.RedisPort),
		Password: config.ExecConfig.RedisPassword, // No password set if empty
		DB:       config.ExecConfig.RedisDB,       // Use default DB if not set
	}

	RedisClient = redis.NewClient(opt)

	// Check connection
	_, err := RedisClient.Ping(config.Ctx).Result()
	if err != nil {
		log.Error().Msg("Failed to connect to Redis: %v", err)
	}
}
