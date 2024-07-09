package execution

import (
	"augotrader/internal/cache"
	"augotrader/internal/notibot"
	"augotrader/internal/services"
	"augotrader/internal/static"
	"augotrader/internal/static/trading"
	"math"

	"github.com/rs/zerolog/log"
)

func executeTrade(accountId string, symbol string) bool {

	// Check if we cannot trade this symbol for this account
	isLocked, err := services.IsLockedAccountSymbol(accountId, symbol)
	if err != nil || isLocked {
		log.Error().Msgf("Account %s is locked for symbol %s (locked=%b) or checking error: %v", accountId, symbol, isLocked, err)
		return false
	}

	// Check if account is disabled
	booksizeConfig, err := services.GetCurrentBooksize(accountId, symbol)
	if err != nil || booksizeConfig.IsDisabled {
		log.Error().Msgf("Account %s is disabled (disabled=%b) or checking error: %v", accountId, booksizeConfig.IsDisabled, err)
		return false
	}

	// Lock account and symbol before execute trade
	services.LockAccountSymbol(accountId, symbol)

	// current tradign strategy
	tradingStrategies, err := services.GetTradingStrategyWeightByAccount(accountId)
	if err != nil {
		log.Error().Msgf("Error getting trading strategy: %v", err)
		return false
	}

	// Calculate expected position based on weights
	positionWeight := 0.0
	for strategyName, strategyValue := range tradingStrategies {
		log.Debug().Msgf("Strategy Name: %s, Weight: %f", strategyName, strategyValue.Weight)

		// Get current signal and calculate expected position
		currentSignal, err := services.GetCurrentSignalByName(strategyValue.Name)
		if err != nil {
			log.Error().Msgf("Error getting current signal: %v", err)
			return false
		}
		positionWeight += currentSignal.Position * strategyValue.Weight
	}

	// Get target position
	targetPosition := booksizeConfig.TargetPosition
	currentOffset := booksizeConfig.Offset

	expectedPosition := positionWeight * targetPosition

	// get current position
	currentPosition, err := services.GetCurrentPosition(accountId, symbol)
	if err != nil {
		log.Error().Msgf("Error while getting position (%s, %s): %v", accountId, symbol, err)
		return false
	}

	// get accouht current offset
	pendingOrders, err := services.GetAccountSymbolPendingOrders(accountId, symbol)
	if err != nil {
		log.Error().Msgf("Error getting pending orders: %v", err)
		return false
	}

	// loop over pending orders and calculate total pending quantity
	totalPendingQty := 0.0
	for _, order := range pendingOrders {
		if order.OrderSide == trading.BUY_SIDE {
			totalPendingQty += order.OrderQty
		} else {
			totalPendingQty -= order.OrderQty
		}
	}

	// Calculate the quantity to execute
	// Remember checking total pending quantity to avoid over trade
	// Round and convert by symbol
	// 1. include 3 characters in upper case: round to lot 100
	// 2. start with VN30: round to lot 1
	// 3. other: float value
	executeQty := expectedPosition - currentPosition.Position + currentOffset - totalPendingQty
	if executeQty == 0 {
		log.Info().Msgf("No trade needed for account %s, symbol %s", accountId, symbol)
		return true
	} else {
		// excute side B or S
		executeSide := trading.BUY_SIDE
		if executeQty < 0 {
			executeSide = "S"
			executeQty = math.Abs(executeQty)
		}

		// Execute trade
		// Create/cancel, etc
		currentExecutePrice := 0.0
		services.CreateOrder(accountId, symbol, currentExecutePrice, executeSide, executeQty)
	}

	// Release the lock for the next trade
	services.UnlockAccountSymbol(accountId, symbol)
	return true
}

// Run is the main function to execute the trade
func Run() {
	pubsub := cache.RedisClient.Subscribe(cache.Ctx, static.CH_SIGNAL_CHANNEL)
	defer pubsub.Close()

	ch := pubsub.Channel()

	for msg := range ch {
		allAccountids, err := services.GetAllLoginInfo()
		if err != nil {
			log.Error().Msgf("Error getting all account ids: %v", err)
			continue
		}
		receivedActionSymbol := msg.Payload // HPG BTCUSDT VN30F2403
		log.Info().Msgf("Received signal action symbol: %s", receivedActionSymbol)

		// Loop all accounts and execute trade
		for _, loginInfo := range allAccountids {
			if loginInfo.IsDisabled {
				log.Info().Msgf("Account %s is disabled", loginInfo.AccountId)
				continue
			}
			// Run in goroutine for each accountId
			go executeTrade(loginInfo.AccountId, receivedActionSymbol)
		}

		// Start send Signal to Telegram
		go notibot.SendSignalToTelegram("")
	}
}
