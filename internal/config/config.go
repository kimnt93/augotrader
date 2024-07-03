package config

import (
	"log"
	"os"
	"strconv"
	"github.com/rs/zerolog/log"
	"github.com/joho/godotenv"
)

type Config struct {
	// Setup
	TelegramBotID			string
	TelegramBotToken		string
	TelegramBotToken 		string

	// Continue
	SendPositionSeconds int

	RedisHost 			string
	RedisPort 			string
	RedisDB 			string
	RedisUser 			string
	RedisPassword 		string
}

var ExecConfig Config


func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
	  log.Error().Msg("Error loading .env file")
	}
  
	sendPositionSeconds, _ := strconv.Atoi(getEnv("SEND_POSITION_SECONDS", "600"))

	ExecConfig = Config{
		TelegramBotID:      	getEnv("TELEGRAM_BOT_ID")
		TelegramBotToken:      	getEnv("TELEGRAM_BOT_TOKEN")
		TelegramBotToken:      	getEnv("TELEGRAM_CHAT_ID"),
		
		SendPositionSeconds: 	sendPositionSeconds

		RedisHost:     		getEnv("POSTGRES_HOST", "localhost"),
		RedisPort:     		getEnv("POSTGRES_PORT", "6379"),
		RedisDB:     		getEnv("POSTGRES_USER", "0"),
		RedisUser: 			getEnv("REDIS_USER", ""),
		RedisPassword:    	getEnv("REDIS_PASSWORD", ""),

	}
}