package config

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

const ()

type TradingAccountConfig struct {
	AccountId       string
	TradingStraregy string
}

type ExecConfig struct {
	AccountId        string
	TradingStrategy  string
	LastestAlpha     string
	TelegramBotId    string
	TelegramBotToken string
	AppPort          int
}

var Config ExecConfig

func LoadExecConfigFromRedis(redisClient *redis.Client, ctx context.Context) {
	accountConfig, err := redisClient.Get(ctx, "AUGOTRADER_MY_ACCOUNT_CONFIG.Q22023-SD01").Result()
	if err != nil {
		log.Error().Msgf("Failed to load account_id from Redis: %v", err)
	}

	tradingStrategy, err := redisClient.Get(ctx, "trading_strategy").Result()
	if err != nil {
		log.Error().Msgf("Failed to load trading_strategy from Redis: %v", err)
	}

	lastestAlpha, err := redisClient.Get(ctx, "lastest_alpha").Result()
	if err != nil {
		log.Error().Msgf("Failed to load lastest_alpha from Redis: %v", err)
	}

	telegramBotId, err := redisClient.Get(ctx, "telegram_bot_id").Result()
	if err != nil {
		log.Error().Msgf("Failed to load telegram_bot_id from Redis: %v", err)
	}

	telegramBotToken, err := redisClient.Get(ctx, "telegram_bot_token").Result()
	if err != nil {
		log.Error().Msgf("Failed to load telegram_bot_token from Redis: %v", err)
	}

	appPort, err := redisClient.Get(ctx, "app_port").Result()
	if err != nil {
		log.Error().Msgf("Failed to load app_port from Redis: %v", err)
	}
	appPortInt, _ := strconv.Atoi(appPort)

	Config = ExecConfig{
		AccountId:        accountId,
		TradingStrategy:  tradingStrategy,
		LastestAlpha:     lastestAlpha,
		TelegramBotId:    telegramBotId,
		TelegramBotToken: telegramBotToken,
		AppPort:          appPortInt,
	}
}
