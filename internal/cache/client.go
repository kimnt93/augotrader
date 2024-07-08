package cache

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

var (
	RedisClient *redis.Client
	Ctx         = context.Background()
)

func InitRedis() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Warn().Msg("Error loading .env file")
	}

	// Get Redis configuration from environment variables
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB := 0 // Default to DB 0 or adjust as needed
	redisDBStr := os.Getenv("REDIS_DB")
	if redisDBStr != "" {
		redisDB, _ = strconv.Atoi(redisDBStr)
	}

	opt := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: redisPassword,
		DB:       redisDB,
	}

	RedisClient = redis.NewClient(opt)

	// Check connection
	_, err = RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Error().Msgf("Failed to connect to Redis, stop program: %v", err)
		os.Exit(1)
	}

	log.Info().Msg("Redis connected successfully")
}

func init() {
	InitRedis()
}
