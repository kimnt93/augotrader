package main

import (
	"augotrader/internal/notibot"
)

func main() {
	notibot.InitTelegramBot()
	notibot.SendSignalToTelegram("Send me")
}
