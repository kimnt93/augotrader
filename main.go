package main

import (
	"augotrader/internal/execution"
	"augotrader/internal/notibot"
)

func main() {
	notibot.InitTelegramBot()
	execution.Run()
}
