// Send message to telegram

package notibot

import (
	"augotrader/internal/cache"
	"augotrader/internal/static"
	"augotrader/internal/static/errors"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/rs/zerolog/log"
)

type TelegramConfig struct {
	BotToken string
	ChatId   int64
}

var (
	tgConfig *TelegramConfig
	tgBotAPI *tgbotapi.BotAPI
	once     sync.Once
)

func InitTelegramBot() {
	once.Do(func() {
		initTelegramAPI()
	})
}

func initTelegramAPI() string {
	// Fetch configuration values from cache or wherever you store them
	botToken, _ := cache.GetKeyStr(static.CFG_TELEGRAM_BOT_TOKEN)
	chatId, _ := cache.GetKeyInt(static.CFG_TELEGRAM_CHAT_ID)

	// Initialize the Telegram configuration
	tgConfig = &TelegramConfig{
		BotToken: botToken,
		ChatId:   int64(chatId),
	}

	// Initialize the Telegram bot API client
	api, err := tgbotapi.NewBotAPI(tgConfig.BotToken)
	if err != nil {
		log.Error().Msgf("Failed to initialize Telegram bot API: %v", err)
		return errors.INIT_FAILED
	}

	log.Info().Msg("Telegram bot API initialized successfully")
	tgBotAPI = api
	return errors.INIT_SUCCESS
}

func SendSignalToTelegram(symbol string) string {

	if tgBotAPI == nil {
		log.Error().Msg("Telegram bot API is not initialized")
		return errors.TELEGRAM_BOT_NOT_INITIALIZED
	}

	// Example: Send message to Telegram
	msg := tgbotapi.NewMessage(tgConfig.ChatId, symbol)
	_, err := tgBotAPI.Send(msg)
	if err != nil {
		log.Error().Err(err).Msg("Failed to send message to Telegram")
		return errors.TELEGRAM_BOT_SEND_MESSAGE_ERROR
	}

	log.Info().Msgf("Signal '%s' sent to Telegram", symbol)
	return errors.TELEGRAM_BOT_NOT_SUCCESSFULLY_SENT
}

func SendTradingSummaryToTelegram() {
	// Send trading summary to telegram
	log.Info().Msg("Send trading summary to Telegram")
}
