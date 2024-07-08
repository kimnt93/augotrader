package cache

import (
	"strconv"

	"github.com/rs/zerolog/log"
)

func GetKeyInt(key string) (int, error) {
	// Use Redis GET command to retrieve the value associated with the key
	v, err := RedisClient.Get(Ctx, key).Result()
	if err != nil {
		log.Error().Msgf("Failed to get key (%s) offset: %v", key, err)
		return 0, err
	}

	offset, err := strconv.Atoi(v)
	if err != nil {
		log.Error().Msgf("Failed to convert (%s) to integer: %v", v, err)
		return 0, err
	}

	return offset, nil
}

func GetKeyFloat(key string) (float64, error) {
	// Use Redis GET command to retrieve the value associated with the key
	v, err := RedisClient.Get(Ctx, key).Result()
	if err != nil {
		log.Error().Msgf("Failed to get key (%s) offset: %v", key, err)
		return 0, err
	}

	offset, err := strconv.ParseFloat(v, 64)
	if err != nil {
		log.Error().Msgf("Failed to convert (%s) to integer: %v", v, err)
		return 0, err
	}

	return offset, nil
}

func GetKeyStr(key string) (string, error) {
	// Use Redis GET command to retrieve the value associated with the key
	str, err := RedisClient.Get(Ctx, key).Result()
	if err != nil {
		log.Error().Msgf("Failed to get key (%s) offset: %v", key, err)
		return "", err
	}
	return str, nil
}

func GetKeyBoolean(key string) (bool, error) {
	// Use Redis GET command to retrieve the value associated with the key
	valStr, err := RedisClient.Get(Ctx, key).Result()
	if err != nil {
		log.Error().Msgf("Failed to get key (%s) offset: %v", key, err)
		return false, err
	}

	// Convert the retrieved string value (valStr) to an integer
	val, err := strconv.Atoi(valStr)
	if err != nil {
		log.Error().Msgf("Failed to convert (%s) to integer: %v", valStr, err)
		return false, err
	}
	return val != 0, nil
}
