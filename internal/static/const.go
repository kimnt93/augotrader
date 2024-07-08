package static

// telegram bot config
const (
	CFG_TELEGRAM_BOT_TOKEN = "AUGOTRADER_CFG_TELEGRAM_BOT_TOKEN"
	CFG_TELEGRAM_CHAT_ID   = "AUGOTRADER_CFG_TELEGRAM_CHAT_ID"
)

const (
	// Market config
	CH_LASTEST_PRICE string = "AUGOTRADER_EXEC_LASTEST_PRICE"
	CH_LASTEST_ROOM  string = "AUGOTRADER_EXEC_LASTEST_ROOM"
	CH_LATEST_BAR    string = "AUGOTRADER_EXEC_LATEST_BAR"

	// Signal config
	CH_SIGNAL_CHANNEL string = "AUGOTRADER_EXEC_SIGNAL_CHANNEL"
	CH_LASTEST_SIGNAL string = "AUGOTRADER_EXEC_LASTEST_SIGNAL"

	// All account config
	CH_ALL_ACCOUNT string = "AUGOTRADER_EXEC_ALL_ACCOUNTS" // contains all account ids

	// Account login info
	CH_ACCOUNT_LOGIN_INFO string = "AUGOTRADER_EXEC_ACCOUNT_LOGIN" // Account login info includes: api key, secret key, account id, etc

	// Account trading config includes: strategy weights, target position, target offset, disabled status
	CH_ACCOUNT_STRATEGY_WEIGHTS string = "AUGOTRADER_EXEC_ACCOUNT_STRATEGY_WEIGHTS"
	CH_TARGET_OFFSET            string = "AUGOTRADER_EXEC_TARGET_OFFSET"
	CH_TARGET_POSITION          string = "AUGOTRADER_EXEC_TARGET_POSITION"
	CH_ACCOUNT_DISABLED         string = "AUGOTRADER_EXEC_ACCOUNT_DISABLED"

	// Lock account and symbol before execute trade config
	CH_ACCOUNT_SYMBOL_LOCKED string = "AUGOTRADER_EXEC_ACCOUNT_SYMBOL_LOCKED"
	DEFAULT_LOCK_TTL         int    = 5
)
